package domain

import "time"

type TrainingType string

const (
	TrainingTypeRegular      TrainingType = "REGULAR"
	TrainingTypeIntroductory TrainingType = "INTRODUCTORY"
	TrainingTypeInitial      TrainingType = "INITIAL"
	TrainingTypeRefresher    TrainingType = "REFRESHER"
)

type Training struct {
	Name        string
	Type        TrainingType
	HasProtocol string
	TrainingDates
}

type TrainingBaseInfo struct {
	ID   int
	Name string
	Type TrainingType
}

type TrainingDates struct {
	PassDate   time.Time
	RePassDate time.Time
}

type UpcomingTraining struct {
	EmployeeID   int
	PositionID   int
	TrainingID   int
	TrainingType TrainingType
	RePassDate   time.Time
	DaysLeft     int
}
