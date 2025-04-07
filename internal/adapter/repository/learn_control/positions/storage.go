package learncontrol

import (
	"database/sql"
)

// Repository with postgress connection
type PositionsStorage struct {
	db *sql.DB
}

// New creates new repository object
func New(db *sql.DB) PositionsStorager {
	return &PositionsStorage{
		db: db,
	}
}
