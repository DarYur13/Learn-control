package storage

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

func (s *Storage) AddTaskTx(ctx context.Context, tx *sql.Tx, task TaskBaseInfo) error {
	return s.addTask(ctx, tx, task)
}

func (s *Storage) AddTask(ctx context.Context, task TaskBaseInfo) error {
	return s.addTask(ctx, s.db, task)
}

func (s *Storage) addTask(ctx context.Context, tx sqlexp.Querier, task TaskBaseInfo) error {
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
