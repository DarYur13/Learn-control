package domain

import "io"

type NotificationType string

const (
	NotificationTypeIntroBrief         NotificationType = "INTRO_BRIEF"
	NotificationTypeInitBrief          NotificationType = "INIT_BRIEF"
	NotificationTypeRefreshBriefFirst  NotificationType = "REFRESH_BRIEF_FIRST"
	NotificationTypeRefreshBriefSecond NotificationType = "REFRESH_BRIEF_SECOND"
)

type SMTPNotification struct {
	Type      NotificationType
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
	ID         int
	EmployeeID int
	TrainingID int
	Type       NotificationType
}
