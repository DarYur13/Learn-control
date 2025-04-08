package learncontrol

import (
	"database/sql"
)

type EmployeesStorage struct {
	db *sql.DB
}

func New(db *sql.DB) EmployeesRepository {
	return &EmployeesStorage{
		db: db,
	}
}
