package learncontrol

import (
	"context"

	desc "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetEmployee(ctx context.Context, req *desc.GetEmployeeRequest) (*desc.GetEmployeeResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	employee, err := i.learnControlSrv.GetEmployee(ctx, int(req.GetId()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal service error: %s", err.Error())
	}

	return employee.ToDesc(), nil
}
