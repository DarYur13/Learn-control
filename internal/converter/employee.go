package converter

import (
	"github.com/DarYur13/learn-control/internal/domain"
	pb "github.com/DarYur13/learn-control/pkg/learn_control"
)

func EmployeeBaseInfoToDesc(e *domain.EmployeeBaseInfo) *pb.EmployeeBaseInfo {
	return &pb.EmployeeBaseInfo{
		Id:        e.ID,
		Fullname:  e.FullName,
		Birthdate: e.BirthDate,
	}
}

func EmployeesBaseInfoToDesc(e []domain.EmployeeBaseInfo) *pb.GetEmployeesByNameResponse {
	var result []*pb.EmployeeBaseInfo

	for _, employee := range e {
		result = append(result, EmployeeBaseInfoToDesc(&employee))
	}

	return &pb.GetEmployeesByNameResponse{Employees: result}
}

func EmployeePersonalCardToDesc(e *domain.EmployeePersonalCard) *pb.GetEmployeePersonalCardResponse {
	result := &pb.GetEmployeePersonalCardResponse{
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

func EmployeesInfoToDesc(e []domain.EmployeeInfo) *pb.GetEmployeesByFiltersResponse {
	employees := make([]*pb.EmployeeInfo, 0, len(e))

	for _, empl := range e {
		employee := &pb.EmployeeInfo{
			Fullname:   empl.FullName,
			Department: empl.Department,
			Position:   empl.Position,
			Trainings:  make([]*pb.Training, 0, len(empl.Trainings)),
		}

		for _, t := range empl.Trainings {
			training := &pb.Training{
				Name:        t.Name,
				PassDate:    t.PassDate,
				RePassDate:  t.RePassDate,
				HasProtocol: t.HasProtocol,
			}

			employee.Trainings = append(employee.Trainings, training)
		}

		employees = append(employees, employee)
	}

	return &pb.GetEmployeesByFiltersResponse{Employees: employees}
}
