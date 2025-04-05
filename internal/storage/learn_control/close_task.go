package storage

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

func (s *Storage) CloseTask(ctx context.Context, taskID int) error {
	return s.closeTask(ctx, s.db, taskID)
}

func (s *Storage) CloseTaskTx(ctx context.Context, tx *sql.Tx, taskID int) error {
	return s.closeTask(ctx, tx, taskID)
}

func (s *Storage) closeTask(ctx context.Context, tx sqlexp.Querier, taskID int) error {
	_, err := tx.ExecContext(ctx, queryCloseTask, taskID)
	if err != nil {
		return err
	}

	return nil
}
