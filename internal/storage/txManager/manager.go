package txmanager

import "database/sql"

type Manager struct {
	db *sql.DB
}

func New(s *sql.DB) IManager {
	return &Manager{db: s}
}
