package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

const (
	// Получаем сотрудников, которым скоро проходить переобучение, исключая тех кто уже его прошел
	queryGetUpcomingTrainings = `
	SELECT 
		et.employee_id,
		et.training_id,
		p.id,
		t.training_type,
		et.retraining_date,
		DATE_PART('day', et.retraining_date - CURRENT_DATE) AS days_left
	FROM employee_trainings et
	JOIN trainings t ON t.id = et.training_id
	JOIN employees e ON e.id = et.employee_id
	JOIN positions p ON e.position = p.position AND e.department = p.department
	WHERE DATE_PART('day', et.retraining_date - CURRENT_DATE) IN (10, 30)
	AND NOT EXISTS (
		SELECT 1
		FROM employee_trainings newer
		WHERE newer.employee_id = et.employee_id
			AND newer.training_id = et.training_id
			AND newer.retraining_date > CURRENT_DATE + INTERVAL '30 days'
	)
	`
)

func (ts *TrainingsStorage) GetUpcomingTrainings(ctx context.Context) ([]domain.UpcomingTraining, error) {
	var trainings []domain.UpcomingTraining

	rows, err := ts.db.QueryContext(ctx, queryGetUpcomingTrainings)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var training domain.UpcomingTraining

		if err := rows.Scan(
			&training.EmployeeID,
			&training.TrainingID,
			&training.PositionID,
			&training.TrainingType,
			&training.RePassDate,
			&training.DaysLeft,
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
