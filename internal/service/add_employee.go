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

	return s.txManager.Do(ctx, func(tx *sql.Tx) error {
		employeeID, err := s.employeesStorage.AddEmployeeTx(ctx, tx, emplStorage.Employee(employee))
		if err != nil {
			return errors.WithMessage(err, "add employee")
		}

		if len(trainingsIDs) > 0 {
			return s.assignTrainingsAndTasks(ctx, tx, employeeID, trainingsIDs)
		}

		return s.addPositionAndTask(ctx, tx, employee)
	})
}

func (s *Service) assignTrainingsAndTasks(ctx context.Context, tx *sql.Tx, employeeID int, trainingIDs []int) error {
	if err := s.employeesStorage.SetEmployeeTrainingsTx(ctx, tx, employeeID, trainingIDs); err != nil {
		return errors.WithMessage(err, "set employee trainings")
	}

	for _, trainingID := range trainingIDs {
		task, err := s.buildTaskForTraining(ctx, employeeID, trainingID)
		if err != nil {
			return errors.WithMessagef(err, "build task for training %d", trainingID)
		}

		if err := s.tasksStorage.AddTaskTx(ctx, tx, tasksStorage.TaskBaseInfo(*task)); err != nil {
			return errors.WithMessage(err, "save task")
		}
	}

	return nil
}

func (s *Service) buildTaskForTraining(ctx context.Context, employeeID int, trainingID int) (*domain.TaskBaseInfo, error) {
	switch trainingID {
	case domain.IntroductoryBriefingID:
		return s.createProvideTask(ctx, employeeID, trainingID)

	case domain.InitialBriefingID, domain.RefresherBriefingID:
		executorID, err := s.employeesStorage.GetEmployeeLeader(ctx, employeeID)
		if err != nil {
			return nil, errors.WithMessage(err, "get department leader")
		}

		return s.createControlTask(ctx, employeeID, trainingID, executorID)

	default:
		return s.createAssignTask(ctx, employeeID, trainingID)
	}
}

func (s *Service) addPositionAndTask(ctx context.Context, tx *sql.Tx, employee domain.Employee) error {
	positionID, err := s.positionsStorage.AddPositionTx(ctx, tx, employee.Position, employee.Department)
	if err != nil {
		return errors.WithMessage(err, "create new position")
	}

	task, err := s.createChooseTask(ctx, positionID)
	if err != nil {
		return errors.WithMessage(err, "create choose task")
	}

	if err := s.tasksStorage.AddTaskTx(ctx, tx, tasksStorage.TaskBaseInfo(*task)); err != nil {
		return errors.WithMessage(err, "save choose task")
	}

	return nil
}
