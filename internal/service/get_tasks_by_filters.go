package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/DarYur13/learn-control/internal/config"
	"github.com/DarYur13/learn-control/internal/domain"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (s *Service) GetTasksByFilters(ctx context.Context, done sql.NullBool) ([]domain.Task, error) {
	tasks, err := s.tasksStorage.GetTasksByFilters(ctx, done)
	if err != nil {
		return nil, err
	}

	result := make([]domain.Task, 0, len(tasks))

	for _, t := range tasks {
		task := domain.Task{
			ID:          t.ID,
			Type:        t.Type,
			Description: t.Description,
			Employee:    t.Employee,
			Training:    t.Training,
			Position:    t.Position,
			Department:  t.Department,
			Executor:    t.Executor,
			Done:        t.Done,
		}

		if t.Type == domain.TaskTypeProvide {
			downloadFileLink, err := s.generateDownloadLink(ctx, int(t.EmployeeID.Int64), int(t.TrainingID.Int64))
			if err != nil {
				return nil, err
			}

			task.FileLink = sql.NullString{
				Valid:  true,
				String: downloadFileLink,
			}
		}

		result = append(result, task)
	}

	return result, nil
}

func (s *Service) generateDownloadLink(ctx context.Context, employeeID, trainingID int) (string, error) {
	token := uuid.New()
	expiresAt := time.Now().AddDate(0, 0, 1)

	if err := s.downloadTokensStorage.AddToken(ctx, domain.DownloadToken{
		Token:      token,
		EmployeeID: employeeID,
		TrainingID: trainingID,
		ExpiresAt:  expiresAt,
	}); err != nil {
		return "", errors.WithMessage(err, "failed to add new token")
	}

	return fmt.Sprintf("%s:%s/files/download?token=%s", config.ApiHost(), config.ApiHttpPort(), token.String()), nil
}
