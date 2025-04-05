package learncontrol

import (
	"context"

	desc "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) CloseTaskWithTrainingDateSet(ctx context.Context, req *desc.CloseTaskWithTrainingDateSetRequest) (*emptypb.Empty, error) {
	emplID := int(req.GetEmployeeID())
	trainingID := int(req.GetTrainingID())
	taskType := req.GetTaskType().String()
	taskID := int(req.GetTaskID())
	date := req.GetDate().AsTime()

	if err := i.learnControlSrv.CloseTaskWithTrainingDateSet(ctx, taskID, emplID, trainingID, taskType, date); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
