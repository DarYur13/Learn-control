package service

import (
	emplRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/employees"
	posRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/positions"
	tasksRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/tasks"
	txManager "github.com/DarYur13/learn-control/internal/adapter/repository/txManager"
	"github.com/minio/minio-go/v7"
)

// Service
type Service struct {
	txManager        *txManager.Manager
	employeesStorage emplRepo.EmoloyeesStorager
	positionsStorage posRepo.PositionsStorager
	tasksStorage     tasksRepo.TasksStorager
	minioFileStor    *minio.Client
}

// New creates new service
func New(
	employeesStorage emplRepo.EmoloyeesStorager,
	positionsStorage posRepo.PositionsStorager,
	tasksStorage tasksRepo.TasksStorager,
	txManager *txManager.Manager,
	minioCli *minio.Client,
) *Service {
	return &Service{
		employeesStorage: employeesStorage,
		positionsStorage: positionsStorage,
		tasksStorage:     tasksStorage,
		txManager:        txManager,
		minioFileStor:    minioCli,
	}
}
