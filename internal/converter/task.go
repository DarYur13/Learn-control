package converter

import (
	"github.com/DarYur13/learn-control/internal/domain"
	desc "github.com/DarYur13/learn-control/pkg/learn_control"
)

func TasksToDesc(tasks []domain.Task) *desc.GetTasksByFiltersResponse {
	result := make([]*desc.Task, 0, len(tasks))

	for _, t := range tasks {
		task := &desc.Task{
			Type:        TypeToDesc(t.Type),
			Description: t.Description,
		}

		if t.Employee.Valid {
			task.Employee = t.Employee.String
		}

		if t.Training.Valid {
			task.Training = t.Training.String
		}

		if t.Position.Valid {
			task.Position = t.Position.String
		}

		if t.Department.Valid {
			task.Department = t.Department.String
		}

		if t.Executor.Valid {
			task.Executor = t.Executor.String
		}

		result = append(result, task)
	}

	return &desc.GetTasksByFiltersResponse{Tasks: result}
}

func TypeToDesc(taskType string) desc.TaskType {
	if value, found := desc.TaskType_value[taskType]; found {
		return desc.TaskType(value)
	}

	return desc.TaskType_UNSPECIFIED
}
