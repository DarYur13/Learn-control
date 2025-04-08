package learncontrol

import (
	"context"
)

const (
	queryGetEmployeeLeader = `
	SELECT id
	FROM employees
	WHERE department = (
		SELECT department FROM employees WHERE id = $1
	) AND is_leader = TRUE
	LIMIT 1
	`
)

func (es *EmployeesStorage) GetEmployeeLeader(ctx context.Context, employeeID int) (int, error) {
	var leaderID int

	if err := es.db.QueryRowContext(ctx, queryGetEmployeeLeader, employeeID).Scan(&leaderID); err != nil {
		return 0, err
	}

	return leaderID, nil
}
