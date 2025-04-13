package learncontrol

import (
	"context"

	trainingsStorage "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/trainings"
)

const (
	queryGetEmployee = `
		SELECT 
			e.full_name, 
			e.birth_date, 
			e.snils, 
			e.department, 
			e.position, 
			e.employment_date,
			t.training_name, 
			t.training_type,
			et.training_date AS pass_date,
			et.retraining_date AS repass_date,
			et.has_protocol
		FROM employees e
		JOIN employee_trainings et ON e.id = et.employee_id
		JOIN trainings t ON et.training_id = t.id
		WHERE e.id = $1
		ORDER BY et.training_date DESC, t.training
	`
)

func (es *EmployeesStorage) GetEmployeePersonalCard(ctx context.Context, id int) (*EmployeePersonalCard, error) {
	var (
		result EmployeePersonalCard
	)

	rows, err := es.db.QueryContext(ctx, queryGetEmployee, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var trainings trainingsStorage.Training

		if err := rows.Scan(
			&result.FullName,
			&result.BirthDate,
			&result.Snils,
			&result.Department,
			&result.Position,
			&result.EmploymentDate,
			&trainings.Name,
			&trainings.Type,
			&trainings.PassDate,
			&trainings.RePassDate,
			&trainings.HasProtocol,
		); err != nil {
			return nil, err
		}

		result.Trainings = append(result.Trainings, trainings)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &result, nil
}
