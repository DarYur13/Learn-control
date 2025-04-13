package learncontrol

import (
	"context"
)

const (
	queryGetTrainings = `SELECT DISTINCT id, training_type, training_name FROM trainings`
)

func (ts *TrainingsStorage) GetTrainings(ctx context.Context) ([]TrainigBaseInfo, error) {
	rows, err := ts.db.QueryContext(ctx, queryGetTrainings)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		trainings []TrainigBaseInfo
	)

	for rows.Next() {
		var training TrainigBaseInfo

		if err := rows.Scan(
			&training.ID,
			&training.Type,
			&training.Name,
		); err != nil {
			return nil, err
		}

		trainings = append(trainings, training)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return trainings, nil
}
