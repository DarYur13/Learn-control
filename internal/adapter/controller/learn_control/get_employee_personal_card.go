package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/converter"
	pb "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetEmployeePersonalCard(ctx context.Context, req *pb.GetEmployeePersonalCardRequest) (*pb.GetEmployeePersonalCardResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	employee, err := i.service.GetEmployeePersonalCard(ctx, int(req.GetId()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal service error: %s", err.Error())
	}

	return converter.EmployeePersonalCardToPb(employee), nil
}
