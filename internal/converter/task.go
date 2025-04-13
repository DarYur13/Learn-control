package converter

import (
	"fmt"

	"github.com/DarYur13/learn-control/internal/domain"
	pb "github.com/DarYur13/learn-control/pkg/learn_control"
)

func TasksToPb(tasks []domain.Task) *pb.GetTasksByFiltersResponse {
	result := make([]*pb.Task, 0, len(tasks))

	for _, t := range tasks {
		task := &pb.Task{
			Id:          int64(t.ID),
			Type:        DomainTaskTypeToPb(t.Type),
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

		if t.FileLink.Valid {
			task.DownloadFileLink = t.FileLink.String
		}

		result = append(result, task)
	}

	return &pb.GetTasksByFiltersResponse{Tasks: result}
}

func DomainTaskTypeToPb(taskType domain.TaskType) pb.TaskType {
	if value, found := pb.TaskType_value[string(taskType)]; found {
		return pb.TaskType(value)
	}

	return pb.TaskType_UNKNOWN_TASK
}

func PbTaskTypeToDomain(pbTaskType pb.TaskType) (domain.TaskType, error) {
	switch pbTaskType {
	case pb.TaskType_ASSIGN:
		return domain.TaskTypeAssign, nil
	case pb.TaskType_SET:
		return domain.TaskTypeSet, nil
	case pb.TaskType_CONFIRM:
		return domain.TaskTypeConfirm, nil
	case pb.TaskType_PROVIDE:
		return domain.TaskTypeProvide, nil
	case pb.TaskType_CONTROL:
		return domain.TaskTypeControl, nil
	case pb.TaskType_CHOOSE:
		return domain.TaskTypeChoose, nil
	default:
		return "", fmt.Errorf("unsupported task type: %v", pbTaskType)
	}
}
