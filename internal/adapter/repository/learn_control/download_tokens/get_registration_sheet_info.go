package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
	"github.com/google/uuid"
)

const (
	queryGetRegistrationSheetInfo = `
	SELECT 
		i.full_name AS instructor_name,
		i.position AS instructor_position,
		e.full_name AS employee_name,
		e.birth_date,
		e.position AS employee_position,
		e.department,
		t.training_type,
		COALESCE(string_agg(la.act_name, E'\n'), '') AS acts
	FROM download_tokens dt
	JOIN employees e ON e.id = dt.employee_id
	JOIN employees i ON i.department = e.department AND i.is_leader = TRUE
	JOIN trainings t ON t.id = dt.training_id
	LEFT JOIN acts_trainings at ON at.training_id = t.id
	LEFT JOIN local_acts la ON la.id = at.local_act_id
	WHERE dt.token = $1
	GROUP BY 
		i.full_name,
		i.position,
		e.full_name,
		e.birth_date,
		e.position,
		e.department,
		t.training_type
	`
)

func (dts *downloadTokensStorage) GetRegistrationSheetInfo(ctx context.Context, token uuid.UUID) (*domain.RegistrationSheetInfo, error) {
	var info domain.RegistrationSheetInfo

	err := dts.db.QueryRowContext(ctx, queryGetRegistrationSheetInfo, token).Scan(
		&info.InstructorName,
		&info.InstructorPosition,
		&info.EmployeeName,
		&info.EmployeeBirthDate,
		&info.EmployeePosition,
		&info.EmployeeDepartment,
		&info.TrainingType,
		&info.Acts,
	)

	if err != nil {
		return nil, err
	}

	return &info, nil
}
