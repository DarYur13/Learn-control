package txmanager

import (
	"context"
	"database/sql"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

func (m *Manager) Do(ctx context.Context, fn func(tx *sql.Tx) error) error {
	var (
		tx  *sql.Tx
		err error
	)

	if tx, err = m.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return err
	}

	defer func() {
		if rbErr := tx.Rollback(); rbErr != nil {
			var mErr *multierror.Error

			mErr = multierror.Append(mErr, errors.WithMessage(err, "transaction"))
			mErr = multierror.Append(mErr, errors.WithMessage(rbErr, "rollback"))

			err = mErr.ErrorOrNil()

			return
		}
	}()

	err = fn(tx)
	if err == nil {
		return tx.Commit()
	}

	return err
}
