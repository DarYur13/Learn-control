package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/converter"
	pb "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) AddEmployee(ctx context.Context, req *pb.AddEmployeeRequest) (*emptypb.Empty, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	employee := converter.PbEmployeeToDomain(req)

	if err := i.service.AddEmployee(ctx, employee); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
