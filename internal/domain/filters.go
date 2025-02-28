package domain

import "time"

type Filters struct {
	Deparment          string
	Position           string
	TrainingID         int
	DateFrom           time.Time
	DateTo             time.Time
	TrainingsNotPassed bool
	RetrainingIn       int
}
