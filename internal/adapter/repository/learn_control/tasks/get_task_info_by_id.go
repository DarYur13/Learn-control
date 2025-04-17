package learncontrol

import (
	"context"
)

const (
	queryGetTaskInfoByID = `
	SELECT
		training_id, 
		employee_id, 
		executor_id, 
		position_id
	FROM tasks
	WHERE id = $1;
	`
)

func (ts *TasksStorage) GetTaskInfoByID(ctx context.Context, taskID int) (TaskBaseInfo, error) {
	var taskInfo TaskBaseInfo

	if err := ts.db.QueryRowContext(ctx, queryGetTaskInfoByID, taskID).Scan(
		&taskInfo.TrainingID,
		&taskInfo.EmployeeID,
		&taskInfo.ExecutorID,
		&taskInfo.PositionID,
	); err != nil {
		return TaskBaseInfo{}, err
	}

	return taskInfo, nil
}
