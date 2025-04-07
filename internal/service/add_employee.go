package service

import (
	"context"
	"database/sql"

	emplStorage "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/employees"
	tasksStorage "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/tasks"
	"github.com/DarYur13/learn-control/internal/domain"
	"github.com/pkg/errors"
)

func (s *Service) AddEmployee(ctx context.Context, employee domain.Employee) error {
	trainingsIDs, err := s.positionsStorage.GetPositionTrainings(ctx, employee.Department, employee.Position)
	if err != nil {
		return err
	}

	if err := s.txManager.Do(ctx, func(tx *sql.Tx) error {
		employeeID, txErr := s.employeesStorage.AddEmployeeTx(ctx, tx, emplStorage.Employee(employee))
		if txErr != nil {
			return errors.WithMessage(txErr, "add employee")
		}

		if trainingsIDs != nil {
			if txErr := s.employeesStorage.SetEmployeeTrainingsTx(ctx, tx, employeeID, trainingsIDs); txErr != nil {
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

				if txErr := s.tasksStorage.AddTaskTx(ctx, tx, tasksStorage.TaskBaseInfo(*task)); txErr != nil {
					return errors.WithMessage(txErr, "add task")
				}
			}

		} else {
			positionID, txErr := s.positionsStorage.AddPositionTx(ctx, tx, employee.Position, employee.Department)
			if txErr != nil {
				return errors.WithMessage(txErr, "set employee trainings")
			}

			task, txErr := s.createChooseTask(ctx, positionID)
			if txErr != nil {
				return errors.WithMessage(txErr, "create choose task")
			}

			if txErr := s.tasksStorage.AddTaskTx(ctx, tx, tasksStorage.TaskBaseInfo(*task)); txErr != nil {
				return errors.WithMessage(txErr, "add assign task")
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
