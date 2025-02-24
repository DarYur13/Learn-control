package service

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

func (s *Service) GetEmployee(ctx context.Context, id int) (*domain.Employee, error) {
	return s.storage.GetEmployee(ctx, id)
}
