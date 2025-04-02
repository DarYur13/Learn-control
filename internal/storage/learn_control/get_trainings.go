package storage

import (
	"context"
)

const (
	queryGetTrainings = `SELECT DISTINCT id, training FROM trainings`
)

func (s *Storage) GetTrainings(ctx context.Context) ([]TrainigBaseInfo, error) {
	rows, err := s.db.QueryContext(ctx, queryGetTrainings)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		trainings []TrainigBaseInfo
	)

	for rows.Next() {
		var training TrainigBaseInfo

		if err := rows.Scan(&training.ID, &training.Name); err != nil {
			return nil, err
		}

		trainings = append(trainings, training)
	}

	return trainings, nil
}
