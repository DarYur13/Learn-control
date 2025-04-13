package domain

import "time"

type Employee struct {
	FullName       string
	BirthDate      time.Time
	Snils          string
	Department     string
	Position       string
	EmploymentDate time.Time
	Email          string
}

type EmployeeBaseInfo struct {
	ID        int64
	FullName  string
	BirthDate time.Time
}

type EmployeePersonalCard struct {
	FullName       string
	BirthDate      time.Time
	Snils          string
	Department     string
	Position       string
	EmploymentDate time.Time
	Trainings      []Training
}

type EmployeeInfo struct {
	FullName   string
	Department string
	Position   string
	Trainings  []Training
}
