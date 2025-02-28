package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/converter"
	desc "github.com/DarYur13/learn-control/pkg/learn_control"
)

func (i *Implementation) GetFilters(ctx context.Context, req *desc.GetFiltersRequest) (*desc.GetFiltersResponse, error) {
	filters, err := i.learnControlSrv.GetFilters(ctx)
	if err != nil {
		return nil, err
	}

	return converter.FiltersToDesc(filters), nil
}
