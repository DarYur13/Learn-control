package learncontrol

import (
	"context"
	"fmt"

	"github.com/DarYur13/learn-control/internal/domain"
	desc "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) AddEmployee(ctx context.Context, req *desc.AddEmployeeRequest) (*emptypb.Empty, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	employee := domain.Employee{
		FullName:       req.GetFullname(),
		BirthDate:      req.GetBirthdate().AsTime().Format("02.01.2006"),
		Snils:          req.GetSnils(),
		Department:     req.GetDepartment(),
		Position:       req.GetPosition(),
		EmploymentDate: req.GetEmploymentDate().AsTime().Format("02.01.2006"),
	}

	fmt.Println(employee)

	if err := i.learnControlSrv.AddEmployee(ctx, employee); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
