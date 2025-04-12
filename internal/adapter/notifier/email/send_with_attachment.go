package notifier

import (
	"context"
	"fmt"
	"io"
	"net/smtp"
	"os"

	"github.com/DarYur13/learn-control/internal/domain"
	"github.com/jordan-wright/email"
)

func (n *notifier) SendWithAttachment(ctx context.Context, notification domain.SMTPNotification) error {
	e := email.NewEmail()
	e.From = n.From
	e.To = []string{notification.Recipient}
	e.Subject = notification.Subject
	e.Text = []byte(notification.Body)

	// создаём временный файл, чтобы прикрепить
	tmpFile, err := os.CreateTemp("", "notif-*")
	if err != nil {
		return fmt.Errorf("create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	_, err = io.Copy(tmpFile, notification.File)
	if err != nil {
		return fmt.Errorf("copy to temp file: %w", err)
	}

	if _, err := tmpFile.Seek(0, 0); err != nil {
		return err
	}

	if _, err := e.Attach(tmpFile, notification.Filename, "application/vnd.openxmlformats-officedocument.wordprocessingml.document"); err != nil {
		return fmt.Errorf("attach file: %w", err)
	}

	auth := smtp.PlainAuth("", n.From, n.Password, n.Host)

	return e.Send(n.Host+":"+n.Port, auth)
}
