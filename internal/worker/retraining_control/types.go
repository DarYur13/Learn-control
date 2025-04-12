package worker

import (
	"context"
)

const (
	daysTillRetrainingFirst  = 30
	daysTillRetrainingSecond = 10
)

type RetrainingControlWorker interface {
	Start(ctx context.Context)
	Interval() float64
}
