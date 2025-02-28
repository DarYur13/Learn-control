package storage

import (
	"context"
	"encoding/json"
)

const (
	queryGetEmployeesByFilters = `
	WITH EmployeeTrainings AS (
    SELECT 
        et.employee_id,
        t.training AS name,
        et.training_date AS pass_date,
        et.retraining_date AS re_pass_date
    FROM employee_trainings et
    JOIN trainings t ON et.training_id = t.id
    WHERE 
        ($3 IS NULL OR et.training_date >= $3)
        AND ($4 IS NULL OR et.training_date <= $4)
	)
	SELECT 
		e.full_name,
		e.department,
		e.position,
		COALESCE(
			json_agg(
				json_build_object(
					'name', et.name,
					'pass_date', et.pass_date,
					're_pass_date', et.re_pass_date
				)
			) FILTER (WHERE et.name IS NOT NULL),
			'[]'
		) AS trainings
	FROM employees e
	LEFT JOIN EmployeeTrainings et ON e.id = et.employee_id
	WHERE 
		($1 IS NULL OR e.department = $1)
		AND ($2 IS NULL OR e.position = $2)
		AND ($5 IS NULL OR e.id IN (
			SELECT et.employee_id FROM employee_trainings et WHERE et.training_id = $5
		))
		AND ($6 IS NULL OR (
			$6 = TRUE AND e.id IN (
				SELECT DISTINCT et.employee_id 
				FROM employee_trainings et 
				WHERE et.training_date IS NULL -- У сотрудника нет даты прохождения обучения
			)
		))
		AND ($7 IS NULL OR e.id IN (
			SELECT et.employee_id FROM employee_trainings et 
			WHERE et.retraining_date IS NOT NULL 
			AND et.retraining_date <= CURRENT_DATE + ($7 || ' days')::INTERVAL
		))
	GROUP BY e.id
	ORDER BY e.full_name;
	`
)

func (s *Storage) GetEmployeesByFilters(ctx context.Context, filters Filters) ([]EmployeeInfo, error) {
	var employees []EmployeeInfo

	rows, err := s.db.QueryContext(ctx, queryGetEmployeesByFilters,
		filters.Department,
		filters.Position,
		filters.DateFrom,
		filters.DateTo,
		filters.TrainingID,
		filters.TrainingsNotPassed,
		filters.RetrainingIn,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee EmployeeInfo
		var trainingsJSON string

		if err := rows.Scan(
			&employee.FullName,
			&employee.Department,
			&employee.Position,
			&trainingsJSON,
		); err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(trainingsJSON), &employee.Trainings); err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}
