package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

type NotificationsRepository interface {
	GetNotificationTemplate(ctx context.Context, notificationType domain.NotificationType) (*domain.SMTPNotificationTemplate, error)
	AddNotificationToQueue(ctx context.Context, employeeID, trainingID int, notificationType domain.NotificationType) error
	GetPendingNotifications(ctx context.Context) ([]domain.PendingNotification, error)
	MarkNotificationAsSent(ctx context.Context, id int) error
}
