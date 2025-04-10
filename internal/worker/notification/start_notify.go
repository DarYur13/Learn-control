package worker

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/DarYur13/learn-control/internal/domain"
	"github.com/DarYur13/learn-control/internal/logger"
)

func (nw *notificationWorker) StartNotify(ctx context.Context) {
	ticker := time.NewTicker(nw.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			logger.Info("🛑 Notification worker stopped")
			return
		case <-ticker.C:
			if err := nw.processNotifications(ctx); err != nil {
				logger.Errorf("⚠️ Failed to process notifications: %v", err)
			}
		}
	}
}

func (nw *notificationWorker) processNotifications(ctx context.Context) error {
	notifications, err := nw.notificationsRepo.GetPendingNotifications(ctx)
	if err != nil {
		return err
	}

	for _, n := range notifications {
		if err := nw.handleNotification(ctx, n); err != nil {
			logger.Errorf("❌ Error handling notification (id=%d): %v", n.ID, err)

			continue
		}

		if err := nw.notificationsRepo.MarkNotificationAsSent(ctx, n.ID); err != nil {
			logger.Errorf("❌ Failed to mark notification as sent (id=%d): %v", n.ID, err)
		}
	}

	return nil
}

func (nw *notificationWorker) handleNotification(ctx context.Context, notification domain.PendingNotification) error {
	filledBody := nw.fillBodyTemplate(notification)

	readyToSend := domain.SMTPNotification{
		Recipient: notification.InstructorEmail,
		Subject:   notification.Subject,
		Body:      filledBody,
	}

	if notification.Type == domain.NotificationTypeInitBrief {
		file, err := nw.getFile(ctx, notification)
		if err != nil {
			return err
		}

		readyToSend.File = file
		readyToSend.Filename = fmt.Sprintf("Первичный инструктаж для %s", notification.EmployeeName)

		if err := nw.notifier.SendWithAttachment(ctx, readyToSend); err != nil {
			return err
		}
	} else {
		if err := nw.notifier.Send(ctx, readyToSend); err != nil {
			return err
		}
	}

	logger.Infof("✅ Notification sent: type=%s to=%s", notification.Type, notification.InstructorEmail)

	return nil
}

func (nw *notificationWorker) fillBodyTemplate(notification domain.PendingNotification) string {

	replacer := strings.NewReplacer(
		"{instructor_name}", notification.InstructorName,
		"{employee_name}", notification.EmployeeName,
		"{department}", notification.EmployeeDepartment,
		"{position}", notification.EmployeePosition,
		"{today_date}", time.Now().Format(domain.DateFormat),
		"{retraining_date}", notification.ReTrainingDate.Format(domain.DateFormat),
	)

	return replacer.Replace(notification.Body)
}

func (nw *notificationWorker) getFile(ctx context.Context, notification domain.PendingNotification) (io.Reader, error) {

	info := domain.RegistrationSheetInfo{
		TrainingType:       notification.TrainingType,
		EmployeeName:       notification.EmployeeName,
		EmployeeBirthDate:  notification.EmployeeBirthDate,
		EmployeePosition:   notification.EmployeePosition,
		EmployeeDepartment: notification.EmployeeDepartment,
		InstructorName:     notification.InstructorName,
		InstructorPosition: notification.InstructorPosition,
		Acts:               notification.Acts,
	}

	return nw.docsGenerator.GenerateRegistrationSheet(ctx, info)
}
