package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/converter"
	pb "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) CloseTaskWithTrainingDateSet(ctx context.Context, req *pb.CloseTaskWithTrainingDateSetRequest) (*emptypb.Empty, error) {
	emplID := int(req.GetEmployeeID())
	trainingID := int(req.GetTrainingID())
	taskID := int(req.GetTaskID())
	date := req.GetDate().AsTime()

	taskType, err := converter.PbTaskTypeToDomain(req.GetTaskType())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	if err := i.service.CloseTaskWithTrainingDateSet(ctx, taskID, emplID, trainingID, taskType, date); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
