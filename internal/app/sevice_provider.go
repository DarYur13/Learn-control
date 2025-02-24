package service_provider

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"

	cmdImpl "github.com/DarYur13/learn-control/internal/api/learn_control"
	"github.com/DarYur13/learn-control/internal/logger"
	cmdService "github.com/DarYur13/learn-control/internal/service"
	cmdStor "github.com/DarYur13/learn-control/internal/storage"
)

// serviceProvider di-container
type serviceProvider struct {
	db             *sql.DB
	storage        cmdStor.IStorage
	service        cmdService.ILearnControlService
	implementation *cmdImpl.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) getDbConn(_ context.Context) *sql.DB {
	if s.db == nil {
		dbDSN := "postgres://postgres:postgres@localhost:5432/learn_control"

		db, err := sql.Open("pgx", dbDSN)
		if err != nil {
			logger.Fatalf("failed to init db connection: %s", err.Error())
		}

		s.db = db
	}

	return s.db
}

func (s *serviceProvider) getStorage(ctx context.Context) cmdStor.IStorage {
	if s.storage == nil {
		s.storage = cmdStor.New(s.getDbConn(ctx))
	}
	return s.storage
}

func (s *serviceProvider) getService(ctx context.Context) cmdService.ILearnControlService {
	if s.service == nil {
		s.service = cmdService.New(s.getStorage(ctx))
	}
	return s.service
}

func (s *serviceProvider) getLearnControl(ctx context.Context) *cmdImpl.Implementation {
	if s.implementation == nil {
		s.implementation = cmdImpl.NewLearnControl(s.getService(ctx))
	}
	return s.implementation
}
