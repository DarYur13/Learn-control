package learncontrol

import (
	"context"
	"database/sql"

	"github.com/DarYur13/learn-control/internal/converter"
	pb "github.com/DarYur13/learn-control/pkg/learn_control"
)

func (i *Implementation) GetTasksByFilters(ctx context.Context, req *pb.GetTasksByFiltersRequest) (*pb.GetTasksByFiltersResponse, error) {
	done := sql.NullBool{}

	if req.Done != nil {
		done.Valid = true
		done.Bool = req.GetDone()
	}

	tasks, err := i.service.GetTasksByFilters(ctx, done)
	if err != nil {
		return nil, err
	}

	return converter.TasksToPb(tasks), nil
}
