package service

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

const dateFormat = "02.01.2006"

// ILearnControlService is the interface of the service layer
type ILearnControlService interface {
	GetEmployeesByName(ctx context.Context, name string) ([]domain.EmployeeBaseInfo, error)
	GetEmployee(ctx context.Context, id int) (*domain.EmployeePersonalCard, error)
	GetFilters(ctx context.Context) (*domain.Filters, error)
}
