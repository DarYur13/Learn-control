package worker

import (
	"context"
	"time"
)

type NotificationWorker interface {
	StartNotify(ctx context.Context)
	Interval() time.Duration
}
