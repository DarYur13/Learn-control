package learncontrol

import (
	"context"
)

const (
	queryGetEmployeeByID = `
	SELECT 
		full_name, 
		birth_date,
		snils,
		department,
		position,
		employment_date,
		email
	FROM employees 
	WHERE id = $1
	`
)

func (es *EmployeesStorage) GetEmployeeByID(ctx context.Context, employeeID int) (*Employee, error) {
	employee := Employee{}

	if err := es.db.QueryRowContext(ctx, queryGetEmployeeByID, employeeID).Scan(
		&employee.FullName,
		&employee.BirthDate,
		&employee.Snils,
		&employee.Department,
		&employee.Position,
		&employee.EmploymentDate,
		&employee.Email,
	); err != nil {
		return nil, err
	}

	return &employee, nil
}
