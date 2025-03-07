package storage

import (
	"context"
)

const (
	queryGetPositions = `SELECT DISTINCT position FROM positions`
)

func (s *Storage) GetPositions(ctx context.Context) ([]string, error) {
	rows, err := s.db.QueryContext(ctx, queryGetPositions)
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

	return positions, nil
}
