package storage

import (
	"context"
	"database/sql"

	"github.com/golang-sql/sqlexp"
)

const (
	queryGetEmployesWithoutTrainings = `
	SELECT e.id FROM employees e
	JOIN positions p ON e.position = p.position AND e.department = p.department
	JOIN position_trainings pt ON p.id = pt.position_id
	LEFT JOIN employee_trainings et ON e.id = et.employee_id AND et.training_id = pt.training_id
	WHERE p.id = $1 AND et.employee_id IS NULL;
	`
)

func (s *Storage) GetEmployeesWithoutTrainings(ctx context.Context, positionID int) ([]int, error) {
	return s.getEmployeesWithoutTrainings(ctx, s.db, positionID)
}

func (s *Storage) GetEmployeesWithoutTrainingsTx(ctx context.Context, tx *sql.Tx, positionID int) ([]int, error) {
	return s.getEmployeesWithoutTrainings(ctx, tx, positionID)
}

func (s *Storage) getEmployeesWithoutTrainings(ctx context.Context, tx sqlexp.Querier, positionID int) ([]int, error) {
	var employeesIDs []int

	rows, err := tx.QueryContext(ctx, queryGetEmployesWithoutTrainings, positionID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var emplID int

		if err := rows.Scan(&emplID); err != nil {
			return nil, err
		}

		employeesIDs = append(employeesIDs, emplID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employeesIDs, nil
}
