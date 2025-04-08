package learncontrol

import (
	"database/sql"
)

type TrainingsStorage struct {
	db *sql.DB
}

func New(db *sql.DB) TrainingsRepository {
	return &TrainingsStorage{
		db: db,
	}
}
