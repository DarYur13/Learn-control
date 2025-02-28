package service

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

func (s *Service) GetEmployeesByName(ctx context.Context, name string) (*domain.EmployeesBaseInfo, error) {
	return s.storage.GetEmployeesByName(ctx, name)
}
