package worker

import (
	"time"
)

func (nw *notificationWorker) Interval() time.Duration {
	return nw.interval
}
