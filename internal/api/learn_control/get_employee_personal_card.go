package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/converter"
	desc "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetEmployeePersonalCard(ctx context.Context, req *desc.GetEmployeePersonalCardRequest) (*desc.GetEmployeePersonalCardResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	employee, err := i.learnControlSrv.GetEmployeePersonalCard(ctx, int(req.GetId()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal service error: %s", err.Error())
	}

	return converter.EmployeePersonalCardToDesc(employee), nil
}
