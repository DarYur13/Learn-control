package domain

type TrainingType string

const (
	TrainingTypeRegular      TrainingType = "REGULAR"
	TrainingTypeIntroductory TrainingType = "INTRODUCTORY"
	TrainingTypeInitial      TrainingType = "INITIAL"
	TrainingTypeRefresher    TrainingType = "REFRESHER"
)

type Training struct {
	Name        string
	HasProtocol string
	TrainingDates
}

type TrainingBaseInfo struct {
	ID   int
	Name string
	Type TrainingType
}

type TrainingDates struct {
	PassDate   string
	RePassDate string
}
