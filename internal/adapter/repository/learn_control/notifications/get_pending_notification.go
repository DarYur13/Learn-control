package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

const queryGetPendingNotifications = `
	SELECT id, employee_id, training_id, notification_type
	FROM notifications_queue
	WHERE is_sent = FALSE
	  AND created_at <= NOW()
	ORDER BY created_at ASC
`

func (ns *NotificationsStorage) GetPendingNotifications(ctx context.Context) ([]domain.PendingNotification, error) {
	rows, err := ns.db.QueryContext(ctx, queryGetPendingNotifications)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.PendingNotification
	for rows.Next() {
		var n domain.PendingNotification
		if err := rows.Scan(&n.ID, &n.EmployeeID, &n.TrainingID, &n.Type); err != nil {
			return nil, err
		}

		result = append(result, n)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
