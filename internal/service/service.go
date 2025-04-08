package service

import (
	emplRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/employees"
	posRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/positions"
	tasksRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/tasks"
	trainingsRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/trainings"
	txManager "github.com/DarYur13/learn-control/internal/adapter/repository/txManager"
	"github.com/minio/minio-go/v7"
)

// Service
type Service struct {
	txManager        *txManager.Manager
	employeesStorage emplRepo.EmployeesRepository
	positionsStorage posRepo.PositionsRepository
	tasksStorage     tasksRepo.TasksRepository
	trainingsStorage trainingsRepo.TrainingsRepository
	minioFileStor    *minio.Client
}

// New creates new service
func New(
	employeesStorage emplRepo.EmployeesRepository,
	positionsStorage posRepo.PositionsRepository,
	tasksStorage tasksRepo.TasksRepository,
	trainingsStorage trainingsRepo.TrainingsRepository,
	txManager *txManager.Manager,
	minioCli *minio.Client,
) *Service {
	return &Service{
		employeesStorage: employeesStorage,
		positionsStorage: positionsStorage,
		tasksStorage:     tasksStorage,
		trainingsStorage: trainingsStorage,
		txManager:        txManager,
		minioFileStor:    minioCli,
	}
}
