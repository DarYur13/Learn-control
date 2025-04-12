package learncontrol

import (
	"context"

	"github.com/google/uuid"
)

const (
	queryGetToken = `
	SELECT 
		token
	FROM download_tokens
	WHERE employee_id = $1 
	  AND training_id = $2
	  AND expires_at > NOW()
	`
)

func (dts *downloadTokensStorage) GetToken(ctx context.Context, employeeID, trainingID int) (uuid.UUID, error) {
	var token uuid.UUID
	err := dts.db.QueryRowContext(ctx, queryGetToken, employeeID, trainingID).Scan(
		&token,
	)

	if err != nil {
		return uuid.UUID{}, err
	}

	return token, nil
}
