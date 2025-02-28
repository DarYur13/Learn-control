package converter

import (
	"github.com/DarYur13/learn-control/internal/domain"
	desc "github.com/DarYur13/learn-control/pkg/learn_control"
)

func TrainingToDesc(t domain.Training) *desc.Training {
	return &desc.Training{
		Name:       t.Name,
		PassDate:   t.PassDate,
		RePassDate: t.RePassDate,
	}
}
