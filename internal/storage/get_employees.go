package storage

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

const (
	queryGetEmployeesByName = `
		SELECT 
			id, 
			full_name, 
			birth_date 
		FROM employees 
		WHERE full_name ILIKE '%' || $1 || '%'
		ORDER BY full_name ASC
		LIMIT 10
	`
)

func (s *Storage) GetEmployeesByName(ctx context.Context, name string) (*domain.EmployeesBaseInfo, error) {
	var (
		result           domain.EmployeesBaseInfo
		employeeBaseInfo domain.EmployeeBaseInfo
	)

	rows, err := s.db.QueryContext(ctx, queryGetEmployeesByName, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&employeeBaseInfo.ID,
			&employeeBaseInfo.FullName,
			&employeeBaseInfo.BirthDate,
		); err != nil {
			return nil, err
		}

		result.Employees = append(result.Employees, employeeBaseInfo)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return &result, nil
}
