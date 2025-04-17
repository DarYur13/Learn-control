package learncontrol

import "context"

const (
	queryGetTrainingsForPosition = `
		SELECT 
			pt.training_id,
			pt.position_id
		FROM positions p
		JOIN position_trainings pt ON p.id = pt.position_id
		WHERE p.position = $1 AND p.department = $2
	`
)

func (ps *PositionsStorage) GetPositionTrainings(ctx context.Context, department, position string) (int, []int, error) {
	rows, err := ps.db.QueryContext(ctx, queryGetTrainingsForPosition, position, department)
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()

	var (
		trainingsIDs []int
		positionID   int
	)

	for rows.Next() {
		var trainingID int

		if err := rows.Scan(
			&trainingID,
			&positionID,
		); err != nil {
			return 0, nil, err
		}

		trainingsIDs = append(trainingsIDs, trainingID)
	}

	if err := rows.Err(); err != nil {
		return 0, nil, err
	}

	return positionID, trainingsIDs, nil
}
