package service

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

func (s *Service) createProvideTask(ctx context.Context, employeeID, trainingID int) (*domain.TaskBaseInfo, error) {
	taskArgs := taskArgs{
		EmployeeID: &employeeID,
		TrainingID: &trainingID,
	}
	return s.buildTask(ctx, domain.TaskTypeProvide, taskArgs)
}

func (s *Service) createAssignTask(ctx context.Context, employeeID, trainingID int) (*domain.TaskBaseInfo, error) {
	taskArgs := taskArgs{
		EmployeeID: &employeeID,
		TrainingID: &trainingID,
	}
	return s.buildTask(ctx, domain.TaskTypeAssign, taskArgs)
}

func (s *Service) createChooseTask(ctx context.Context, positionID int) (*domain.TaskBaseInfo, error) {
	taskArgs := taskArgs{
		PositionID: &positionID,
	}
	return s.buildTask(ctx, domain.TaskTypeChoose, taskArgs)
}

func (s *Service) createSetTask(ctx context.Context, employeeID, trainingID int) (*domain.TaskBaseInfo, error) {
	taskArgs := taskArgs{
		EmployeeID: &employeeID,
		TrainingID: &trainingID,
	}
	return s.buildTask(ctx, domain.TaskTypeSet, taskArgs)
}

func (s *Service) createConfirmTask(ctx context.Context, employeeID, trainingID int) (*domain.TaskBaseInfo, error) {
	taskArgs := taskArgs{
		EmployeeID: &employeeID,
		TrainingID: &trainingID,
	}
	return s.buildTask(ctx, domain.TaskTypeConfirm, taskArgs)
}

func (s *Service) createControlTask(ctx context.Context, employeeID, trainingID, executorID int) (*domain.TaskBaseInfo, error) {
	taskArgs := taskArgs{
		EmployeeID: &employeeID,
		TrainingID: &trainingID,
		ExecutorID: &executorID,
	}
	return s.buildTask(ctx, domain.TaskTypeControl, taskArgs)
}

func (s *Service) buildTask(_ context.Context, taskType domain.TaskType, args taskArgs) (*domain.TaskBaseInfo, error) {
	task := &domain.TaskBaseInfo{
		Type: taskType,
	}

	if args.EmployeeID != nil {
		task.EmployeeID.Valid = true
		task.EmployeeID.Int64 = int64(*args.EmployeeID)
	}
	if args.ExecutorID != nil {
		task.ExecutorID.Valid = true
		task.ExecutorID.Int64 = int64(*args.ExecutorID)
	}
	if args.TrainingID != nil {
		task.TrainingID.Valid = true
		task.TrainingID.Int64 = int64(*args.TrainingID)
	}
	if args.PositionID != nil {
		task.PositionID.Valid = true
		task.PositionID.Int64 = int64(*args.PositionID)
	}

	return task, nil
}
