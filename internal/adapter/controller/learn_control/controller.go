package learncontrol

import (
	"github.com/DarYur13/learn-control/internal/service"
	pb "github.com/DarYur13/learn-control/pkg/learn_control"
)

type Implementation struct {
	pb.UnimplementedLearnControlServer

	service service.Servicer
}

func NewLearnControl(service service.Servicer) *Implementation {
	return &Implementation{
		service: service,
	}
}
