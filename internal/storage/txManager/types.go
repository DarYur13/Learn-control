package txmanager

import (
	"context"
	"database/sql"
)

type IManager interface {
	Do(ctx context.Context, fn func(tx *sql.Tx) error) error
}
