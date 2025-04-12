package worker

import (
	"time"

	emplRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/employees"
	notificationsRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/notifications"
	tasksRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/tasks"
	trainingsRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/trainings"
	txManager "github.com/DarYur13/learn-control/internal/adapter/repository/txManager"
	"github.com/DarYur13/learn-control/internal/service"
)

type retrainingControlWorker struct {
	txManager         *txManager.Manager
	employeesRepo     emplRepo.EmployeesRepository
	trainingsRepo     trainingsRepo.TrainingsRepository
	notificationsRepo notificationsRepo.NotificationsRepository
	tasksRepo         tasksRepo.TasksRepository
	service           service.Servicer
	interval          time.Duration
}

func New(
	interval time.Duration,
	txManager *txManager.Manager,
	employeesRepo emplRepo.EmployeesRepository,
	trainingsRepo trainingsRepo.TrainingsRepository,
	notificationsRepo notificationsRepo.NotificationsRepository,
	tasksRepo tasksRepo.TasksRepository,
	service service.Servicer,
) RetrainingControlWorker {
	return &retrainingControlWorker{
		service:           service,
		txManager:         txManager,
		employeesRepo:     employeesRepo,
		trainingsRepo:     trainingsRepo,
		notificationsRepo: notificationsRepo,
		tasksRepo:         tasksRepo,
		interval:          interval,
	}
}
