package learncontrol

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const (
	queryIncreaseLoadsCount = `
	UPDATE download_tokens
	SET loads_count = loads_count + 1
	WHERE token = $1
	`
)

func (dts *downloadTokensStorage) IncreaseLoadsCountTx(ctx context.Context, tx *sql.Tx, token uuid.UUID) error {
	_, err := tx.ExecContext(ctx, queryIncreaseLoadsCount, token)
	if err != nil {
		return err
	}

	return nil
}
