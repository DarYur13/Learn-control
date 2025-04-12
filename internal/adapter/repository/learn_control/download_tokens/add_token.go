package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

const queryInsertDownloadToken = `
	INSERT INTO download_tokens (
		token, 
		employee_id, 
		training_id, 
		expires_at
	)
	VALUES ($1, $2, $3, $4)
`

func (dts *downloadTokensStorage) AddToken(ctx context.Context, token domain.DownloadToken) error {
	_, err := dts.db.ExecContext(ctx, queryInsertDownloadToken,
		token.Token,
		token.EmployeeID,
		token.TrainingID,
		token.ExpiresAt,
	)
	return err
}
