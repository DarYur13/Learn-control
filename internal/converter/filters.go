package converter

import (
	"github.com/DarYur13/learn-control/internal/domain"
	desc "github.com/DarYur13/learn-control/pkg/learn_control"
)

func FiltersToDesc(filters *domain.Filters) *desc.GetFiltersResponse {
	result := &desc.GetFiltersResponse{
		Departments: filters.Deparments,
		Positions:   filters.Positions,
	}

	for _, t := range filters.Trainings {
		training := desc.TrainingBaseInfo{
			Id:   int64(t.ID),
			Name: t.Name,
		}
		result.Trainings = append(result.Trainings, &training)
	}

	return result
}
