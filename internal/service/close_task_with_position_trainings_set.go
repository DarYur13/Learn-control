package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

func (s *Service) CloseTaskWithPositionTrainingsSet(ctx context.Context, taskID, positionID int, trainingsIDs []int) error {
	return s.txManager.Do(ctx, func(tx *sql.Tx) error {
		if txErr := s.positionsStorage.SetPositionTrainingsTx(ctx, tx, positionID, trainingsIDs); txErr != nil {
			return errors.WithMessage(txErr, "set position trainings")
		}

		if txErr := s.tasksStorage.CloseTaskTx(ctx, tx, taskID); txErr != nil {
			return errors.WithMessage(txErr, "close task")
		}

		employeesIDs, txErr := s.employeesStorage.GetEmployeesWithoutTrainingsTx(ctx, tx, positionID)
		if txErr != nil {
			return errors.WithMessage(txErr, "get employees without trainings")
		}

		for _, emplID := range employeesIDs {
			if txErr := s.assignTrainingsAndTasks(ctx, tx, emplID, trainingsIDs); txErr != nil {
				return errors.WithMessagef(txErr, "assign trainings and tasks for employee %d", emplID)
			}
		}

		return nil
	})
}
