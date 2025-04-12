package service

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

func (s *Service) GetEmployeeByID(ctx context.Context, employeeID int) (*domain.Employee, error) {
	storageEmpl, err := s.employeesStorage.GetEmployeeByID(ctx, employeeID)
	if err != nil {
		return nil, err
	}

	employee := &domain.Employee{
		FullName:       storageEmpl.FullName,
		BirthDate:      storageEmpl.BirthDate,
		Snils:          storageEmpl.Snils,
		Department:     storageEmpl.Department,
		Position:       storageEmpl.Position,
		EmploymentDate: storageEmpl.EmploymentDate,
		Email:          storageEmpl.Email,
	}

	return employee, nil
}
