package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
	"github.com/google/uuid"
)

type DownloadTokensRepository interface {
	AddToken(ctx context.Context, token domain.DownloadToken) error
	GetRegistrationSheetInfo(ctx context.Context, token uuid.UUID) (*domain.RegistrationSheetInfo, error)
}
