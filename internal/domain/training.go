package domain

import (
	desc "github.com/DarYur13/learn-control/pkg/learn_control"
)

type Training struct {
	Name       string
	PassDate   string
	RePassDate string
}

func (t *Training) ToDesc() *desc.Training {
	return &desc.Training{
		Name:       t.Name,
		PassDate:   t.PassDate,
		RePassDate: t.RePassDate,
	}
}
