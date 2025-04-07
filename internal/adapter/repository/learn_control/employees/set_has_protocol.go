package learncontrol

import (
	"context"
	"database/sql"

	"github.com/golang-sql/sqlexp"
)

const (
	querySetHasProtocol = `
	UPDATE employee_trainings SET has_protocol = true
	WHERE employee_id = $1 AND training_id = $2
	`
)

func (es *EmployeesStorage) SetHasProtocol(ctx context.Context, employeeID, trainingID int) error {
	return es.setHasProtocol(ctx, es.db, employeeID, trainingID)
}

func (es *EmployeesStorage) SetHasProtocolTx(ctx context.Context, tx *sql.Tx, employeeID, trainingID int) error {
	return es.setHasProtocol(ctx, tx, employeeID, trainingID)
}

func (es *EmployeesStorage) setHasProtocol(ctx context.Context, tx sqlexp.Querier, employeeID, trainingID int) error {
	_, err := tx.ExecContext(ctx, querySetHasProtocol, employeeID, trainingID)
	if err != nil {
		return err
	}

	return nil
}
