package learncontrol

import (
	"context"
	"database/sql"
)

type TasksStorager interface {
	AddTaskTx(ctx context.Context, tx *sql.Tx, task TaskBaseInfo) error
	AddTask(ctx context.Context, task TaskBaseInfo) error
	GetTasksByFilters(ctx context.Context, done sql.NullBool) ([]Task, error)
	CloseTask(ctx context.Context, taskID int) error
	CloseTaskTx(ctx context.Context, tx *sql.Tx, taskID int) error
}

type TaskBaseInfo struct {
	Type       string
	TrainingID sql.NullInt64
	EmployeeID sql.NullInt64
	ExecutorID sql.NullInt64
	PositionID sql.NullInt64
}

type Task struct {
	ID          int
	Type        string
	Description string
	Employee    sql.NullString
	Training    sql.NullString
	Position    sql.NullString
	Department  sql.NullString
	Executor    sql.NullString
	Done        bool
}
