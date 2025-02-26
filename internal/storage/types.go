package storage

import (
	"context"
	"database/sql"
	"time"

	"github.com/DarYur13/learn-control/internal/domain"
)

type IStorage interface {
	GetEmployees(ctx context.Context, name string) (*domain.EmployeesBaseInfo, error)
	GetEmployee(ctx context.Context, id int) (*domain.Employee, error)
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
	PassDate   time.Time    `db:"pass_date"`
	RePassDate sql.NullTime `db:"repass_date"`
}

func (e *Employee) toDomain() *domain.Employee {
	result := &domain.Employee{
		FullName:   e.FullName,
		BirthDate:  e.BirthDate,
		Snils:      e.Snils,
		Department: e.Department,
		Position:   e.Position,
	}

	for _, training := range e.Trainings {
		result.Trainings = append(result.Trainings, domain.Training(training))
	}

	return result
}
