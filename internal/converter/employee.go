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

func EmployeePersonalCardToDesc(e *domain.EmployeePersonalCard) *desc.GetEmployeePersonalCardResponse {
	result := &desc.GetEmployeePersonalCardResponse{
		Fullname:       e.FullName,
		Birthdate:      e.BirthDate,
		Snils:          e.Snils,
		Department:     e.Department,
		Position:       e.Position,
		EmploymentDate: e.EmploymentDate,
	}

	for _, training := range e.Trainings {
		result.Trainings = append(result.Trainings, TrainingToDesc(training))
	}

	return result
}

func EmployeesInfoToDesc(e []domain.EmployeeInfo) *desc.GetEmployeesByFiltersResponse {
	employees := make([]*desc.EmployeeInfo, 0, len(e))

	for _, empl := range e {
		employee := &desc.EmployeeInfo{
			Fullname:   empl.FullName,
			Department: empl.Department,
			Position:   empl.Position,
			Trainings:  make([]*desc.Training, 0, len(empl.Trainings)),
		}

		for _, t := range empl.Trainings {
			training := &desc.Training{
				Name:        t.Name,
				PassDate:    t.PassDate,
				RePassDate:  t.RePassDate,
				HasProtocol: t.HasProtocol,
			}

			employee.Trainings = append(employee.Trainings, training)
		}

		employees = append(employees, employee)
	}

	return &desc.GetEmployeesByFiltersResponse{Employees: employees}
}
