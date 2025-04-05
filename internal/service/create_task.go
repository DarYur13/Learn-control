package service

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

func (s *Service) createProvideTask(ctx context.Context, employeeID, trainingID int) (*domain.TaskBaseInfo, error) {
	ids := map[string]int{
		"employee": employeeID,
		"training": trainingID,
	}
	return s.createTask(ctx, "PROVIDE", ids)
}

func (s *Service) createAssignTask(ctx context.Context, employeeID, trainingID int) (*domain.TaskBaseInfo, error) {
	ids := map[string]int{
		"employee": employeeID,
		"training": trainingID,
	}
	return s.createTask(ctx, "ASSIGN", ids)
}

func (s *Service) createChooseTask(ctx context.Context, positionID int) (*domain.TaskBaseInfo, error) {
	ids := map[string]int{
		"position": positionID,
	}
	return s.createTask(ctx, "CHOOSE", ids)
}

func (s *Service) createSetTask(ctx context.Context, employeeID, trainingID int) (*domain.TaskBaseInfo, error) {
	ids := map[string]int{
		"employee": employeeID,
		"training": trainingID,
	}
	return s.createTask(ctx, "SET", ids)
}

func (s *Service) createConfirmTask(ctx context.Context, employeeID, trainingID int) (*domain.TaskBaseInfo, error) {
	ids := map[string]int{
		"employee": employeeID,
		"training": trainingID,
	}
	return s.createTask(ctx, "CONFIRM", ids)
}

func (s *Service) createTask(_ context.Context, taskType string, ids map[string]int) (*domain.TaskBaseInfo, error) {
	task := &domain.TaskBaseInfo{
		Type: taskType,
	}

	for key, val := range ids {
		switch key {
		case "employee":
			task.EmployeeID.Valid = true
			task.EmployeeID.Int64 = int64(val)
		case "executor":
			task.ExecutorID.Valid = true
			task.ExecutorID.Int64 = int64(val)
		case "training":
			task.TrainingID.Valid = true
			task.TrainingID.Int64 = int64(val)
		case "position":
			task.PositionID.Valid = true
			task.PositionID.Int64 = int64(val)
		}

	}

	return task, nil
}
