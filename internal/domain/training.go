package domain

import (
	"database/sql"
	"time"

	desc "github.com/DarYur13/learn-control/pkg/learn_control"
)

const dateFormat = "02.01.2006"

type Training struct {
	Name       string
	PassDate   time.Time
	RePassDate sql.NullTime
}

func (t *Training) ToDesc() *desc.Training {
	result := &desc.Training{
		Name: t.Name,
		Date: t.PassDate.Format(dateFormat),
	}

	if t.RePassDate.Valid {
		result.Nextdate = t.RePassDate.Time.Format(dateFormat)
	} else {
		result.Nextdate = "Не требуется"
	}

	return result
}
