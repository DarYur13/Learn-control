package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
	pb "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) AddEmployee(ctx context.Context, req *pb.AddEmployeeRequest) (*emptypb.Empty, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	employee := domain.Employee{
		FullName:       req.GetFullname(),
		BirthDate:      req.GetBirthdate(),
		Snils:          req.GetSnils(),
		Department:     req.GetDepartment(),
		Position:       req.GetPosition(),
		EmploymentDate: req.GetEmploymentDate(),
	}

	if err := i.service.AddEmployee(ctx, employee); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
