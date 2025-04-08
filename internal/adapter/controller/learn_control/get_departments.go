package learncontrol

import (
	"context"

	pb "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) GetDepartments(ctx context.Context, _ *emptypb.Empty) (*pb.GetDepartmentsResponse, error) {
	departments, err := i.service.GetDepartments(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.GetDepartmentsResponse{Departments: departments}, nil
}
