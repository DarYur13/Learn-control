package service

import (
	"context"
	"time"

	"github.com/DarYur13/learn-control/internal/domain"
)

func (s *Service) UpdateEmployeeTrainingDate(ctx context.Context, employeeID int, trainingID int, date time.Time) (*domain.TrainingDates, error) {
	dates, err := s.storage.UpdateEmployeeTrainingDate(ctx, employeeID, trainingID, date)
	if err != nil {
		return nil, err
	}

	result := formatTrainingDates(*dates)

	return &result, nil
}
