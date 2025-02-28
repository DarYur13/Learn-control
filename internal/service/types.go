package service

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

const dateFormat = "02.01.2006"

// ILearnControlService is the interface of the service layer
type ILearnControlService interface {
	GetEmployeesByName(ctx context.Context, name string) (*domain.EmployeesBaseInfo, error)
	GetEmployee(ctx context.Context, id int) (*domain.Employee, error)
}
