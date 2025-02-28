package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/converter"
	desc "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) GetTrainings(ctx context.Context, _ *emptypb.Empty) (*desc.GetTrainingsResponse, error) {
	trainings, err := i.learnControlSrv.GetTrainings(ctx)
	if err != nil {
		return nil, err
	}

	return converter.TrainingsToDesc(trainings), nil
}
