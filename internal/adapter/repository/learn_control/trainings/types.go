package learncontrol

import (
	"context"
	"database/sql"
)

type TrainingsRepository interface {
	GetTrainings(ctx context.Context) ([]TrainigBaseInfo, error)
}

type Training struct {
	Name        string       `db:"training" json:"name"`
	HasProtocol sql.NullBool `db:"has_protocol" json:"has_protocol"`
	TrainingDates
}

type TrainingDates struct {
	PassDate   sql.NullTime `db:"training_date" json:"pass_date"`
	RePassDate sql.NullTime `db:"retraining_date" json:"repass_date"`
}

type TrainigBaseInfo struct {
	ID   int    `db:"id"`
	Name string `db:"training"`
}
