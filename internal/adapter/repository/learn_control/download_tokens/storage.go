package learncontrol

import (
	"database/sql"
)

type downloadTokensStorage struct {
	db *sql.DB
}

func New(db *sql.DB) DownloadTokensRepository {
	return &downloadTokensStorage{
		db: db,
	}
}
