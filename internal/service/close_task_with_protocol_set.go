package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

func (s *Service) CloseTaskWithTrainingProtocolConfirm(ctx context.Context, taskID, employeeID, trainingID int) error {
	if err := s.txManager.Do(ctx, func(tx *sql.Tx) error {
		if txErr := s.storage.SetHasProtocolTx(ctx, tx, employeeID, trainingID); txErr != nil {
			return errors.WithMessage(txErr, "set has protocol")
		}

		if txErr := s.storage.CloseTaskTx(ctx, tx, taskID); txErr != nil {
			return errors.WithMessage(txErr, "close task")
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
