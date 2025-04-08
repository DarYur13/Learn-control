package learncontrol

import (
	"context"
)

const (
	queryGetPositions = `SELECT DISTINCT position FROM positions`
)

func (ps *PositionsStorage) GetPositions(ctx context.Context) ([]string, error) {
	rows, err := ps.db.QueryContext(ctx, queryGetPositions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var positions []string

	for rows.Next() {
		var position string

		if err := rows.Scan(&position); err != nil {
			return nil, err
		}

		positions = append(positions, position)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return positions, nil
}
