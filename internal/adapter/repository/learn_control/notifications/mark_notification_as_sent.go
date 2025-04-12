package learncontrol

import "context"

const queryMarkNotificationAsSent = `
	UPDATE notifications_queue
	SET is_sent = TRUE
	WHERE id = $1
`

func (ns *NotificationsStorage) MarkNotificationAsSent(ctx context.Context, id int) error {
	_, err := ns.db.ExecContext(ctx, queryMarkNotificationAsSent, id)
	return err
}
