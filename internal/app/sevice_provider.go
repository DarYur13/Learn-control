package service_provider

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"

	impl "github.com/DarYur13/learn-control/internal/adapter/controller/learn_control"
	docsgenerator "github.com/DarYur13/learn-control/internal/adapter/docs_generator/registration_form"
	notifier "github.com/DarYur13/learn-control/internal/adapter/notifier/email"
	downloadTokensRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/download_tokens"
	emplRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/employees"
	notificationsRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/notifications"
	posRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/positions"
	tasksRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/tasks"
	trainingsRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/trainings"
	txManager "github.com/DarYur13/learn-control/internal/adapter/repository/txManager"
	"github.com/DarYur13/learn-control/internal/config"
	"github.com/DarYur13/learn-control/internal/logger"
	service "github.com/DarYur13/learn-control/internal/service"
)

// serviceProvider di-container
type serviceProvider struct {
	db                 *sql.DB
	txManager          *txManager.Manager
	EmployeesRepo      emplRepo.EmployeesRepository
	PositionsRepo      posRepo.PositionsRepository
	TasksRepo          tasksRepo.TasksRepository
	TrainingsRepo      trainingsRepo.TrainingsRepository
	service            service.Servicer
	implementation     *impl.Implementation
	docsGenerator      docsgenerator.DocsGenerator
	notifyer           notifier.Notifier
	notificationsRepo  notificationsRepo.NotificationsRepository
	downloadTokensRepo downloadTokensRepo.DownloadTokensRepository
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) getDbConn(_ context.Context) *sql.DB {
	if s.db == nil {
		dbDSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			config.PgUser(),
			config.PgPassword(),
			config.PgHost(),
			config.PgPort(),
			config.PgDatabase(),
		)

		db, err := sql.Open("pgx", dbDSN)
		if err != nil {
			logger.Fatalf("failed to init db connection: %s", err.Error())
		}

		s.db = db
	}

	return s.db
}

func (s *serviceProvider) getEmplRepo(ctx context.Context) emplRepo.EmployeesRepository {
	if s.EmployeesRepo == nil {
		s.EmployeesRepo = emplRepo.New(s.getDbConn(ctx))
	}
	return s.EmployeesRepo
}

func (s *serviceProvider) getPosRepo(ctx context.Context) posRepo.PositionsRepository {
	if s.PositionsRepo == nil {
		s.PositionsRepo = posRepo.New(s.getDbConn(ctx))
	}
	return s.PositionsRepo
}

func (s *serviceProvider) getTasksRepo(ctx context.Context) tasksRepo.TasksRepository {
	if s.TasksRepo == nil {
		s.TasksRepo = tasksRepo.New(s.getDbConn(ctx))
	}
	return s.TasksRepo
}

func (s *serviceProvider) getTrainingsRepo(ctx context.Context) trainingsRepo.TrainingsRepository {
	if s.TrainingsRepo == nil {
		s.TrainingsRepo = trainingsRepo.New(s.getDbConn(ctx))
	}
	return s.TrainingsRepo
}

func (s *serviceProvider) getNotificationsRepo(ctx context.Context) notificationsRepo.NotificationsRepository {
	if s.notificationsRepo == nil {
		s.notificationsRepo = notificationsRepo.New(s.getDbConn(ctx))
	}
	return s.notificationsRepo
}

func (s *serviceProvider) getdownloadTokensRepo(ctx context.Context) downloadTokensRepo.DownloadTokensRepository {
	if s.downloadTokensRepo == nil {
		s.downloadTokensRepo = downloadTokensRepo.New(s.getDbConn(ctx))
	}
	return s.downloadTokensRepo
}

func (s *serviceProvider) getTxManager(ctx context.Context) *txManager.Manager {
	if s.txManager == nil {
		s.txManager = txManager.New(s.getDbConn(ctx))
	}
	return s.txManager
}

func (s *serviceProvider) getDocsGenerator(_ context.Context) docsgenerator.DocsGenerator {
	if s.docsGenerator == nil {

		s.docsGenerator = docsgenerator.New(config.DocsGeneratorTamplatePath())
	}

	return s.docsGenerator
}

func (s *serviceProvider) getNotifier(_ context.Context) notifier.Notifier {
	if s.notifyer == nil {
		emailNotifier := notifier.New(
			config.NotifierEmailFrom(),
			config.NotifierEmailPassword(),
			config.NotifierSMTPHost(),
			config.NotifierSMTPPort(),
		)

		s.notifyer = emailNotifier
	}

	return s.notifyer
}

func (s *serviceProvider) getService(ctx context.Context) service.Servicer {
	if s.service == nil {
		s.service = service.New(
			s.getEmplRepo(ctx),
			s.getPosRepo(ctx),
			s.getTasksRepo(ctx),
			s.getTrainingsRepo(ctx),
			s.getTxManager(ctx),
			s.getDocsGenerator(ctx),
			s.getNotifier(ctx),
			s.getNotificationsRepo(ctx),
			s.getdownloadTokensRepo(ctx),
		)
	}
	return s.service
}

func (s *serviceProvider) getLearnControl(ctx context.Context) *impl.Implementation {
	if s.implementation == nil {
		s.implementation = impl.NewLearnControl(s.getService(ctx))
	}
	return s.implementation
}
