package service

import (
	"context"

	emplStorage "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/employees"
	"github.com/DarYur13/learn-control/internal/domain"
)

func (s *Service) GetEmployeesByFilters(ctx context.Context, filters domain.Filters) ([]domain.EmployeeInfo, error) {

	employees, err := s.employeesStorage.GetEmployeesByFilters(ctx, emplStorage.Filters(filters))
	if err != nil {
		return nil, err
	}

	result := make([]domain.EmployeeInfo, 0, len(employees))

	for _, e := range employees {
		employee := domain.EmployeeInfo{
			FullName:   e.FullName,
			Department: e.Department,
			Position:   e.Position,
		}

		for _, t := range e.Trainings {
			training := domain.Training{
				Name:          t.Name,
				TrainingDates: formatTrainingDates(t.TrainingDates),
				HasProtocol:   formatTrainingHasProtocol(t.HasProtocol),
			}

			employee.Trainings = append(employee.Trainings, training)
		}

		result = append(result, employee)
	}

	return result, nil
}
