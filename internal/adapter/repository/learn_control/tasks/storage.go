package learncontrol

import (
	"database/sql"
)

type TasksStorage struct {
	db *sql.DB
}

func New(db *sql.DB) TasksRepository {
	return &TasksStorage{
		db: db,
	}
}
