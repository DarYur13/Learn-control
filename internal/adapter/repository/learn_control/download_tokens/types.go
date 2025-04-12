package learncontrol

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DarYur13/learn-control/internal/domain"
	"github.com/google/uuid"
)

const (
	loadsCountLimit = 10
)

var (
	ErrTooManyLoads = errors.New("too many loads")
)

type DownloadTokensRepository interface {
	AddToken(ctx context.Context, token domain.DownloadToken) error
	GetRegistrationSheetInfoTx(ctx context.Context, tx *sql.Tx, token uuid.UUID) (*domain.RegistrationSheetInfo, error)
	GetToken(ctx context.Context, employeeID, trainingID int) (uuid.UUID, error)
	GetTokenInfo(ctx context.Context, token uuid.UUID) (domain.DownloadToken, error)
	GetTokenInfoTx(ctx context.Context, tx *sql.Tx, token uuid.UUID) (domain.DownloadToken, error)
	IncreaseLoadsCountTx(ctx context.Context, tx *sql.Tx, token uuid.UUID) error
}
