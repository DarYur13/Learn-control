package service

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
	"github.com/DarYur13/learn-control/internal/storage"
)

func (s *Service) GetEmployeesByFilters(ctx context.Context, filters domain.Filters) ([]domain.EmployeeInfo, error) {
	storageFilters := s.validateFilters(filters)

	employees, err := s.storage.GetEmployeesByFilters(ctx, storageFilters)
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
				Name: t.Name,
			}

			s.validateTrainingDates(t, &training)

			employee.Trainings = append(employee.Trainings, training)
		}

		result = append(result, employee)
	}

	return result, nil
}

func (s *Service) validateFilters(filters domain.Filters) storage.Filters {
	var storageFilters storage.Filters

	if filters.Deparment != "" {
		storageFilters.Department.Valid = true
		storageFilters.Department.String = filters.Deparment
	}

	if filters.Position != "" {
		storageFilters.Position.Valid = true
		storageFilters.Position.String = filters.Position
	}

	if filters.TrainingID > 0 {
		storageFilters.TrainingID.Valid = true
		storageFilters.TrainingID.Int64 = int64(filters.TrainingID)
	}

	if filters.RetrainingIn > 0 {
		storageFilters.RetrainingIn.Valid = true
		storageFilters.RetrainingIn.Int64 = int64(filters.RetrainingIn)
	}

	if filters.TrainingsNotPassed {
		storageFilters.TrainingsNotPassed.Valid = true
		storageFilters.TrainingsNotPassed.Bool = filters.TrainingsNotPassed
	}

	if !filters.DateFrom.IsZero() {
		storageFilters.DateFrom.Valid = true
		storageFilters.TrainingsNotPassed.Bool = filters.TrainingsNotPassed
	}

	if !filters.DateTo.IsZero() {
		storageFilters.DateTo.Valid = true
		storageFilters.DateTo.Time = filters.DateTo
	}

	return storageFilters
}
