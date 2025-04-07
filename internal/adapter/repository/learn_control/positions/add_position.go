package learncontrol

import (
	"context"
	"database/sql"
)

const (
	queryAddPosition = `
		INSERT INTO positions (
			position,
			department
		) VALUES ($1, $2)
		RETURNING id
	`
)

func (ps *PositionsStorage) AddPositionTx(ctx context.Context, tx *sql.Tx, position, department string) (int, error) {
	var positionID int

	if err := tx.QueryRowContext(ctx, queryAddPosition, position, department).Scan(&positionID); err != nil {
		return 0, err
	}

	return positionID, nil
}
