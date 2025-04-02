package storage

import (
	"context"
	"database/sql"
	"time"
)

const (
	queryUpdateEmployeeTrainingDate = `
	UPDATE employee_trainings 
	SET training_date = $1
	WHERE employee_id = $2 AND training_id = $3
	RETURNING training_date, retraining_date;
	`
)

func (s *Storage) UpdateEmployeeTrainingDateTx(ctx context.Context, tx *sql.Tx, employeeID int, trainingID int, date time.Time) (*TrainingDates, error) {
	var trainingDates TrainingDates

	err := tx.QueryRowContext(ctx, queryUpdateEmployeeTrainingDate,
		date,
		employeeID,
		trainingID,
	).Scan(
		&trainingDates.PassDate,
		&trainingDates.RePassDate,
	)
	if err != nil {
		return nil, err
	}

	return &trainingDates, nil
}
