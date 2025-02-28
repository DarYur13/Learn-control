package storage

import (
	"context"
	"database/sql"

	"github.com/DarYur13/learn-control/internal/domain"
)

type IStorage interface {
	GetEmployees(ctx context.Context, name string) (*domain.EmployeesBaseInfo, error)
	GetEmployee(ctx context.Context, id int) (*Employee, error)
}

type Employee struct {
	FullName   string `db:"full_name"`
	BirthDate  string `db:"birth_date"`
	Snils      string `db:"snils"`
	Department string `db:"department"`
	Position   string `db:"position"`
	Trainings  []Training
}

type Training struct {
	Name       string       `db:"training"`
	PassDate   sql.NullTime `db:"pass_date"`
	RePassDate sql.NullTime `db:"repass_date"`
}
