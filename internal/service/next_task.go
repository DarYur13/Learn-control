package service

import (
	"context"
	"fmt"

	"github.com/DarYur13/learn-control/internal/domain"
)

// TODO transaction
func (s *Service) nextTask(ctx context.Context, employeeID, trainingID int, taskType string) (*domain.TaskBaseInfo, bool, error) {
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
	case domain.TaskTypeProvide:
		return nil, false, nil
	default:
		return nil, false, fmt.Errorf("unknown task type")
	}

	return task, true, nil
}
