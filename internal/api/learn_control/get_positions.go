package learncontrol

import (
	"context"

	desc "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) GetPositions(ctx context.Context, _ *emptypb.Empty) (*desc.GetPositionsResponse, error) {
	positions, err := i.learnControlSrv.GetPositions(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.GetPositionsResponse{Positions: positions}, nil
}
