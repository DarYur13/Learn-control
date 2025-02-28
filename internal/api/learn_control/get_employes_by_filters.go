package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/converter"
	"github.com/DarYur13/learn-control/internal/domain"
	desc "github.com/DarYur13/learn-control/pkg/learn_control"
)

func (i *Implementation) GetEmployeesByFilters(ctx context.Context, req *desc.GetEmployeesByFiltersRequest) (*desc.GetEmployeesByFiltersResponse, error) {
	filters := domain.Filters{
		Deparment:          req.GetDepartment(),
		Position:           req.GetPosition(),
		TrainingID:         int(req.GetTrainingID()),
		DateFrom:           req.GetDateFrom().AsTime(),
		DateTo:             req.DateTo.AsTime(),
		TrainingsNotPassed: req.GetTrainigsNotPassed(),
		RetrainingIn:       int(req.GetRetrainingIn()),
	}

	employees, err := i.learnControlSrv.GetEmployeesByFilters(ctx, filters)
	if err != nil {
		return nil, err
	}

	return converter.EmployeesInfoToDesc(employees), nil
}
