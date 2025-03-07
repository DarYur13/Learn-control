package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/converter"
	desc "github.com/DarYur13/learn-control/pkg/learn_control"
)

func (i *Implementation) GetEmployeesByFilters(ctx context.Context, req *desc.GetEmployeesByFiltersRequest) (*desc.GetEmployeesByFiltersResponse, error) {
	filters := converter.FiltersToDomain(req)

	employees, err := i.learnControlSrv.GetEmployeesByFilters(ctx, filters)
	if err != nil {
		return nil, err
	}

	return converter.EmployeesInfoToDesc(employees), nil
}
