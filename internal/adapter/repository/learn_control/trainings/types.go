package learncontrol

import (
	"context"
	"database/sql"

	"github.com/DarYur13/learn-control/internal/domain"
)

type TrainingsRepository interface {
	GetTrainings(ctx context.Context) ([]TrainigBaseInfo, error)
	GetTrainingType(ctx context.Context, trainingID int) (domain.TrainingType, error)
	GetTrainingAct(ctx context.Context, trainingID int) (string, error)
	GetUpcomingTrainings(ctx context.Context) ([]domain.UpcomingTraining, error)
}

type Training struct {
	ID          int                 `db:"id" json:"id"`
	Name        string              `db:"training_name" json:"name"`
	Type        domain.TrainingType `db:"training_type"`
	HasProtocol sql.NullBool        `db:"has_protocol" json:"has_protocol"`
	TrainingDates
}

type TrainingDates struct {
	PassDate   sql.NullTime `db:"training_date" json:"pass_date"`
	RePassDate sql.NullTime `db:"retraining_date" json:"repass_date"`
}

type TrainigBaseInfo struct {
	ID   int                 `db:"id"`
	Type domain.TrainingType `db:"training_type"`
	Name string              `db:"training_name"`
}
