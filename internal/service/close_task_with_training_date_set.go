package service

import (
	"context"
	"database/sql"
	"time"

	tasksStorage "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/tasks"
	"github.com/DarYur13/learn-control/internal/domain"
	"github.com/pkg/errors"
)

func (s *Service) CloseTaskWithTrainingDateSet(ctx context.Context, taskID, emplID, trainingID int, taskType domain.TaskType, date time.Time) error {
	task, needNextTask, err := s.nextTask(ctx, emplID, trainingID, taskType)
	if err != nil {
		return errors.WithMessage(err, "create next task")
	}

	if err := s.txManager.Do(ctx, func(tx *sql.Tx) error {
		_, txErr := s.employeesStorage.UpdateEmployeeTrainingDateTx(ctx, tx, emplID, trainingID, date)
		if txErr != nil {
			return errors.WithMessage(txErr, "set training date")
		}

		if txErr := s.tasksStorage.CloseTaskTx(ctx, tx, taskID); txErr != nil {
			return errors.WithMessage(txErr, "close task")
		}

		if needNextTask {
			if txErr := s.tasksStorage.AddTaskTx(ctx, tx, tasksStorage.TaskBaseInfo(*task)); txErr != nil {
				return errors.WithMessage(txErr, "save task")
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
