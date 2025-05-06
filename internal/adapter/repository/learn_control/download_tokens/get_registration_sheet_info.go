package learncontrol

import (
	"context"
	"database/sql"

	"github.com/DarYur13/learn-control/internal/domain"
	"github.com/google/uuid"
)

const (
	queryGetRegistrationSheetInfo = `
	SELECT 
		oss.full_name AS occupational_safety_specialist_name,
		oss.position AS occupational_safety_specialist_position,
		i.full_name AS instructor_name,
		i.position AS instructor_position,
		e.full_name AS employee_name,
		e.birth_date,
		e.position AS employee_position,
		e.department,
		t.training_type,
		COALESCE(string_agg(la.act_name, E'\n'), '') AS acts,
		dt.loads_count
	FROM download_tokens dt
	JOIN employees e ON e.id = dt.employee_id
	JOIN employees i ON i.department = e.department AND i.is_leader = TRUE
	JOIN employees oss ON oss.department = 'Отдел охраны труда'
	JOIN trainings t ON t.id = dt.training_id
	LEFT JOIN acts_trainings at ON at.training_id = t.id
	LEFT JOIN local_acts la ON la.id = at.local_act_id
	WHERE dt.token = $1
	GROUP BY 
		oss.full_name,
		oss.position,
		i.full_name,
		i.position,
		e.full_name,
		e.birth_date,
		e.position,
		e.department,
		t.training_type,
		dt.loads_count
	`
)

// TODO fix case when more than one department leader and more than one occupational safety specialist
func (dts *downloadTokensStorage) GetRegistrationSheetInfoTx(ctx context.Context, tx *sql.Tx, token uuid.UUID) (*domain.RegistrationSheetInfo, error) {
	var (
		info       domain.RegistrationSheetInfo
		loadsCount int
	)

	err := tx.QueryRowContext(ctx, queryGetRegistrationSheetInfo, token).Scan(
		&info.OccupSafetySpecName,
		&info.OccupSafetySpecPosition,
		&info.InstructorName,
		&info.InstructorPosition,
		&info.EmployeeName,
		&info.EmployeeBirthDate,
		&info.EmployeePosition,
		&info.EmployeeDepartment,
		&info.TrainingType,
		&info.Acts,
		&loadsCount,
	)

	if err != nil {
		return nil, err
	}

	if loadsCount > loadsCountLimit {
		return nil, ErrTooManyLoads
	}

	return &info, nil
}
