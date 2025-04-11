package worker

import (
	"context"
)

type NotificationWorker interface {
	StartNotify(ctx context.Context)
	Interval() float64
}
