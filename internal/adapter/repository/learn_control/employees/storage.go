package learncontrol

import (
	"database/sql"
)

// Repository with postgress connection
type EmployeesStorage struct {
	db *sql.DB
}

// New creates new repository object
func New(db *sql.DB) EmoloyeesStorager {
	return &EmployeesStorage{
		db: db,
	}
}
