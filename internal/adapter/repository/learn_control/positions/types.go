package learncontrol

import (
	"context"
	"database/sql"
)

type PositionsStorager interface {
	GetPositions(ctx context.Context) ([]string, error)
	GetDepartments(ctx context.Context) ([]string, error)
	GetPositionTrainings(ctx context.Context, department, position string) ([]int, error)
	AddPositionTx(ctx context.Context, tx *sql.Tx, position, department string) (int, error)
	SetPositionTrainingsTx(ctx context.Context, tx *sql.Tx, positionID int, trainingsIDs []int) error
}
