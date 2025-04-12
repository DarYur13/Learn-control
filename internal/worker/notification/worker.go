package worker

import (
	"time"

	docsgenerator "github.com/DarYur13/learn-control/internal/adapter/docs_generator/registration_form"
	notifier "github.com/DarYur13/learn-control/internal/adapter/notifier/email"
	downloadTokensRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/download_tokens"
	emplRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/employees"
	notificationsRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/notifications"
	trainingsRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/trainings"
)

type notificationWorker struct {
	interval           time.Duration
	employeesRepo      emplRepo.EmployeesRepository
	trainingsRepo      trainingsRepo.TrainingsRepository
	notificationsRepo  notificationsRepo.NotificationsRepository
	downloadTokensRepo downloadTokensRepo.DownloadTokensRepository
	notifier           notifier.Notifier
	docsGenerator      docsgenerator.DocsGenerator
}

func New(
	employeesRepo emplRepo.EmployeesRepository,
	trainingsRepo trainingsRepo.TrainingsRepository,
	notificationsRepo notificationsRepo.NotificationsRepository,
	downloadTokensRepo downloadTokensRepo.DownloadTokensRepository,
	docsGenerator docsgenerator.DocsGenerator,
	notifier notifier.Notifier,
	interval time.Duration,
) NotificationWorker {
	return &notificationWorker{
		employeesRepo:      employeesRepo,
		trainingsRepo:      trainingsRepo,
		notificationsRepo:  notificationsRepo,
		downloadTokensRepo: downloadTokensRepo,
		notifier:           notifier,
		docsGenerator:      docsGenerator,
		interval:           interval,
	}
}
