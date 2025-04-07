package learncontrol

import (
	"context"
	"database/sql"
)

const (
	queryGetEmployeesByFilters = `
	WITH filtered_employees AS (
		SELECT * FROM employees
		WHERE ($1 IS NULL OR department = $1)
		AND ($2 IS NULL OR position = $2)
	),
	filtered_trainings AS (
		SELECT * FROM employee_trainings
		WHERE ($3 IS NULL OR training_id = $3)
		AND (
			($4 IS NULL AND $5 IS NULL)
			OR ($4 IS NOT NULL AND training_date >= $4)
			OR ($5 IS NOT NULL AND training_date <= $5)
		)
		AND ($7 IS NULL OR (retraining_date IS NOT NULL AND retraining_date <= CURRENT_DATE + INTERVAL '1 day' * $7))
		AND ($8 IS NULL OR has_protocol = $8)
	)
	SELECT 
		e.id, 
		e.full_name, 
		e.department, 
		e.position, 
		t.training, 
		et.training_date, 
		et.retraining_date,
		et.has_protocol
	FROM filtered_employees e
	LEFT JOIN filtered_trainings et ON e.id = et.employee_id
	LEFT JOIN trainings t ON et.training_id = t.id
	ORDER BY e.id;
	`
)

func (es *EmployeesStorage) GetEmployeesByFilters(ctx context.Context, filters Filters) ([]EmployeeInfo, error) {
	rows, err := es.db.QueryContext(ctx, queryGetEmployeesByFilters,
		filters.Department,
		filters.Position.String,
		filters.TrainingID,
		filters.DateFrom,
		filters.DateTo,
		filters.TrainingsNotPassed,
		filters.RetrainingIn,
		filters.HasProtocol,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Хранение данных сотрудников
	employeesMap := make(map[string]*EmployeeInfo)

	for rows.Next() {
		var (
			id          int
			fullName    string
			department  string
			position    string
			training    sql.NullString
			passDate    sql.NullTime
			rePassDate  sql.NullTime
			hasProtocol sql.NullBool
		)

		if err := rows.Scan(
			&id,
			&fullName,
			&department,
			&position,
			&training,
			&passDate,
			&rePassDate,
			&hasProtocol,
		); err != nil {
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
				HasProtocol: hasProtocol,
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
