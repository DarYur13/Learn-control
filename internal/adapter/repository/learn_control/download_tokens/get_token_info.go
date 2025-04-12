package learncontrol

import (
	"context"
	"database/sql"

	"github.com/DarYur13/learn-control/internal/domain"
	"github.com/golang-sql/sqlexp"
	"github.com/google/uuid"
)

const (
	queryGetTokenInfo = `
	SELECT 
		employee_id,
		training_id,
		created_at,
		expires_at,
		loads_count
	FROM download_tokens
	WHERE token = $1
	`
)

func (dts *downloadTokensStorage) GetTokenInfoTx(ctx context.Context, tx *sql.Tx, token uuid.UUID) (domain.DownloadToken, error) {
	return dts.getTokenInfo(ctx, tx, token)
}

func (dts *downloadTokensStorage) GetTokenInfo(ctx context.Context, token uuid.UUID) (domain.DownloadToken, error) {
	return dts.getTokenInfo(ctx, dts.db, token)
}

func (dts *downloadTokensStorage) getTokenInfo(ctx context.Context, tx sqlexp.Querier, token uuid.UUID) (domain.DownloadToken, error) {
	var info domain.DownloadToken

	err := tx.QueryRowContext(ctx, queryGetTokenInfo, token).Scan(
		&info.EmployeeID,
		&info.TrainingID,
		&info.CreatedAt,
		&info.ExpiresAt,
		&info.LoadsCount,
	)
	if err != nil {
		return domain.DownloadToken{}, err
	}

	return info, nil
}
