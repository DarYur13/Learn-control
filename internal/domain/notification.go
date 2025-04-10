package domain

import (
	"io"
	"time"
)

type NotificationType string

const (
	NotificationTypeIntroBrief         NotificationType = "INTRO_BRIEF"
	NotificationTypeInitBrief          NotificationType = "INIT_BRIEF"
	NotificationTypeRefreshBriefFirst  NotificationType = "REFRESH_BRIEF_FIRST"
	NotificationTypeRefreshBriefSecond NotificationType = "REFRESH_BRIEF_SECOND"
)

type SMTPNotification struct {
	Recipient string
	Subject   string
	Body      string
	File      io.Reader
	Filename  string
}

type SMTPNotificationTemplate struct {
	Subject string
	Body    string
}

type PendingNotification struct {
	ID                 int
	Type               NotificationType
	Subject            string
	Body               string
	InstructorName     string
	InstructorEmail    string
	InstructorPosition string
	EmployeeName       string
	EmployeeBirthDate  time.Time
	EmployeePosition   string
	EmployeeDepartment string
	TrainingType       TrainingType
	Acts               string
	ReTrainingDate     time.Time
}
