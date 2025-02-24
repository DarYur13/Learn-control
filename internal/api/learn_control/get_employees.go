package learncontrol

import (
	"context"

	desc "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetEmployees(ctx context.Context, req *desc.GetEmployeesRequest) (*desc.GetEmployeesResponse, error) {
	employees, err := i.learnControlSrv.GetEmployees(ctx, req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal service error: %s", err.Error())
	}

	return employees.ToDesc(), nil
}
