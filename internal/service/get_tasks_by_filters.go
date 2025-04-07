package service

import (
	"context"
	"database/sql"

	"github.com/DarYur13/learn-control/internal/domain"
)

func (s *Service) GetTasksByFilters(ctx context.Context, done sql.NullBool) ([]domain.Task, error) {
	tasks, err := s.tasksStorage.GetTasksByFilters(ctx, done)
	if err != nil {
		return nil, err
	}

	result := make([]domain.Task, 0, len(tasks))

	for _, task := range tasks {
		t := domain.Task(task)

		result = append(result, t)
	}

	return result, nil
}
