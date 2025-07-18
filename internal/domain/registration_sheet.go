package domain

import "time"

type RegistrationSheetInfo struct {
	TrainingType            TrainingType
	EmployeeName            string
	EmployeeBirthDate       time.Time
	EmployeePosition        string
	EmployeeDepartment      string
	InstructorName          string
	InstructorPosition      string
	OccupSafetySpecName     string
	OccupSafetySpecPosition string
	Acts                    string
}
