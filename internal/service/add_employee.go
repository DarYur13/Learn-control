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
	positionID, trainingsIDs, err := s.positionsStorage.GetPositionTrainings(ctx, employee.Department, employee.Position)
	if err != nil {
		return err
	}

	return s.txManager.Do(ctx, func(tx *sql.Tx) error {
		employeeID, err := s.employeesStorage.AddEmployeeTx(ctx, tx, emplStorage.Employee(employee))
		if err != nil {
			return errors.WithMessage(err, "add employee")
		}

		if len(trainingsIDs) > 0 {
			return s.assignTrainingsAndTasks(ctx, tx, employeeID, positionID, trainingsIDs)
		}

		return s.addPositionAndTask(ctx, tx, employee)
	})
}

func (s *Service) assignTrainingsAndTasks(ctx context.Context, tx *sql.Tx, employeeID, positionID int, trainingIDs []int) error {
	if err := s.employeesStorage.SetEmployeeTrainingsTx(ctx, tx, employeeID, trainingIDs); err != nil {
		return errors.WithMessage(err, "set employee trainings")
	}

	for _, trainingID := range trainingIDs {
		trainingType, err := s.trainingsStorage.GetTrainingType(ctx, trainingID)
		if err != nil {
			return errors.WithMessage(err, "get training type")
		}

		task, err := s.buildTaskForTraining(ctx, employeeID, trainingID, positionID, trainingType)
		if err != nil {
			return errors.WithMessagef(err, "build task for training %d", trainingID)
		}

		if err := s.tasksStorage.AddTaskTx(ctx, tx, tasksStorage.TaskBaseInfo(*task)); err != nil {
			return errors.WithMessage(err, "save task")
		}
	}

	return nil
}

func (s *Service) buildTaskForTraining(ctx context.Context, employeeID, trainingID, positionID int, trainingType domain.TrainingType) (*domain.TaskBaseInfo, error) {
	switch trainingType {
	case domain.TrainingTypeIntroductory:
		return s.CreateProvideTask(ctx, employeeID, trainingID, positionID)

	case domain.TrainingTypeInitial:
		executorID, err := s.employeesStorage.GetEmployeeLeader(ctx, employeeID)
		if err != nil {
			return nil, errors.WithMessage(err, "get department leader")
		}

		if err := s.notificationsStorage.AddNotificationToQueue(
			ctx,
			employeeID,
			trainingID,
			domain.NotificationTypeInitBrief,
		); err != nil {
			return nil, errors.WithMessage(err, "enqueue init brief")
		}

		return s.CreateControlTask(ctx, employeeID, trainingID, executorID, positionID)

	case domain.TrainingTypeRefresher:
		executorID, err := s.employeesStorage.GetEmployeeLeader(ctx, employeeID)
		if err != nil {
			return nil, errors.WithMessage(err, "get department leader")
		}

		return s.CreateControlTask(ctx, employeeID, trainingID, executorID, positionID)

	default:
		return s.CreateAssignTask(ctx, employeeID, trainingID, positionID)
	}
}

func (s *Service) addPositionAndTask(ctx context.Context, tx *sql.Tx, employee domain.Employee) error {
	positionID, err := s.positionsStorage.AddPositionTx(ctx, tx, employee.Position, employee.Department)
	if err != nil {
		return errors.WithMessage(err, "create new position")
	}

	task, err := s.CreateChooseTask(ctx, positionID)
	if err != nil {
		return errors.WithMessage(err, "create choose task")
	}

	if err := s.tasksStorage.AddTaskTx(ctx, tx, tasksStorage.TaskBaseInfo(*task)); err != nil {
		return errors.WithMessage(err, "save choose task")
	}

	return nil
}
