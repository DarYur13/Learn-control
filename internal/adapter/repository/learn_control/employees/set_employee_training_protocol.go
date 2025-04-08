package learncontrol

import (
	"context"
	"database/sql"

	"github.com/golang-sql/sqlexp"
)

const (
	querySetEmployeeTrainnigProtocol = `
	UPDATE employee_trainings SET has_protocol = true
	WHERE employee_id = $1 AND training_id = $2
	`
)

func (es *EmployeesStorage) SetEmployeeTrainnigProtocol(ctx context.Context, employeeID, trainingID int) error {
	return es.setEmployeeTrainnigProtocol(ctx, es.db, employeeID, trainingID)
}

func (es *EmployeesStorage) SetEmployeeTrainnigProtocolTx(ctx context.Context, tx *sql.Tx, employeeID, trainingID int) error {
	return es.setEmployeeTrainnigProtocol(ctx, tx, employeeID, trainingID)
}

func (es *EmployeesStorage) setEmployeeTrainnigProtocol(ctx context.Context, tx sqlexp.Querier, employeeID, trainingID int) error {
	_, err := tx.ExecContext(ctx, querySetEmployeeTrainnigProtocol, employeeID, trainingID)
	if err != nil {
		return err
	}

	return nil
}
