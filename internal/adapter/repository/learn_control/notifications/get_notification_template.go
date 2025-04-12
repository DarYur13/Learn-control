package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

const (
	queryGetNotificationTemplates = `
	SELECT 
		subject_template, 
		body_template
	FROM notification_types_templates
	WHERE notificationType = $1
	`
)

func (ns *NotificationsStorage) GetNotificationTemplate(ctx context.Context, notificationType domain.NotificationType) (*domain.SMTPNotificationTemplate, error) {
	var template domain.SMTPNotificationTemplate

	if err := ns.db.QueryRowContext(ctx, queryGetNotificationTemplates).Scan(
		&template.Subject,
		&template.Body,
	); err != nil {
		return nil, err
	}

	return &template, nil
}
