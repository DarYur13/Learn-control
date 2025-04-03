package learncontrol

import (
	"context"
	"database/sql"

	"github.com/DarYur13/learn-control/internal/converter"
	desc "github.com/DarYur13/learn-control/pkg/learn_control"
)

func (i *Implementation) GetTasksByFilters(ctx context.Context, req *desc.GetTasksByFiltersRequest) (*desc.GetTasksByFiltersResponse, error) {
	done := sql.NullBool{}

	if req.Done != nil {
		done.Valid = true
		done.Bool = req.GetDone()
	}

	tasks, err := i.learnControlSrv.GetTasksByFilters(ctx, done)
	if err != nil {
		return nil, err
	}

	return converter.TasksToDesc(tasks), nil
}
