package storage

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

const (
	queryGetTasks = `
	SELECT 
		t.task_type,
		ttt.task_text,
		e_empl.full_name AS employee,
		tr.training,
		p.position,
		p.department,
		e_exec.full_name AS executor
	FROM tasks t
	LEFT JOIN task_types_texts ttt ON t.task_type = ttt.task_type
	LEFT JOIN employees e_empl ON t.employee_id = e_empl.id
	LEFT JOIN employees e_exec ON t.executor_id = e_exec.id
	LEFT JOIN trainings tr ON t.training_id = tr.id
	LEFT JOIN positions p ON t.position_id = p.id
	`
)

func (s *Storage) GetTasksByFilters(ctx context.Context, done sql.NullBool) ([]Task, error) {
	var filters []string
	var args []interface{}

	if done.Valid {
		paramIndex := len(args) + 1
		filters = append(filters, fmt.Sprintf("t.done = $%d", paramIndex))
		args = append(args, done.Bool)
	}

	var sb strings.Builder

	sb.WriteString(queryGetTasks)
	if len(filters) > 0 {
		sb.WriteString(" WHERE ")
		sb.WriteString(strings.Join(filters, " AND "))
	}

	queryWithFilters := sb.String()

	var tasks []Task

	rows, err := s.db.QueryContext(ctx, queryWithFilters, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task Task

		if err := rows.Scan(
			&task.Type,
			&task.Description,
			&task.Employee,
			&task.Training,
			&task.Position,
			&task.Department,
			&task.Executor,
		); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
