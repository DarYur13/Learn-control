package learncontrol

import (
	"context"
	"database/sql"
	"time"

	trainingsRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/trainings"
)

type EmployeesRepository interface {
	AddEmployeeTx(ctx context.Context, tx *sql.Tx, employee Employee) (int, error)

	GetEmployeeByID(ctx context.Context, employeeID int) (*Employee, error)
	GetEmployeesByName(ctx context.Context, name string) ([]EmployeeBaseInfo, error)
	GetEmployeePersonalCard(ctx context.Context, id int) (*EmployeePersonalCard, error)
	GetEmployeesByFilters(ctx context.Context, filters Filters) ([]EmployeeInfo, error)
	GetEmployeesWithoutTrainings(ctx context.Context, positionID int) ([]int, error)
	GetEmployeesWithoutTrainingsTx(ctx context.Context, tx *sql.Tx, positionID int) ([]int, error)
	GetEmployeeLeader(ctx context.Context, employeeID int) (int, error)

	SetEmployeeTrainingsTx(ctx context.Context, tx *sql.Tx, employeeID int, trainingIDs []int) error
	SetEmployeeTrainnigProtocol(ctx context.Context, employeeID, trainingID int) error
	SetEmployeeTrainnigProtocolTx(ctx context.Context, tx *sql.Tx, employeeID, trainingID int) error

	UpdateEmployeeTrainingDateTx(ctx context.Context, tx *sql.Tx, employeeID int, trainingID int, date time.Time) (*trainingsRepo.TrainingDates, error)
}

type Employee struct {
	FullName       string `db:"full_name"`
	BirthDate      string `db:"birth_date"`
	Snils          string `db:"snils"`
	Department     string `db:"department"`
	Position       string `db:"position"`
	EmploymentDate string `db:"employment_date"`
	Email          string `db:"email"`
}

type EmployeePersonalCard struct {
	FullName       string `db:"full_name"`
	BirthDate      string `db:"birth_date"`
	Snils          string `db:"snils"`
	Department     string `db:"department"`
	Position       string `db:"position"`
	EmploymentDate string `db:"employment_date"`
	Trainings      []trainingsRepo.Training
}

type EmployeeBaseInfo struct {
	ID        int64
	FullName  string
	BirthDate string
}

type EmployeeInfo struct {
	FullName   string                   `db:"full_name" json:"full_name"`
	Department string                   `db:"department" json:"department"`
	Position   string                   `db:"position" json:"position"`
	Trainings  []trainingsRepo.Training `json:"trainings"`
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
