package service

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

func (s *Service) GetTrainings(ctx context.Context) ([]domain.TrainingBaseInfo, error) {
	trainigs, err := s.storage.GetTrainings(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.TrainingBaseInfo, 0, len(trainigs))

	for _, t := range trainigs {
		training := domain.TrainingBaseInfo{
			ID:   t.ID,
			Name: t.Name,
		}

		result = append(result, training)
	}

	return result, nil
}
