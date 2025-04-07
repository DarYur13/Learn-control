package storage

import (
	"context"
	"database/sql"
	"time"
)

const (
	queryUpdateEmployeeTrainingDate = `
	UPDATE employee_trainings et
	SET 
		et.training_date = $1,
		et.retraining_date = $1 + (t.valid_period || ' months')::interval
	FROM trainings t
	WHERE 
		et.training_id = t.id AND
		et.employee_id = $2 AND 
		et.training_id = $3
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
