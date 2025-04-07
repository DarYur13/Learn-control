package learncontrol

import (
	"context"
	"database/sql"
)

const (
	queryAddEmloyee = `
	INSERT INTO employees (
		full_name, 
		birth_date, 
		department, 
		position, 
		snils, 
		employment_date
	) VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id
	`
)

func (es *EmployeesStorage) AddEmployeeTx(ctx context.Context, tx *sql.Tx, employee Employee) (int, error) {
	var emplyeeID int

	err := tx.QueryRowContext(ctx, queryAddEmloyee,
		employee.FullName,
		employee.BirthDate,
		employee.Department,
		employee.Position,
		employee.Snils,
		employee.EmploymentDate,
	).Scan(&emplyeeID)
	if err != nil {
		return 0, err
	}

	return emplyeeID, nil
}
