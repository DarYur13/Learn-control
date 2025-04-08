package learncontrol

import (
	"context"

	pb "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) GetPositions(ctx context.Context, _ *emptypb.Empty) (*pb.GetPositionsResponse, error) {
	positions, err := i.service.GetPositions(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.GetPositionsResponse{Positions: positions}, nil
}
