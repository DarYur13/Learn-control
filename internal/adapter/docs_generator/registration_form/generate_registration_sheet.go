package docsgenerator

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"path/filepath"
	"time"

	"github.com/DarYur13/learn-control/internal/domain"
	"github.com/nguyenthenguyen/docx"
)

func (g *docxGenerator) GenerateRegistrationSheet(ctx context.Context, info domain.RegistrationSheetInfo) (io.Reader, error) {
	fileName, err := g.getTemplateFilename(info.TrainingType)
	if err != nil {
		return nil, err
	}

	templatePath := filepath.Join(g.templateDir, fileName)

	r, err := docx.ReadDocxFile(templatePath)
	if err != nil {
		return nil, fmt.Errorf("read template: %w", err)
	}
	defer r.Close()

	doc := r.Editable()

	doc.Replace("{date}", time.Now().Format(domain.DateFormat), -1)
	doc.Replace("{full_name}", info.EmployeeName, -1)
	doc.Replace("{position}", info.EmployeePosition, -1)
	doc.Replace("{birth_date}", info.EmployeeBirthDate.Format(domain.DateFormat), -1)

	if info.TrainingType == domain.TrainingTypeIntroductory {
		doc.Replace("{department}", info.EmployeeDepartment, -1)
		doc.Replace("{executor}", info.OccupSafetySpecName, -1)
		doc.Replace("{executor_position}", info.OccupSafetySpecPosition, -1)
	}

	if info.TrainingType == domain.TrainingTypeInitial || info.TrainingType == domain.TrainingTypeRefresher {
		doc.Replace("{act}", info.Acts, -1)
		doc.Replace("{executor}", info.InstructorName, -1)
		doc.Replace("{executor_position}", info.InstructorPosition, -1)
	}

	buf := new(bytes.Buffer)
	if err := doc.Write(buf); err != nil {
		return nil, fmt.Errorf("write filled docx: %w", err)
	}

	return buf, nil
}

func (g *docxGenerator) getTemplateFilename(trainingType domain.TrainingType) (string, error) {
	switch trainingType {
	case domain.TrainingTypeIntroductory:
		return "вводный_инструктаж_шаблон.docx", nil
	case domain.TrainingTypeInitial:
		return "первичный_инструктаж_шаблон.docx", nil
	case domain.TrainingTypeRefresher:
		return "повторный_инструктаж_шаблон.docx", nil
	default:
		return "", fmt.Errorf("unknown training type: %s", trainingType)
	}
}
