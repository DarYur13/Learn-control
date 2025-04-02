package storage

import "context"

const (
	queryGetTrainingsForPosition = `
		SELECT pt.training_id FROM positions p
		JOIN position_trainings pt ON p.id = pt.position_id
		WHERE p.position = $1 AND p.department = $2
	`
)

func (s *Storage) GetTrainingsForPosition(ctx context.Context, department, position string) ([]int, error) {
	rows, err := s.db.QueryContext(ctx, queryGetTrainingsForPosition, position, department)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []int
	for rows.Next() {
		var trainingID int
		if err := rows.Scan(&trainingID); err != nil {
			return nil, err
		}
		result = append(result, trainingID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
