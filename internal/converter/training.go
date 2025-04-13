package converter

import (
	"github.com/DarYur13/learn-control/internal/domain"
	pb "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TrainingToPb(t domain.Training) *pb.Training {
	return &pb.Training{
		Name:        t.Name,
		Type:        DomainTriningTypeToPb(t.Type),
		PassDate:    timestamppb.New(t.PassDate),
		RePassDate:  timestamppb.New(t.RePassDate),
		HasProtocol: t.HasProtocol,
	}
}

func TrainingsToPb(trainings []domain.TrainingBaseInfo) *pb.GetTrainingsResponse {
	result := &pb.GetTrainingsResponse{}

	for _, t := range trainings {
		training := &pb.TrainingBaseInfo{
			Id:   int64(t.ID),
			Name: t.Name,
			Type: DomainTriningTypeToPb(t.Type),
		}

		result.Trainings = append(result.Trainings, training)
	}

	return result
}

func TrainingDatesToPb(dates *domain.TrainingDates) *pb.UpdateEmployeeTrainingDateResponse {
	return &pb.UpdateEmployeeTrainingDateResponse{
		PassDate:   timestamppb.New(dates.PassDate),
		RePassDate: timestamppb.New(dates.RePassDate),
	}
}

func DomainTriningTypeToPb(trainingType domain.TrainingType) pb.TrainingType {
	if value, found := pb.TrainingType_value[string(trainingType)]; found {
		return pb.TrainingType(value)
	}

	return pb.TrainingType_UNKNOWN_TRAINING
}
