package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/converter"
	pb "github.com/DarYur13/learn-control/pkg/learn_control"
)

func (i *Implementation) UpdateEmployeeTrainingDate(ctx context.Context, req *pb.UpdateEmployeeTrainingDateRequest) (*pb.UpdateEmployeeTrainingDateResponse, error) {
	employeeID := int(req.GetEmployeeID())
	trainingID := int(req.GetTrainingID())
	date := req.GetDate().AsTime()

	newDates, err := i.service.UpdateEmployeeTrainingDate(ctx, employeeID, trainingID, date)
	if err != nil {
		return nil, err
	}

	return converter.TrainingDatesToPb(newDates), nil
}
