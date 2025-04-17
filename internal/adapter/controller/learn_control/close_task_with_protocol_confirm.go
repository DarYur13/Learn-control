package learncontrol

import (
	"context"

	pb "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) CloseTaskWithTrainingProtocolConfirm(ctx context.Context, req *pb.CloseTaskWithTrainingProtocolConfirmRequest) (*emptypb.Empty, error) {
	taskID := int(req.GetTaskID())

	if err := i.service.CloseTaskWithTrainingProtocolConfirm(ctx, taskID); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
