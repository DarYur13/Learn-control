package service

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

func (s *Service) GetEmployee(ctx context.Context, id int) (*domain.EmployeePersonalCard, error) {
	employee, err := s.storage.GetEmployee(ctx, id)
	if err != nil {
		return nil, err
	}

	result := domain.EmployeePersonalCard{
		FullName:   employee.FullName,
		BirthDate:  employee.BirthDate,
		Snils:      employee.Snils,
		Department: employee.Department,
		Position:   employee.Position,
	}

	var training domain.Training

	for _, t := range employee.Trainings {
		training.Name = t.Name

		if t.PassDate.Valid {
			training.PassDate = t.PassDate.Time.Format(dateFormat)

			if t.RePassDate.Valid {
				training.RePassDate = t.RePassDate.Time.Format(dateFormat)
			} else {
				training.RePassDate = "Не требуется"
			}

		} else {
			training.PassDate = "Обучение не пройдено"
			training.RePassDate = "Не установлена"
		}

		result.Trainings = append(result.Trainings, training)
	}

	return &result, nil
}
