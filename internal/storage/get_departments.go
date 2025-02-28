package storage

import (
	"context"
)

const (
	queryGetDepartments = `SELECT DISTINCT department FROM positions`
)

func (s *Storage) GetDepartments(ctx context.Context) ([]string, error) {
	rows, err := s.db.QueryContext(ctx, queryGetDepartments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []string

	for rows.Next() {
		var department string

		if err := rows.Scan(&department); err != nil {
			return nil, err
		}

		departments = append(departments, department)
	}

	return departments, nil
}
