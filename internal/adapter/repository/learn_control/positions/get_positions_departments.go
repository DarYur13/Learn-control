package learncontrol

import (
	"context"
)

const (
	queryGetDepartments = `SELECT DISTINCT department FROM positions`
)

func (ps *PositionsStorage) GetPositionsDepartments(ctx context.Context) ([]string, error) {
	rows, err := ps.db.QueryContext(ctx, queryGetDepartments)
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

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return departments, nil
}
