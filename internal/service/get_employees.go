package service

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

func (s *Service) GetEmployees(ctx context.Context, name string) (*domain.EmployeesBaseInfo, error) {
	return s.storage.GetEmployees(ctx, name)
}
