package learncontrol

import (
	"database/sql"
)

type TasksStorage struct {
	db *sql.DB
}

func New(db *sql.DB) TasksStorager {
	return &TasksStorage{
		db: db,
	}
}
