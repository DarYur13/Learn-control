package service

import (
	"context"
	"database/sql"

	storage "github.com/DarYur13/learn-control/internal/storage/learn_control"
	"github.com/pkg/errors"
)

func (s *Service) CloseAssignTask(ctx context.Context, taskID, employeeID, trainingID int, taskType string) error {
	task, needNextTask, err := s.nextTask(ctx, employeeID, trainingID, taskType)
	if err != nil {
		return errors.WithMessage(err, "create next task")
	}

	if err := s.txManager.Do(ctx, func(tx *sql.Tx) error {
		if txErr := s.storage.CloseTaskTx(ctx, tx, taskID); txErr != nil {
			return errors.WithMessage(txErr, "close task")
		}

		if needNextTask {
			if txErr := s.storage.AddTaskTx(ctx, tx, storage.TaskBaseInfo(*task)); txErr != nil {
				return errors.WithMessage(txErr, "add task")
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
