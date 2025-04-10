package docsgenerator

import (
	"context"
	"io"

	"github.com/DarYur13/learn-control/internal/domain"
)

type DocsGenerator interface {
	GenerateRegistrationSheet(ctx context.Context, employee domain.Employee, briefingInfo domain.BriefingInfo) (io.Reader, error)
}
