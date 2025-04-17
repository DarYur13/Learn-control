package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

func (s *Service) CloseTaskWithTrainingProtocolConfirm(ctx context.Context, taskID int) error {
	taskInfo, err := s.tasksStorage.GetTaskInfoByID(ctx, taskID)
	if err != nil {
		return errors.WithMessage(err, "get task info")
	}

	if err := s.txManager.Do(ctx, func(tx *sql.Tx) error {
		if txErr := s.employeesStorage.SetEmployeeTrainnigProtocolTx(ctx, tx, int(taskInfo.EmployeeID.Int64), int(taskInfo.TrainingID.Int64)); txErr != nil {
			return errors.WithMessage(txErr, "set has protocol")
		}

		if txErr := s.tasksStorage.CloseTaskTx(ctx, tx, taskID); txErr != nil {
			return errors.WithMessage(txErr, "close task")
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
