package converter

import (
	"github.com/DarYur13/learn-control/internal/domain"
	desc "github.com/DarYur13/learn-control/pkg/learn_control"
)

func EmployeeBaseInfoToDesc(e *domain.EmployeeBaseInfo) *desc.EmployeeBaseInfo {
	return &desc.EmployeeBaseInfo{
		Id:        e.ID,
		Fullname:  e.FullName,
		Birthdate: e.BirthDate,
	}
}

func EmployeesBaseInfoToDesc(e []domain.EmployeeBaseInfo) *desc.GetEmployeesByNameResponse {
	var result []*desc.EmployeeBaseInfo

	for _, employee := range e {
		result = append(result, EmployeeBaseInfoToDesc(&employee))
	}

	return &desc.GetEmployeesByNameResponse{Employees: result}
}

func EmployeePersonalCardToDesc(e *domain.EmployeePersonalCard) *desc.GetEmployeeResponse {
	result := &desc.GetEmployeeResponse{
		Fullname:   e.FullName,
		Birthdate:  e.BirthDate,
		Snils:      e.Snils,
		Department: e.Department,
		Position:   e.Position,
	}

	for _, training := range e.Trainings {
		result.Trainings = append(result.Trainings, TrainingToDesc(training))
	}

	return result
}
