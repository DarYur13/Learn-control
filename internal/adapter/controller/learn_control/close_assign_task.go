package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/converter"
	pb "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) CloseAssignTask(ctx context.Context, req *pb.CloseAssignTaskRequest) (*emptypb.Empty, error) {
	taskID := int(req.GetTaskID())

	taskType, err := converter.PbTaskTypeToDomain(req.GetTaskType())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	if err := i.service.CloseAssignTask(ctx, taskID, taskType); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
