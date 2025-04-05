package service

import (
	"context"
	"database/sql"

	"github.com/DarYur13/learn-control/internal/domain"
	storage "github.com/DarYur13/learn-control/internal/storage/learn_control"
	"github.com/pkg/errors"
)

func (s *Service) CloseTaskWithPositionTrainingsSet(ctx context.Context, taskID, positionID int, trainingsIDs []int) error {
	if err := s.txManager.Do(ctx, func(tx *sql.Tx) error {
		if txErr := s.storage.SetPositionTrainingsTx(ctx, tx, positionID, trainingsIDs); txErr != nil {
			return errors.WithMessage(txErr, "set position trainings")
		}

		if txErr := s.storage.CloseTaskTx(ctx, tx, taskID); txErr != nil {
			return errors.WithMessage(txErr, "close task")
		}

		employeesIDs, txErr := s.storage.GetEmployeesWithoutTrainingsTx(ctx, tx, positionID)
		if txErr != nil {
			return errors.WithMessage(txErr, "get employees without trainings")
		}

		for _, emplID := range employeesIDs {
			if txErr := s.storage.SetEmployeeTrainingsTx(ctx, tx, emplID, trainingsIDs); txErr != nil {
				return errors.WithMessage(txErr, "set employee trainings")
			}

			for _, trainingID := range trainingsIDs {
				var task *domain.TaskBaseInfo

				if trainingID == 2 {
					task, txErr = s.createProvideTask(ctx, emplID, trainingID)
					if txErr != nil {
						return errors.WithMessage(txErr, "create provide task")
					}
				} else {
					task, txErr = s.createAssignTask(ctx, emplID, trainingID)
					if txErr != nil {
						return errors.WithMessage(txErr, "create assign task")
					}
				}

				if txErr := s.storage.AddTaskTx(ctx, tx, storage.TaskBaseInfo(*task)); txErr != nil {
					return errors.WithMessage(txErr, "add task")
				}
			}
		}

		return nil

	}); err != nil {
		return err
	}

	return nil
}
