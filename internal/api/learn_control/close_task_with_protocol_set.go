package learncontrol

import (
	"context"

	desc "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) CloseTaskWithTrainingProtocolConfirm(ctx context.Context, req *desc.CloseTaskWithTrainingProtocolConfirmRequest) (*emptypb.Empty, error) {
	taskID := int(req.GetTaskID())
	emplID := int(req.GetEmployeeID())
	trainingID := int(req.GetTrainingID())

	if err := i.learnControlSrv.CloseTaskWithTrainingProtocolConfirm(ctx, taskID, emplID, trainingID); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
