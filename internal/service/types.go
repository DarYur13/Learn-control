package service

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

// ILearnControlService is the interface of the service layer
type ILearnControlService interface {
	GetEmployees(ctx context.Context, name string) (*domain.EmployeesBaseInfo, error)
	GetEmployee(ctx context.Context, id int) (*domain.Employee, error)
}
