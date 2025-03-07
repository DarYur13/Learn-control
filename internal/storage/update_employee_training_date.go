package storage

import (
	"context"
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

func (s *Storage) UpdateEmployeeTrainingDate(ctx context.Context, employeeID int, trainingID int, date time.Time) (*TrainingDates, error) {
	var trainingDates TrainingDates

	err := s.db.QueryRowContext(ctx, queryUpdateEmployeeTrainingDate,
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
