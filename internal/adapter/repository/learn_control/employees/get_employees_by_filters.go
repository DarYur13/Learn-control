package learncontrol

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	trainingsStorage "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/trainings"
	"github.com/DarYur13/learn-control/internal/domain"
	sq "github.com/Masterminds/squirrel"
)

const (
	queryGetEmployees = `
	SELECT 
		e.id,
		e.full_name,
		e.department,
		e.position,
		t.training_name,
		t.training_type,
		et.training_date,
		et.retraining_date,
		et.has_protocol
	FROM employees e
	LEFT JOIN employee_trainings et ON e.id = et.employee_id
	LEFT JOIN trainings t ON et.training_id = t.id
	`
)

func (es *EmployeesStorage) GetEmployeesByFilters(ctx context.Context, filters Filters) ([]EmployeeInfo, error) {
	queryWithFilters, args, err := buildQueryWithFilters(filters)
	if err != nil {
		return nil, err
	}

	rows, err := es.db.QueryContext(ctx, queryWithFilters, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employeesMap := make(map[string]*EmployeeInfo)

	for rows.Next() {
		var (
			id           int
			fullName     string
			department   string
			position     string
			training     sql.NullString
			trainingType domain.TrainingType
			passDate     sql.NullTime
			rePassDate   sql.NullTime
			hasProtocol  sql.NullBool
		)

		if err := rows.Scan(
			&id,
			&fullName,
			&department,
			&position,
			&training,
			&trainingType,
			&passDate,
			&rePassDate,
			&hasProtocol,
		); err != nil {
			return nil, err
		}

		key := fmt.Sprintf("%s|%s|%s", fullName, department, position)
		if _, ok := employeesMap[key]; !ok {
			employeesMap[key] = &EmployeeInfo{
				FullName:   fullName,
				Department: department,
				Position:   position,
				Trainings:  []trainingsStorage.Training{},
			}
		}

		if training.Valid {
			employeesMap[key].Trainings = append(employeesMap[key].Trainings, trainingsStorage.Training{
				Name: training.String,
				Type: trainingType,
				TrainingDates: trainingsStorage.TrainingDates{
					PassDate:   passDate,
					RePassDate: rePassDate,
				},
				HasProtocol: hasProtocol,
			})
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	result := make([]EmployeeInfo, 0, len(employeesMap))
	for _, emp := range employeesMap {
		result = append(result, *emp)
	}

	return result, nil
}

func buildQueryWithFilters(filters Filters) (string, []interface{}, error) {
	builder := sq.SelectBuilder{}.
		Prefix(queryGetEmployees).
		PlaceholderFormat(sq.Dollar).
		Suffix("ORDER BY e.full_name ASC")

	if filters.Department.Valid {
		builder = builder.Where(sq.Eq{"e.department": filters.Department.String})
	}

	if filters.Position.Valid {
		builder = builder.Where(sq.Eq{"e.position": filters.Position.String})
	}

	if filters.TrainingID.Valid {
		builder = builder.Where(sq.Eq{"et.training_id": filters.TrainingID.Int64})
	}

	if filters.DateFrom.Valid {
		builder = builder.Where(sq.GtOrEq{"et.training_date": filters.DateFrom.Time})
	}

	if filters.DateTo.Valid {
		builder = builder.Where(sq.LtOrEq{"et.training_date": filters.DateTo.Time})
	}
	if filters.RetrainingIn.Valid {
		builder = builder.Where(sq.GtOrEq{"et.retraining_date": time.Now().AddDate(0, 0, int(filters.RetrainingIn.Int64))})
	}

	if filters.HasProtocol.Valid {
		builder = builder.Where(sq.Eq{"et.has_protocol": filters.HasProtocol.Bool})
	}

	if filters.TrainingsNotPassed.Valid && filters.TrainingsNotPassed.Bool {
		builder = builder.Where("et.training_date IS NULL")
	}

	return builder.ToSql()
}
