package domain

type Training struct {
	Name        string
	HasProtocol string
	TrainingDates
}

type TrainingBaseInfo struct {
	ID   int
	Name string
}

type TrainingDates struct {
	PassDate   string
	RePassDate string
}
