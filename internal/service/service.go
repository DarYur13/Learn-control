package service

import (
	docsgenerator "github.com/DarYur13/learn-control/internal/adapter/docs_generator/registration_form"
	notifier "github.com/DarYur13/learn-control/internal/adapter/notifier/email"
	downloadTokensRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/download_tokens"
	emplRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/employees"
	notificationsRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/notifications"
	posRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/positions"
	tasksRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/tasks"
	trainingsRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/trainings"
	txManager "github.com/DarYur13/learn-control/internal/adapter/repository/txManager"
)

// Service
type Service struct {
	txManager             *txManager.Manager
	employeesStorage      emplRepo.EmployeesRepository
	positionsStorage      posRepo.PositionsRepository
	tasksStorage          tasksRepo.TasksRepository
	trainingsStorage      trainingsRepo.TrainingsRepository
	notificationsStorage  notificationsRepo.NotificationsRepository
	docsGenerator         docsgenerator.DocsGenerator
	notifier              notifier.Notifier
	downloadTokensStorage downloadTokensRepo.DownloadTokensRepository
}

// New creates new service
func New(
	employeesStorage emplRepo.EmployeesRepository,
	positionsStorage posRepo.PositionsRepository,
	tasksStorage tasksRepo.TasksRepository,
	trainingsStorage trainingsRepo.TrainingsRepository,
	txManager *txManager.Manager,
	docsGenerator docsgenerator.DocsGenerator,
	notifier notifier.Notifier,
	notificationsStorage notificationsRepo.NotificationsRepository,
	downloadTokensStorage downloadTokensRepo.DownloadTokensRepository,
) *Service {
	return &Service{
		employeesStorage:      employeesStorage,
		positionsStorage:      positionsStorage,
		tasksStorage:          tasksStorage,
		trainingsStorage:      trainingsStorage,
		txManager:             txManager,
		docsGenerator:         docsGenerator,
		notifier:              notifier,
		notificationsStorage:  notificationsStorage,
		downloadTokensStorage: downloadTokensStorage,
	}
}
