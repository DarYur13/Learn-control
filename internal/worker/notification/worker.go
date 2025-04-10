package worker

import (
	"time"

	docsgenerator "github.com/DarYur13/learn-control/internal/adapter/docs_generator/registration_form"
	"github.com/DarYur13/learn-control/internal/adapter/notifier/email"
	emplRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/employees"
	notificationsRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/notifications"
	trainingsRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/trainings"
)

type notificationWorker struct {
	employeesRepo     emplRepo.EmployeesRepository
	trainingsRepo     trainingsRepo.TrainingsRepository
	notificationsRepo notificationsRepo.NotificationsRepository
	notifier          email.Notifier
	docsGenerator     docsgenerator.DocsGenerator
	interval          time.Duration
}

func New(
	employeesRepo emplRepo.EmployeesRepository,
	trainingsRepo trainingsRepo.TrainingsRepository,
	notificationsRepo notificationsRepo.NotificationsRepository,
	docsGenerator docsgenerator.DocsGenerator,
	notifier email.Notifier,
	interval time.Duration,
) NotificationWorker {
	return &notificationWorker{
		employeesRepo:     employeesRepo,
		trainingsRepo:     trainingsRepo,
		notificationsRepo: notificationsRepo,
		notifier:          notifier,
		docsGenerator:     docsGenerator,
		interval:          interval,
	}
}
