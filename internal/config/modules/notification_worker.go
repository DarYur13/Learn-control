package modules

import (
	"github.com/kelseyhightower/envconfig"
)

const NotificationWorkerModulePrefix = "NOTIFICATION_WORKER"

type NotificationWorker struct {
	QueueCheckPeriod int `envconfig:"QUEUE_CHECK_PERIOD"`
}

func LoadNotificationWorker() (*NotificationWorker, error) {
	var nw NotificationWorker
	err := envconfig.Process(NotificationWorkerModulePrefix, &nw)

	if err != nil {
		return nil, err
	}

	return &nw, nil
}
