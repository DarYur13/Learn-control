package storage

import (
	"context"
	"database/sql"
)

type IStorage interface {
	GetEmployeesByName(ctx context.Context, name string) ([]EmployeeBaseInfo, error)
	GetEmployee(ctx context.Context, id int) (*Employee, error)
	GetPositions(ctx context.Context) ([]string, error)
	GetDepartments(ctx context.Context) ([]string, error)
	GetTrainings(ctx context.Context) ([]TrainigBaseInfo, error)
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

type EmployeeBaseInfo struct {
	ID        int64
	FullName  string
	BirthDate string
}

type TrainigBaseInfo struct {
	ID   int    `db:"id"`
	Name string `db:"training"`
}
