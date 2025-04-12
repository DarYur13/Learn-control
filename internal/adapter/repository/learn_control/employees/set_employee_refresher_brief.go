package learncontrol

import (
	"context"
)

const (
	querySetEmployeeRefresherBrief = `
	INSERT INTO employee_trainings (
		employee_id,
		training_id
	)
	SELECT 
		$1,
		id
	FROM trainings t
	WHERE t.training_type = 'REFRESHER'
	LIMIT 1
	RETURNING training_id;
	`
)

func (es *EmployeesStorage) SetEmployeeRefresherBrief(ctx context.Context, employeeID int) (int, error) {
	var trainingID int

	if err := es.db.QueryRowContext(ctx, querySetEmployeeRefresherBrief, employeeID).Scan(
		&trainingID,
	); err != nil {
		return 0, err
	}

	return trainingID, nil
}
