package service

import (
	"context"
	"database/sql"

	"github.com/DarYur13/learn-control/internal/domain"
	storage "github.com/DarYur13/learn-control/internal/storage/learn_control"
	"github.com/pkg/errors"
)

func (s *Service) AddEmployee(ctx context.Context, employee domain.Employee) error {
	trainingsIDs, err := s.storage.GetTrainingsForPosition(ctx, employee.Department, employee.Position)
	if err != nil {
		return err
	}

	if err := s.txManager.Do(ctx, func(tx *sql.Tx) error {
		employeeID, txErr := s.storage.AddEmployeeTx(ctx, tx, storage.Employee(employee))
		if txErr != nil {
			return errors.WithMessage(txErr, "add employee")
		}

		if trainingsIDs != nil {
			if txErr := s.storage.SetEmployeeTrainingsTx(ctx, tx, employeeID, trainingsIDs); txErr != nil {
				return errors.WithMessage(txErr, "set employee trainings")
			}

			for _, trainingID := range trainingsIDs {
				var task *domain.TaskBaseInfo

				if trainingID == 2 {
					task, txErr = s.createProvideTask(ctx, employeeID, trainingID)
					if txErr != nil {
						return errors.WithMessage(txErr, "create provide task")
					}
				} else {
					task, txErr = s.createAssignTask(ctx, employeeID, trainingID)
					if txErr != nil {
						return errors.WithMessage(txErr, "create assign task")
					}
				}

				if txErr := s.storage.AddTaskTx(ctx, tx, storage.TaskBaseInfo(*task)); txErr != nil {
					return errors.WithMessage(txErr, "add task")
				}
			}

		} else {
			positionID, txErr := s.storage.AddPositionTx(ctx, tx, employee.Position, employee.Department)
			if txErr != nil {
				return errors.WithMessage(txErr, "set employee trainings")
			}

			task, txErr := s.createChooseTask(ctx, positionID)
			if txErr != nil {
				return errors.WithMessage(txErr, "create choose task")
			}

			if txErr := s.storage.AddTaskTx(ctx, tx, storage.TaskBaseInfo(*task)); txErr != nil {
				return errors.WithMessage(txErr, "add assign task")
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
