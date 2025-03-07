package service

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
	"github.com/DarYur13/learn-control/internal/storage"
)

func (s *Service) GetEmployeePersonalCard(ctx context.Context, id int) (*domain.EmployeePersonalCard, error) {
	employee, err := s.storage.GetEmployeePersonalCard(ctx, id)
	if err != nil {
		return nil, err
	}

	result := domain.EmployeePersonalCard{
		FullName:       employee.FullName,
		BirthDate:      employee.BirthDate,
		Snils:          employee.Snils,
		Department:     employee.Department,
		Position:       employee.Position,
		EmploymentDate: employee.EmploymentDate,
	}

	for _, t := range employee.Trainings {
		training := domain.Training{
			Name:          t.Name,
			TrainingDates: s.formatTrainingDates(t.TrainingDates),
		}

		result.Trainings = append(result.Trainings, training)
	}

	return &result, nil
}

func (s *Service) formatTrainingDates(st storage.TrainingDates) domain.TrainingDates {
	dt := domain.TrainingDates{}
	if st.PassDate.Valid {
		dt.PassDate = st.PassDate.Time.Format(dateFormat)

		if st.RePassDate.Valid {
			dt.RePassDate = st.RePassDate.Time.Format(dateFormat)
		} else {
			dt.RePassDate = noNeedRepassDate
		}
	} else {
		dt.PassDate = noPassDate
		dt.RePassDate = noRepassDate
	}

	return dt
}
