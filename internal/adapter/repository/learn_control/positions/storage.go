package learncontrol

import (
	"database/sql"
)

type PositionsStorage struct {
	db *sql.DB
}

func New(db *sql.DB) PositionsRepository {
	return &PositionsStorage{
		db: db,
	}
}
