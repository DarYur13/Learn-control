package domain

import "database/sql"

type Filters struct {
	Department         sql.NullString
	Position           sql.NullString
	TrainingID         sql.NullInt64
	DateFrom           sql.NullTime
	DateTo             sql.NullTime
	TrainingsNotPassed sql.NullBool
	RetrainingIn       sql.NullInt64
}
