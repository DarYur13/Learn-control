package service

import (
	storage "github.com/DarYur13/learn-control/internal/storage/learn_control"
	txmanager "github.com/DarYur13/learn-control/internal/storage/txManager"
)

// Service
type Service struct {
	txManager txmanager.IManager
	storage   storage.IStorage
}

// New creates new service
func New(storage storage.IStorage, txManager txmanager.IManager) *Service {
	return &Service{
		storage:   storage,
		txManager: txManager,
	}
}
