package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/converter"
	pb "github.com/DarYur13/learn-control/pkg/learn_control"
)

func (i *Implementation) GetEmployeesByFilters(ctx context.Context, req *pb.GetEmployeesByFiltersRequest) (*pb.GetEmployeesByFiltersResponse, error) {
	filters := converter.FiltersToDomain(req)

	employees, err := i.service.GetEmployeesByFilters(ctx, filters)
	if err != nil {
		return nil, err
	}

	return converter.EmployeesInfoToDesc(employees), nil
}
