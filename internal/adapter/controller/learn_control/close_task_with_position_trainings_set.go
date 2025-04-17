package learncontrol

import (
	"context"

	pb "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) CloseTaskWithPositionTrainingsSet(ctx context.Context, req *pb.CloseTaskWithPositionTrainingsSetRequest) (*emptypb.Empty, error) {
	taskID := int(req.GetTaskID())
	trainingsIDs := make([]int, 0, len(req.GetTrainingsIDs()))

	for _, trainingID := range req.GetTrainingsIDs() {
		trainingsIDs = append(trainingsIDs, int(trainingID))
	}

	if err := i.service.CloseTaskWithPositionTrainingsSet(ctx, taskID, trainingsIDs); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
