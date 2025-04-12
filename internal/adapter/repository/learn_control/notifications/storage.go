package learncontrol

import (
	"database/sql"
)

type NotificationsStorage struct {
	db *sql.DB
}

func New(db *sql.DB) NotificationsRepository {
	return &NotificationsStorage{
		db: db,
	}
}
