package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/converter"
	pb "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetEmployeesByName(ctx context.Context, req *pb.GetEmployeesByNameRequest) (*pb.GetEmployeesByNameResponse, error) {
	employees, err := i.service.GetEmployeesByName(ctx, req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal service error: %s", err.Error())
	}

	return converter.EmployeesBaseInfoToDesc(employees), nil
}
