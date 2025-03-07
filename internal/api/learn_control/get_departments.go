package learncontrol

import (
	"context"

	desc "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) GetDepartments(ctx context.Context, _ *emptypb.Empty) (*desc.GetDepartmentsResponse, error) {
	departments, err := i.learnControlSrv.GetDepartments(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.GetDepartmentsResponse{Departments: departments}, nil
}
