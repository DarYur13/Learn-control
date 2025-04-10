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

func (g *docxGenerator) GenerateRegistrationSheet(ctx context.Context, employee domain.Employee, briefingInfo domain.BriefingInfo) (io.Reader, error) {
	fileName, err := g.getTemplateFilename(briefingInfo.TrainingType)
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

	doc.Replace("{date}", time.Now().Format("02.01.2006"), -1)
	doc.Replace("{full_name}", employee.FullName, -1)
	doc.Replace("{position}", employee.Position, -1)
	doc.Replace("{birth_date}", employee.BirthDate, -1)
	doc.Replace("{executor}", briefingInfo.Instructor.FullName, -1)
	doc.Replace("{executor_position}", briefingInfo.Instructor.Position, -1)

	if briefingInfo.TrainingType == domain.TrainingTypeIntroductory {
		doc.Replace("{department}", employee.Department, -1)
	}

	if briefingInfo.TrainingType == domain.TrainingTypeInitial || briefingInfo.TrainingType == domain.TrainingTypeRefresher {
		doc.Replace("{act}", briefingInfo.Act, -1)
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
