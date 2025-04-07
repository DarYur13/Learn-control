package service_provider

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	impl "github.com/DarYur13/learn-control/internal/adapter/controller/learn_control"
	emplRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/employees"
	posRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/positions"
	tasksRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/tasks"
	txManager "github.com/DarYur13/learn-control/internal/adapter/repository/txManager"
	"github.com/DarYur13/learn-control/internal/config"
	"github.com/DarYur13/learn-control/internal/logger"
	service "github.com/DarYur13/learn-control/internal/service"
)

// serviceProvider di-container
type serviceProvider struct {
	db             *sql.DB
	txManager      *txManager.Manager
	EmployeesRepo  emplRepo.EmoloyeesStorager
	PositionsRepo  posRepo.PositionsStorager
	TasksRepo      tasksRepo.TasksStorager
	service        service.Servicer
	implementation *impl.Implementation
	minioCli       *minio.Client
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

func (s *serviceProvider) getEmplRepo(ctx context.Context) emplRepo.EmoloyeesStorager {
	if s.EmployeesRepo == nil {
		s.EmployeesRepo = emplRepo.New(s.getDbConn(ctx))
	}
	return s.EmployeesRepo
}

func (s *serviceProvider) getPosRepo(ctx context.Context) posRepo.PositionsStorager {
	if s.PositionsRepo == nil {
		s.PositionsRepo = posRepo.New(s.getDbConn(ctx))
	}
	return s.PositionsRepo
}

func (s *serviceProvider) getTasksRepo(ctx context.Context) tasksRepo.TasksStorager {
	if s.TasksRepo == nil {
		s.TasksRepo = tasksRepo.New(s.getDbConn(ctx))
	}
	return s.TasksRepo
}

func (s *serviceProvider) getTxManager(ctx context.Context) *txManager.Manager {
	if s.txManager == nil {
		s.txManager = txManager.New(s.getDbConn(ctx))
	}
	return s.txManager
}

func (s *serviceProvider) getMinioCli(_ context.Context) *minio.Client {
	if s.minioCli == nil {
		minioEndpoint := fmt.Sprintf("%s:%s", config.MinioHost(), config.MinioPort())

		minioAccessKey := config.MinioUser()
		minioSecretKey := config.MinioPassword()

		minioClient, err := minio.New(minioEndpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(minioAccessKey, minioSecretKey, ""),
			Secure: false,
		})
		if err != nil {
			logger.Fatalf("failed to create MinIO client: %s", err.Error())
		}

		s.minioCli = minioClient
	}

	return s.minioCli
}

func (s *serviceProvider) getService(ctx context.Context) service.Servicer {
	if s.service == nil {
		s.service = service.New(
			s.getEmplRepo(ctx),
			s.getPosRepo(ctx),
			s.getTasksRepo(ctx),
			s.getTxManager(ctx),
			s.getMinioCli(ctx),
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
