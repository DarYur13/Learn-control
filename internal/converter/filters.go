package converter

import (
	"database/sql"

	"github.com/DarYur13/learn-control/internal/domain"
	pb "github.com/DarYur13/learn-control/pkg/learn_control"
)

func PbFiltersToDomain(req *pb.GetEmployeesByFiltersRequest) domain.Filters {
	filters := domain.Filters{
		Department: sql.NullString{
			Valid:  req.Department != nil,
			String: req.GetDepartment(),
		},
		Position: sql.NullString{
			Valid:  req.Position != nil,
			String: req.GetPosition(),
		},
		TrainingID: sql.NullInt64{
			Valid: req.TrainingID != nil,
			Int64: req.GetTrainingID(),
		},
		DateFrom: sql.NullTime{
			Valid: req.DateFrom.IsValid(),
			Time:  req.GetDateFrom().AsTime(),
		},
		DateTo: sql.NullTime{
			Valid: req.DateTo.IsValid(),
			Time:  req.GetDateTo().AsTime(),
		},
		TrainingsNotPassed: sql.NullBool{
			Valid: req.TrainigsNotPassed != nil,
			Bool:  req.GetTrainigsNotPassed(),
		},
		RetrainingIn: sql.NullInt64{
			Valid: req.RetrainingIn != nil,
			Int64: req.GetRetrainingIn(),
		},
		HasProtocol: sql.NullBool{
			Valid: req.HasProtocol != nil,
			Bool:  req.GetHasProtocol(),
		},
	}

	return filters
}
