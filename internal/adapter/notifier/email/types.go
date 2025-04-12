package notifier

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

type Notifier interface {
	Send(ctx context.Context, notification domain.SMTPNotification) error
	SendWithAttachment(ctx context.Context, notification domain.SMTPNotification) error
}
