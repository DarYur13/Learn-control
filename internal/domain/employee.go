package domain

type EmployeeBaseInfo struct {
	ID        int64
	FullName  string
	BirthDate string
}

type EmployeePersonalCard struct {
	FullName   string
	BirthDate  string
	Snils      string
	Department string
	Position   string
	Trainings  []Training
}

type EmployeeWorkInfo struct {
	FullName   string
	Department string
	Position   string
	Trainings  []Training
}
