package learncontrol

import (
	"context"
	"database/sql"

	"github.com/golang-sql/sqlexp"
)

const (
	queryCloseTask = `
		UPDATE tasks
		SET done = true AND done_at = CURRENT_DATE
		WHERE id = $1
	`
)

func (ts *TasksStorage) CloseTask(ctx context.Context, taskID int) error {
	return ts.closeTask(ctx, ts.db, taskID)
}

func (ts *TasksStorage) CloseTaskTx(ctx context.Context, tx *sql.Tx, taskID int) error {
	return ts.closeTask(ctx, tx, taskID)
}

func (ts *TasksStorage) closeTask(ctx context.Context, tx sqlexp.Querier, taskID int) error {
	_, err := tx.ExecContext(ctx, queryCloseTask, taskID)
	if err != nil {
		return err
	}

	return nil
}
