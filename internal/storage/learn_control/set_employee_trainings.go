package storage

import (
	"context"
	"database/sql"
)

const (
	querySetEmployeeTrainings = `
	INSERT INTO employee_trainings (
		employee_id,
		training_id
	) VALUES ($1, $2)
	ON CONFLICT (employee_id, training_id) DO NOTHING
	`
)

func (s *Storage) SetEmployeeTrainingsTx(ctx context.Context, tx *sql.Tx, employeeID int, trainingIDs []int) error {
	stmt, err := tx.PrepareContext(ctx, querySetEmployeeTrainings)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, trainingID := range trainingIDs {
		if _, err := stmt.ExecContext(ctx, employeeID, trainingID); err != nil {
			return err
		}
	}

	return nil
}
