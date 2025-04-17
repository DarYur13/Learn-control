package service

import (
	"context"
	"database/sql"
	"io"
	"time"

	"github.com/DarYur13/learn-control/internal/domain"
	"github.com/google/uuid"
)

type Servicer interface {
	GetEmployeeByID(ctx context.Context, employeeID int) (*domain.Employee, error)
	GetEmployeesByName(ctx context.Context, name string) ([]domain.EmployeeBaseInfo, error)
	GetEmployeePersonalCard(ctx context.Context, id int) (*domain.EmployeePersonalCard, error)
	GetEmployeesByFilters(ctx context.Context, filters domain.Filters) ([]domain.EmployeeInfo, error)
	UpdateEmployeeTrainingDate(ctx context.Context, employeeID int, trainingID int, date time.Time) (*domain.TrainingDates, error)
	AddEmployee(ctx context.Context, employee domain.Employee) error

	GetTrainings(ctx context.Context) ([]domain.TrainingBaseInfo, error)
	GetDepartments(ctx context.Context) ([]string, error)
	GetPositions(ctx context.Context) ([]string, error)

	GetTasksByFilters(ctx context.Context, done sql.NullBool) ([]domain.Task, error)
	CloseAssignTask(ctx context.Context, taskID int, taskType domain.TaskType) error
	CloseTaskWithPositionTrainingsSet(ctx context.Context, taskID int, trainingsIDs []int) error
	CloseTaskWithTrainingProtocolConfirm(ctx context.Context, taskID int) error
	CloseTaskWithTrainingDateSet(ctx context.Context, taskID int, taskType domain.TaskType, date time.Time) error

	GetFileByToken(ctx context.Context, token uuid.UUID) (io.Reader, error)

	CreateProvideTask(ctx context.Context, employeeID, trainingID, positionID int) (*domain.TaskBaseInfo, error)
	CreateAssignTask(ctx context.Context, employeeID, trainingID, positionID int) (*domain.TaskBaseInfo, error)
	CreateSetTask(ctx context.Context, employeeID, trainingID, positionID int) (*domain.TaskBaseInfo, error)
	CreateConfirmTask(ctx context.Context, employeeID, trainingID, positionID int) (*domain.TaskBaseInfo, error)
	CreateControlTask(ctx context.Context, employeeID, trainingID, executorID, positionID int) (*domain.TaskBaseInfo, error)
	CreateChooseTask(ctx context.Context, positionID int) (*domain.TaskBaseInfo, error)
}

type taskArgs struct {
	EmployeeID *int
	ExecutorID *int
	TrainingID *int
	PositionID *int
}
