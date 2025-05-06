package learncontrol

import (
	"context"
	"database/sql"
	"time"

	trainingsStorage "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/trainings"
	"github.com/golang-sql/sqlexp"
)

const (
	queryUpdateEmployeeTrainingDate = `
	UPDATE employee_trainings
	SET 
		training_date = $1::timestamp,
		retraining_date = $1::timestamp + (t.valid_period || ' months')::interval
	FROM trainings t
	WHERE 
		training_id = t.id AND
		employee_id = $2 AND 
		training_id = $3
	RETURNING training_date, retraining_date;
	`
)

func (es *EmployeesStorage) UpdateEmployeeTrainingDateTx(ctx context.Context, tx *sql.Tx, employeeID int, trainingID int, date time.Time) (*trainingsStorage.TrainingDates, error) {
	return es.updateEmployeeTrainingDate(ctx, tx, employeeID, trainingID, date)
}

func (es *EmployeesStorage) UpdateEmployeeTrainingDate(ctx context.Context, employeeID int, trainingID int, date time.Time) (*trainingsStorage.TrainingDates, error) {
	return es.updateEmployeeTrainingDate(ctx, es.db, employeeID, trainingID, date)
}

func (es *EmployeesStorage) updateEmployeeTrainingDate(ctx context.Context, tx sqlexp.Querier, employeeID int, trainingID int, date time.Time) (*trainingsStorage.TrainingDates, error) {
	var trainingDates trainingsStorage.TrainingDates

	err := tx.QueryRowContext(ctx, queryUpdateEmployeeTrainingDate,
		date,
		employeeID,
		trainingID,
	).Scan(
		&trainingDates.PassDate,
		&trainingDates.RePassDate,
	)
	if err != nil {
		return nil, err
	}

	return &trainingDates, nil
}
