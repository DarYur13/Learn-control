package modules

import (
	"github.com/kelseyhightower/envconfig"
)

const RetrainingControlWorkerModulePrefix = "RETRAINING_CONTROL_WORKER"

type RetrainingControlWorker struct {
	QueueCheckPeriod int `envconfig:"QUEUE_CHECK_PERIOD"`
}

func LoadRetrainingControlWorker() (*RetrainingControlWorker, error) {
	var nw RetrainingControlWorker
	err := envconfig.Process(RetrainingControlWorkerModulePrefix, &nw)

	if err != nil {
		return nil, err
	}

	return &nw, nil
}
