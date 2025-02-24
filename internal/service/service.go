package service

import "github.com/DarYur13/learn-control/internal/storage"

// Service
type Service struct {
	storage storage.IStorage
}

// New creates new service
func New(storage storage.IStorage) *Service {
	return &Service{
		storage: storage,
	}
}
