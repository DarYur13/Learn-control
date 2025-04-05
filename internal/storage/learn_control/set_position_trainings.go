package storage

import (
	"context"
	"database/sql"
)

const (
	querySetPositionTrainings = `
	INSERT INTO position_trainings
		position_id,
		training_id
	) VALUES ($1, $2)
	ON CONFLICT (employee_id, training_id) DO NOTHING;
	`
)

func (s *Storage) SetPositionTrainingsTx(ctx context.Context, tx *sql.Tx, positionID int, trainingsIDs []int) error {
	stmt, err := tx.PrepareContext(ctx, querySetPositionTrainings)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, trainingID := range trainingsIDs {
		if _, err := stmt.ExecContext(ctx, positionID, trainingID); err != nil {
			return err
		}
	}

	return nil
}
