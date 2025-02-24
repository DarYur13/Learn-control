package storage

import (
	"database/sql"
)

// Repository with postgress connection
type Storage struct {
	db *sql.DB
}

// New creates new repository object
func New(db *sql.DB) IStorage {
	return &Storage{
		db: db,
	}
}
