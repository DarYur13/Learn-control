package storage

import (
	"context"
)

const (
	queryGetEmployee = `
		SELECT 
			e.full_name, 
			e.birth_date, 
			e.snils, 
			e.department, 
			e.position, 
			t.training, 
			et.training_date AS pass_date,       -- добавляем алиас
			et.retraining_date AS repass_date    -- добавляем алиас
		FROM employees e
		JOIN employee_trainings et ON e.id = et.employee_id
		JOIN trainings t ON et.training_id = t.id
		WHERE e.id = $1
		ORDER BY et.training_date DESC, t.training
	`
)

func (s *Storage) GetEmployee(ctx context.Context, id int) (*Employee, error) {
	var (
		result Employee
	)

	rows, err := s.db.QueryContext(ctx, queryGetEmployee, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var trainings Training

		rows.Scan(
			&result.FullName,
			&result.BirthDate,
			&result.Snils,
			&result.Department,
			&result.Position,
			&trainings.Name,
			&trainings.PassDate,
			&trainings.RePassDate,
		)

		result.Trainings = append(result.Trainings, trainings)
	}

	return &result, nil
}
