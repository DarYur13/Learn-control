package service

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

func (s *Service) GetEmployeesByName(ctx context.Context, name string) ([]domain.EmployeeBaseInfo, error) {
	employeesBaseInfo, err := s.employeesStorage.GetEmployeesByName(ctx, name)
	if err != nil {
		return nil, err
	}

	result := make([]domain.EmployeeBaseInfo, 0, len(employeesBaseInfo))

	for _, employeeBaseInfo := range employeesBaseInfo {
		result = append(result, domain.EmployeeBaseInfo(employeeBaseInfo))
	}

	return result, nil
}
