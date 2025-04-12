package domain

import (
	"time"

	"github.com/google/uuid"
)

type DownloadToken struct {
	Token      uuid.UUID
	EmployeeID int
	TrainingID int
	CreatedAt  time.Time
	ExpiresAt  time.Time
	LoadsCount int
}
