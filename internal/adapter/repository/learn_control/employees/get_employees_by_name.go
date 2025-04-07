package learncontrol

import (
	"context"
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

func (es *EmployeesStorage) GetEmployeesByName(ctx context.Context, name string) ([]EmployeeBaseInfo, error) {
	rows, err := es.db.QueryContext(ctx, queryGetEmployeesByName, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []EmployeeBaseInfo

	for rows.Next() {
		var employeeBaseInfo EmployeeBaseInfo

		if err := rows.Scan(
			&employeeBaseInfo.ID,
			&employeeBaseInfo.FullName,
			&employeeBaseInfo.BirthDate,
		); err != nil {
			return nil, err
		}

		result = append(result, employeeBaseInfo)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return result, nil
}
