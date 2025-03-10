package converter

import (
	"github.com/DarYur13/learn-control/internal/domain"
	desc "github.com/DarYur13/learn-control/pkg/learn_control"
)

func TrainingToDesc(t domain.Training) *desc.Training {
	return &desc.Training{
		Name:        t.Name,
		PassDate:    t.PassDate,
		RePassDate:  t.RePassDate,
		HasProtocol: t.HasProtocol,
	}
}

func TrainingsToDesc(trainings []domain.TrainingBaseInfo) *desc.GetTrainingsResponse {
	result := &desc.GetTrainingsResponse{}

	for _, t := range trainings {
		training := &desc.TrainingBaseInfo{
			Id:   int64(t.ID),
			Name: t.Name,
		}

		result.Trainings = append(result.Trainings, training)
	}

	return result
}

func TrainingDatesToDesc(dates *domain.TrainingDates) *desc.UpdateEmployeeTrainingDateResponse {
	return &desc.UpdateEmployeeTrainingDateResponse{
		PassDate:   dates.PassDate,
		RePassDate: dates.RePassDate,
	}
}
