package modules

import (
	"github.com/kelseyhightower/envconfig"
)

const NotifierModulePrefix = "NOTIFIER"

type Notifier struct {
	EmailFrom     string `envconfig:"EMAIL_FROM"`
	EmailPassword string `envconfig:"EMAIL_PASSWORD"`
	SMTPHost      string `envconfig:"SMTP_HOST"`
	SMTPPort      string `envconfig:"SMTP_PORT"`
	EmailUseTLS   string `envconfig:"EMAIL_USE_TLS"`
}

func LoadNotifier() (*Notifier, error) {
	var s Notifier
	err := envconfig.Process(NotifierModulePrefix, &s)

	if err != nil {
		return nil, err
	}

	return &s, nil
}
