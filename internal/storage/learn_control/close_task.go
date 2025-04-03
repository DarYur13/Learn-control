package storage

import "context"

const (
	queryCloseTask = `
		UPDATE tasks
		SET done = true AND done_at = CURRENT_DATE
		WHERE id = $1
	`
)

func (s *Storage) CloseTask(ctx context.Context, taskID int) error {
	_, err := s.db.ExecContext(ctx, queryCloseTask, taskID)
	if err != nil {
		return err
	}

	return nil
}
