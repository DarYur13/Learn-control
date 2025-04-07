package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/converter"
	pb "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) GetTrainings(ctx context.Context, _ *emptypb.Empty) (*pb.GetTrainingsResponse, error) {
	trainings, err := i.service.GetTrainings(ctx)
	if err != nil {
		return nil, err
	}

	return converter.TrainingsToDesc(trainings), nil
}
