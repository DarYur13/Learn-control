package storage

import (
	"context"
	"database/sql"
)

const (
	queryGetEmployeesByFilters = `
	SELECT 
		e.id, 
		e.full_name, 
		e.department, 
		e.position, 
		COALESCE(t.training, 'Нет обучения') AS training, 
		et.training_date, 
		et.retraining_date
	FROM employees e
	LEFT JOIN employee_trainings et ON e.id = et.employee_id
	LEFT JOIN trainings t ON et.training_id = t.id
	WHERE 
		(COALESCE($1::TEXT, '') = '' OR e.department = $1::TEXT)
		AND (COALESCE($2::TEXT, '') = '' OR e.position = $2::TEXT)
		AND (COALESCE($3::INTEGER, -1) = -1 OR et.training_id = $3::INTEGER)
		AND (
			(COALESCE($4::DATE, '1000-01-01') = '1000-01-01' AND COALESCE($5::DATE, '9999-12-31') = '9999-12-31')
			OR (COALESCE($4::DATE, '1000-01-01') != '1000-01-01' AND et.training_date >= $4::DATE)
			OR (COALESCE($5::DATE, '9999-12-31') != '9999-12-31' AND et.training_date <= $5::DATE)
		)
		AND (COALESCE($6::BOOLEAN, FALSE) = FALSE OR (et.training_id IS NOT NULL AND et.training_date IS NULL))
		AND (COALESCE($7::INTEGER, -1) = -1 OR (et.retraining_date IS NOT NULL AND et.retraining_date <= CURRENT_DATE + INTERVAL '1 day' * $7::INTEGER))
	ORDER BY e.id;
	`
)

func (s *Storage) GetEmployeesByFilters(ctx context.Context, filters Filters) ([]EmployeeInfo, error) {
	rows, err := s.db.QueryContext(ctx, queryGetEmployeesByFilters,
		filters.Department,
		filters.Position.String,
		filters.TrainingID,
		filters.DateFrom,
		filters.DateTo,
		filters.TrainingsNotPassed,
		filters.RetrainingIn,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Хранение данных сотрудников
	employeesMap := make(map[string]*EmployeeInfo)

	for rows.Next() {
		var (
			id         int
			fullName   string
			department string
			position   string
			training   sql.NullString
			passDate   sql.NullTime
			rePassDate sql.NullTime
		)

		if err := rows.Scan(&id, &fullName, &department, &position, &training, &passDate, &rePassDate); err != nil {
			return nil, err
		}

		empKey := fullName + department + position // Ключ для группировки сотрудников
		if _, exists := employeesMap[empKey]; !exists {
			employeesMap[empKey] = &EmployeeInfo{
				FullName:   fullName,
				Department: department,
				Position:   position,
				Trainings:  []Training{},
			}
		}

		// Добавляем обучение, если оно есть
		if training.Valid {
			employeesMap[empKey].Trainings = append(employeesMap[empKey].Trainings, Training{
				Name: training.String,
				TrainingDates: TrainingDates{
					PassDate:   passDate,
					RePassDate: rePassDate,
				},
			})
		}
	}

	// Преобразуем map в slice
	employees := make([]EmployeeInfo, 0, len(employeesMap))
	for _, emp := range employeesMap {
		employees = append(employees, *emp)
	}

	return employees, nil
}
