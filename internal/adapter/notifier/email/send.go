package notifier

import (
	"context"
	"net/smtp"

	"github.com/DarYur13/learn-control/internal/domain"
	"github.com/jordan-wright/email"
)

func (n *notifier) Send(ctx context.Context, notification domain.SMTPNotification) error {
	e := email.NewEmail()
	e.From = n.From
	e.To = []string{notification.Recipient}
	e.Subject = notification.Subject
	e.Text = []byte(notification.Body)

	auth := smtp.PlainAuth("", n.From, n.Password, n.Host)
	return e.Send(n.Host+":"+n.Port, auth)
}
