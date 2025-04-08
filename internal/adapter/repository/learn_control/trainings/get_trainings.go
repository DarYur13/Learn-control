package learncontrol

import (
	"context"
)

const (
	queryGetTrainings = `SELECT DISTINCT id, training FROM trainings`
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

		if err := rows.Scan(&training.ID, &training.Name); err != nil {
			return nil, err
		}

		trainings = append(trainings, training)
	}

	return trainings, nil
}
