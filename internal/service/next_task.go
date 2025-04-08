package service

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

func (s *Service) nextTask(ctx context.Context, employeeID, trainingID int, taskType domain.TaskType) (*domain.TaskBaseInfo, bool, error) {
	var task *domain.TaskBaseInfo
	var err error

	switch taskType {
	case domain.TaskTypeAssign:
		task, err = s.createSetTask(ctx, employeeID, trainingID)
		if err != nil {
			return nil, false, err
		}

	case domain.TaskTypeSet:
		task, err = s.createConfirmTask(ctx, employeeID, trainingID)
		if err != nil {
			return nil, false, err
		}

	default:
		return nil, false, nil
	}

	return task, true, nil
}
