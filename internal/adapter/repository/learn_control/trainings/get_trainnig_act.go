package learncontrol

import "context"

const (
	queryGetTrainingAct = `
	SELECT la.act_name
	FROM local_acts la
	JOIN trainings t ON la.training_id = t.id
	WHERE t.id = $1;
	`
)

func (ts *TrainingsStorage) GetTrainingAct(ctx context.Context, trainingID int) (string, error) {
	var act string

	if err := ts.db.QueryRowContext(ctx, queryGetTrainingAct, trainingID).Scan(&act); err != nil {
		return "", err
	}

	return act, nil
}
