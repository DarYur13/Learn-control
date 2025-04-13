package converter

import (
	"github.com/DarYur13/learn-control/internal/domain"
	pb "github.com/DarYur13/learn-control/pkg/learn_control"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func EmployeeBaseInfoToPb(e *domain.EmployeeBaseInfo) *pb.EmployeeBaseInfo {
	return &pb.EmployeeBaseInfo{
		Id:        e.ID,
		Fullname:  e.FullName,
		Birthdate: timestamppb.New(e.BirthDate),
	}
}

func EmployeesBaseInfoToPb(e []domain.EmployeeBaseInfo) *pb.GetEmployeesByNameResponse {
	var result []*pb.EmployeeBaseInfo

	for _, employee := range e {
		result = append(result, EmployeeBaseInfoToPb(&employee))
	}

	return &pb.GetEmployeesByNameResponse{Employees: result}
}

func EmployeePersonalCardToPb(e *domain.EmployeePersonalCard) *pb.GetEmployeePersonalCardResponse {
	result := &pb.GetEmployeePersonalCardResponse{
		Fullname:       e.FullName,
		Birthdate:      timestamppb.New(e.BirthDate),
		Snils:          e.Snils,
		Department:     e.Department,
		Position:       e.Position,
		EmploymentDate: timestamppb.New(e.EmploymentDate),
	}

	for _, training := range e.Trainings {
		result.Trainings = append(result.Trainings, TrainingToPb(training))
	}

	return result
}

func EmployeesInfoToPb(e []domain.EmployeeInfo) *pb.GetEmployeesByFiltersResponse {
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
				PassDate:    timestamppb.New(t.PassDate),
				RePassDate:  timestamppb.New(t.RePassDate),
				HasProtocol: t.HasProtocol,
			}

			employee.Trainings = append(employee.Trainings, training)
		}

		employees = append(employees, employee)
	}

	return &pb.GetEmployeesByFiltersResponse{Employees: employees}
}

func PbEmployeeToDomain(req *pb.AddEmployeeRequest) domain.Employee {
	return domain.Employee{
		FullName:       req.GetFullname(),
		BirthDate:      req.GetBirthdate().AsTime(),
		Snils:          req.GetSnils(),
		Department:     req.GetDepartment(),
		Position:       req.GetPosition(),
		EmploymentDate: req.GetEmploymentDate().AsTime(),
		Email:          req.GetEmail(),
	}
}
