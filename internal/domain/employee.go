package domain

import (
	desc "github.com/DarYur13/learn-control/pkg/learn_control"
)

type EmployeeBaseInfo struct {
	ID        int64
	FullName  string
	BirthDate string
}

func (e *EmployeeBaseInfo) ToDesc() *desc.EmployeeBaseInfo {
	return &desc.EmployeeBaseInfo{
		Id:        e.ID,
		Fullname:  e.FullName,
		Birthdate: e.BirthDate,
	}
}

type EmployeesBaseInfo struct {
	Employees []EmployeeBaseInfo
}

func (e *EmployeesBaseInfo) ToDesc() *desc.GetEmployeesResponse {
	var result []*desc.EmployeeBaseInfo

	for _, employee := range e.Employees {
		result = append(result, employee.ToDesc())
	}

	return &desc.GetEmployeesResponse{Employees: result}
}

type Employee struct {
	FullName   string
	BirthDate  string
	Snils      string
	Department string
	Position   string
	Trainings  []Training
}

func (e *Employee) ToDesc() *desc.GetEmployeeResponse {
	result := &desc.GetEmployeeResponse{
		Fullname:   e.FullName,
		Birthdate:  e.BirthDate,
		Snils:      e.Snils,
		Department: e.Department,
		Position:   e.Position,
	}

	for _, training := range e.Trainings {
		result.Trainings = append(result.Trainings, training.ToDesc())
	}

	return result
}
