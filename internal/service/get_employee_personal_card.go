package service

import (
	"context"
	"database/sql"

	trainingsStorage "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/trainings"
	"github.com/DarYur13/learn-control/internal/domain"
)

func (s *Service) GetEmployeePersonalCard(ctx context.Context, id int) (*domain.EmployeePersonalCard, error) {
	employee, err := s.employeesStorage.GetEmployeePersonalCard(ctx, id)
	if err != nil {
		return nil, err
	}

	result := domain.EmployeePersonalCard{
		EmployeeID:     employee.EmployeeID,
		FullName:       employee.FullName,
		BirthDate:      employee.BirthDate,
		Snils:          employee.Snils,
		Department:     employee.Department,
		Position:       employee.Position,
		EmploymentDate: employee.EmploymentDate,
	}

	for _, t := range employee.Trainings {
		training := domain.Training{
			ID:            t.ID,
			Name:          t.Name,
			TrainingDates: formatTrainingDates(t.TrainingDates),
			HasProtocol:   formatTrainingHasProtocol(t.HasProtocol),
		}

		result.Trainings = append(result.Trainings, training)
	}

	return &result, nil
}

func formatTrainingDates(st trainingsStorage.TrainingDates) domain.TrainingDates {
	dt := domain.TrainingDates{}
	if st.PassDate.Valid {
		dt.PassDate = st.PassDate.Time

		if st.RePassDate.Valid {
			dt.RePassDate = st.RePassDate.Time
		}
	}

	return dt
}

func formatTrainingHasProtocol(hp sql.NullBool) string {
	if hp.Valid {
		if hp.Bool {
			return "Получен"
		}

		return "Ожидается"
	}

	return "Не требуется"
}
