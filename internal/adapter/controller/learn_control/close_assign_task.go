package learncontrol

import (
	"context"

	pb "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) CloseAssignTask(ctx context.Context, req *pb.CloseAssignTaskRequest) (*emptypb.Empty, error) {
	emplID := int(req.GetEmployeeID())
	trainingID := int(req.GetTrainingID())
	taskType := req.GetTaskType().String()
	taskID := int(req.GetTaskID())

	if err := i.service.CloseAssignTask(ctx, taskID, emplID, trainingID, taskType); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
