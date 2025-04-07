package converter

import (
	"github.com/DarYur13/learn-control/internal/domain"
	pb "github.com/DarYur13/learn-control/pkg/learn_control"
)

func TasksToDesc(tasks []domain.Task) *pb.GetTasksByFiltersResponse {
	result := make([]*pb.Task, 0, len(tasks))

	for _, t := range tasks {
		task := &pb.Task{
			Id:          int64(t.ID),
			Type:        TypeToDesc(t.Type),
			Description: t.Description,
			Done:        t.Done,
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

	return &pb.GetTasksByFiltersResponse{Tasks: result}
}

func TypeToDesc(taskType string) pb.TaskType {
	if value, found := pb.TaskType_value[taskType]; found {
		return pb.TaskType(value)
	}

	return pb.TaskType_UNSPECIFIED
}
