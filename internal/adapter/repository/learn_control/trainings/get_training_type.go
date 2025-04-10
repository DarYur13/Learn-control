package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

const (
	queryGetTrainingType = `
	SELECT training_type 
	FROM trainings
	WHERE id = $1;
	`
)

func (ts *TrainingsStorage) GetTrainingType(ctx context.Context, trainingID int) (domain.TrainingType, error) {
	var trainingType domain.TrainingType

	if err := ts.db.QueryRowContext(ctx, queryGetTrainingType, trainingID).Scan(&trainingType); err != nil {
		return "", err
	}

	return trainingType, nil
}
