package domain

// const (
// 	InitialBriefingID      = 1 // id первичного иструктажа
// 	RefresherBriefingID    = 2 // id повторного иструктажа
// 	IntroductoryBriefingID = 3 // id вводного иструктажа
// )

type BriefingInfo struct {
	TrainingType TrainingType
	Instructor   Employee
	Act          string
}
