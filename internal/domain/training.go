package domain

import (
	"time"

	desc "github.com/DarYur13/learn-control/pkg/learn_control"
)

const dateFormat = "02.01.2006"

type Training struct {
	Name       string
	PassDate   time.Time
	RePassDate time.Time
}

func (t *Training) ToDesc() *desc.Training {
	return &desc.Training{
		Name:     t.Name,
		Date:     t.PassDate.Format(dateFormat),
		Nextdate: t.RePassDate.Format(dateFormat),
	}
}
