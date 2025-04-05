package service

import (
	"context"
	"database/sql"
	"time"

	storage "github.com/DarYur13/learn-control/internal/storage/learn_control"
	"github.com/pkg/errors"
)

func (s *Service) CloseTaskWithTrainingDateSet(ctx context.Context, taskID, emplID, trainingID int, taskType string, date time.Time) error {
	task, needNextTask, err := s.nextTask(ctx, emplID, trainingID, taskType)
	if err != nil {
		return errors.WithMessage(err, "create next task")
	}

	if err := s.txManager.Do(ctx, func(tx *sql.Tx) error {
		_, txErr := s.storage.UpdateEmployeeTrainingDateTx(ctx, tx, emplID, trainingID, date)
		if txErr != nil {
			return errors.WithMessage(txErr, "set training date")
		}

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
