package domain

type Employee struct {
	FullName       string
	BirthDate      string
	Snils          string
	Department     string
	Position       string
	EmploymentDate string
}

type EmployeeBaseInfo struct {
	ID        int64
	FullName  string
	BirthDate string
}

type EmployeePersonalCard struct {
	FullName       string
	BirthDate      string
	Snils          string
	Department     string
	Position       string
	EmploymentDate string
	Trainings      []Training
}

type EmployeeInfo struct {
	FullName   string
	Department string
	Position   string
	Trainings  []Training
}
