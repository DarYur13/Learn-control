package service

import (
	"context"
	"database/sql"

	tasksStorage "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/tasks"
	"github.com/DarYur13/learn-control/internal/domain"
	"github.com/pkg/errors"
)

func (s *Service) CloseAssignTask(ctx context.Context, taskID int, taskType domain.TaskType) error {
	taskInfo, err := s.tasksStorage.GetTaskInfoByID(ctx, taskID)
	if err != nil {
		return errors.WithMessage(err, "get task info")
	}

	task, needNextTask, err := s.nextTask(ctx, int(taskInfo.EmployeeID.Int64), int(taskInfo.TrainingID.Int64), taskType)
	if err != nil {
		return errors.WithMessage(err, "create next task")
	}

	if err := s.txManager.Do(ctx, func(tx *sql.Tx) error {
		if txErr := s.tasksStorage.CloseTaskTx(ctx, tx, taskID); txErr != nil {
			return errors.WithMessage(txErr, "close task")
		}

		if needNextTask {
			if txErr := s.tasksStorage.AddTaskTx(ctx, tx, tasksStorage.TaskBaseInfo(*task)); txErr != nil {
				return errors.WithMessage(txErr, "add task")
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
