package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/DarYur13/learn-control/internal/domain"
	storage "github.com/DarYur13/learn-control/internal/storage/learn_control"
	"github.com/pkg/errors"
)

func (s *Service) UpdateEmployeeTrainingDate(ctx context.Context, employeeID int, trainingID int, date time.Time) (*domain.TrainingDates, error) {
	var dates *storage.TrainingDates

	if err := s.txManager.Do(ctx, func(tx *sql.Tx) error {
		var txErr error

		dates, txErr = s.storage.UpdateEmployeeTrainingDateTx(ctx, tx, employeeID, trainingID, date)
		if txErr != nil {
			return errors.WithMessage(txErr, "update employee training date")
		}

		return nil
	}); err != nil {
		return nil, err
	}

	result := formatTrainingDates(*dates)

	return &result, nil
}
