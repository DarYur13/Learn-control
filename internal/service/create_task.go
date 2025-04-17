package service

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

func (s *Service) CreateProvideTask(ctx context.Context, employeeID, trainingID, positionID int) (*domain.TaskBaseInfo, error) {
	taskArgs := taskArgs{
		EmployeeID: &employeeID,
		TrainingID: &trainingID,
		PositionID: &positionID,
	}
	return s.buildTask(ctx, domain.TaskTypeProvide, taskArgs)
}

func (s *Service) CreateAssignTask(ctx context.Context, employeeID, trainingID, positionID int) (*domain.TaskBaseInfo, error) {
	taskArgs := taskArgs{
		EmployeeID: &employeeID,
		TrainingID: &trainingID,
		PositionID: &positionID,
	}
	return s.buildTask(ctx, domain.TaskTypeAssign, taskArgs)
}

func (s *Service) CreateChooseTask(ctx context.Context, positionID int) (*domain.TaskBaseInfo, error) {
	taskArgs := taskArgs{
		PositionID: &positionID,
	}
	return s.buildTask(ctx, domain.TaskTypeChoose, taskArgs)
}

func (s *Service) CreateSetTask(ctx context.Context, employeeID, trainingID, positionID int) (*domain.TaskBaseInfo, error) {
	taskArgs := taskArgs{
		EmployeeID: &employeeID,
		TrainingID: &trainingID,
		PositionID: &positionID,
	}
	return s.buildTask(ctx, domain.TaskTypeSet, taskArgs)
}

func (s *Service) CreateConfirmTask(ctx context.Context, employeeID, trainingID, positionID int) (*domain.TaskBaseInfo, error) {
	taskArgs := taskArgs{
		EmployeeID: &employeeID,
		TrainingID: &trainingID,
		PositionID: &positionID,
	}
	return s.buildTask(ctx, domain.TaskTypeConfirm, taskArgs)
}

func (s *Service) CreateControlTask(ctx context.Context, employeeID, trainingID, executorID, positionID int) (*domain.TaskBaseInfo, error) {
	taskArgs := taskArgs{
		EmployeeID: &employeeID,
		TrainingID: &trainingID,
		ExecutorID: &executorID,
		PositionID: &positionID,
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
