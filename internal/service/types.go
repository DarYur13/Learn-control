package service

import (
	"context"
	"time"

	"github.com/DarYur13/learn-control/internal/domain"
)

const (
	dateFormat = "02.01.2006"

	noPassDate       = "Не пройдено"
	noRepassDate     = "Не установлено"
	noNeedRepassDate = "Не требуется"
)

// ILearnControlService is the interface of the service layer
type ILearnControlService interface {
	GetEmployeesByName(ctx context.Context, name string) ([]domain.EmployeeBaseInfo, error)
	GetEmployeePersonalCard(ctx context.Context, id int) (*domain.EmployeePersonalCard, error)
	GetEmployeesByFilters(ctx context.Context, filters domain.Filters) ([]domain.EmployeeInfo, error)
	GetTrainings(ctx context.Context) ([]domain.TrainingBaseInfo, error)
	GetDepartments(ctx context.Context) ([]string, error)
	GetPositions(ctx context.Context) ([]string, error)
	UpdateEmployeeTrainingDate(ctx context.Context, employeeID int, trainingID int, date time.Time) (*domain.TrainingDates, error)
	AddEmployee(ctx context.Context, employee domain.Employee) error
}
