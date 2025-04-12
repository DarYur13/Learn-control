package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

const queryInsertNotificationToQueue = `
	INSERT INTO notification_queue (employee_id, training_id, notification_type)
	VALUES ($1, $2, $3)
`

func (ns *NotificationsStorage) AddNotificationToQueue(ctx context.Context, employeeID, trainingID int, notificationType domain.NotificationType) error {
	_, err := ns.db.ExecContext(ctx, queryInsertNotificationToQueue, employeeID, trainingID, notificationType)
	return err
}
