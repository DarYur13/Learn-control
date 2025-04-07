package learncontrol

import (
	"context"
	"database/sql"
)

const (
	querySetEmployeeTrainings = `
	INSERT INTO employee_trainings (
		employee_id,
		training_id,
		has_protocol
	)
	SELECT 
		$1,
		$2,
		CASE 
			WHEN t.need_protocol THEN FALSE
			ELSE NULL
		END
	FROM trainings t
	WHERE t.id = $2
	ON CONFLICT (employee_id, training_id) DO NOTHING;
	`
)

func (es *EmployeesStorage) SetEmployeeTrainingsTx(ctx context.Context, tx *sql.Tx, employeeID int, trainingIDs []int) error {
	stmt, err := tx.PrepareContext(ctx, querySetEmployeeTrainings)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, trainingID := range trainingIDs {
		if _, err := stmt.ExecContext(ctx, employeeID, trainingID); err != nil {
			return err
		}
	}

	return nil
}
