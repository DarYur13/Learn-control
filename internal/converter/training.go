package converter

import (
	"github.com/DarYur13/learn-control/internal/domain"
	pb "github.com/DarYur13/learn-control/pkg/learn_control"
)

func TrainingToPb(t domain.Training) *pb.Training {
	return &pb.Training{
		Name:        t.Name,
		PassDate:    t.PassDate,
		RePassDate:  t.RePassDate,
		HasProtocol: t.HasProtocol,
	}
}

func TrainingsToPb(trainings []domain.TrainingBaseInfo) *pb.GetTrainingsResponse {
	result := &pb.GetTrainingsResponse{}

	for _, t := range trainings {
		training := &pb.TrainingBaseInfo{
			Id:   int64(t.ID),
			Name: t.Name,
		}

		result.Trainings = append(result.Trainings, training)
	}

	return result
}

func TrainingDatesToPb(dates *domain.TrainingDates) *pb.UpdateEmployeeTrainingDateResponse {
	return &pb.UpdateEmployeeTrainingDateResponse{
		PassDate:   dates.PassDate,
		RePassDate: dates.RePassDate,
	}
}
