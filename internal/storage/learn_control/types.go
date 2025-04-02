package storage

import (
	"context"
	"database/sql"
	"time"
)

type IStorage interface {
	GetEmployeesByName(ctx context.Context, name string) ([]EmployeeBaseInfo, error)
	GetEmployeePersonalCard(ctx context.Context, id int) (*EmployeePersonalCard, error)
	GetPositions(ctx context.Context) ([]string, error)
	GetDepartments(ctx context.Context) ([]string, error)
	GetTrainings(ctx context.Context) ([]TrainigBaseInfo, error)
	GetEmployeesByFilters(ctx context.Context, filters Filters) ([]EmployeeInfo, error)
	UpdateEmployeeTrainingDateTx(ctx context.Context, tx *sql.Tx, employeeID int, trainingID int, date time.Time) (*TrainingDates, error)
	AddEmployeeTx(ctx context.Context, tx *sql.Tx, employee Employee) (int, error)
	GetTrainingsForPosition(ctx context.Context, department, position string) ([]int, error)
	SetEmployeeTrainingsTx(ctx context.Context, tx *sql.Tx, employeeID int, trainingsIDs []int) error
	AddPositionTx(ctx context.Context, tx *sql.Tx, position, department string) (int, error)
}

type Employee struct {
	FullName       string `db:"full_name"`
	BirthDate      string `db:"birth_date"`
	Snils          string `db:"snils"`
	Department     string `db:"department"`
	Position       string `db:"position"`
	EmploymentDate string `db:"employment_date"`
}

type EmployeePersonalCard struct {
	FullName       string `db:"full_name"`
	BirthDate      string `db:"birth_date"`
	Snils          string `db:"snils"`
	Department     string `db:"department"`
	Position       string `db:"position"`
	EmploymentDate string `db:"employment_date"`
	Trainings      []Training
}

type Training struct {
	Name        string       `db:"training" json:"name"`
	HasProtocol sql.NullBool `db:"has_protocol" json:"has_protocol"`
	TrainingDates
}

type TrainingDates struct {
	PassDate   sql.NullTime `db:"training_date" json:"pass_date"`
	RePassDate sql.NullTime `db:"retraining_date" json:"repass_date"`
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

type EmployeeInfo struct {
	FullName   string     `db:"full_name" json:"full_name"`
	Department string     `db:"department" json:"department"`
	Position   string     `db:"position" json:"position"`
	Trainings  []Training `json:"trainings"`
}

type Filters struct {
	Department         sql.NullString
	Position           sql.NullString
	TrainingID         sql.NullInt64
	DateFrom           sql.NullTime
	DateTo             sql.NullTime
	TrainingsNotPassed sql.NullBool
	RetrainingIn       sql.NullInt64
	HasProtocol        sql.NullBool
}
