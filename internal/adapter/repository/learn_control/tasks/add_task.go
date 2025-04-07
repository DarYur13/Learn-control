package learncontrol

import (
	"context"
	"database/sql"

	"github.com/golang-sql/sqlexp"
)

const (
	queryAddTask = `
		INSERT INTO tasks (
			task_type, 
			training_id, 
			employee_id, 
			executor_id, 
			position_id
		) VALUES ($1, $2, $3, $4, $5)
	`
)

func (ts *TasksStorage) AddTaskTx(ctx context.Context, tx *sql.Tx, task TaskBaseInfo) error {
	return ts.addTask(ctx, tx, task)
}

func (ts *TasksStorage) AddTask(ctx context.Context, task TaskBaseInfo) error {
	return ts.addTask(ctx, ts.db, task)
}

func (ts *TasksStorage) addTask(ctx context.Context, tx sqlexp.Querier, task TaskBaseInfo) error {
	_, err := tx.ExecContext(ctx, queryAddTask,
		task.Type,
		task.TrainingID,
		task.EmployeeID,
		task.ExecutorID,
		task.PositionID,
	)
	if err != nil {
		return err
	}

	return nil
}
