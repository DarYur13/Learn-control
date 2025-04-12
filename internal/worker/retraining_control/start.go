package worker

import (
	"context"
	"fmt"
	"time"

	tasksRepo "github.com/DarYur13/learn-control/internal/adapter/repository/learn_control/tasks"
	"github.com/DarYur13/learn-control/internal/domain"
	"github.com/DarYur13/learn-control/internal/logger"
	"github.com/pkg/errors"
)

func (rcw *retrainingControlWorker) Start(ctx context.Context) {
	ticker := time.NewTicker(rcw.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			logger.Info("🛑 Retraining control worker stopped")
			return
		case <-ticker.C:
			if err := rcw.processUpcomingTrainings(ctx); err != nil {
				logger.Errorf("⚠️ Failed to process upcoming trainings: %v", err)
			}
		}
	}
}

// TODO transactions
func (rcw *retrainingControlWorker) processUpcomingTrainings(ctx context.Context) error {
	trainings, err := rcw.trainingsRepo.GetUpcomingTrainings(ctx)
	if err != nil {
		return errors.WithMessage(err, "failed to get upcoming trainings")
	}

	for _, training := range trainings {
		if training.TrainingType == domain.TrainingTypeInitial {
			training.TrainingID, err = rcw.employeesRepo.SetEmployeeRefresherBrief(ctx, training.EmployeeID)
			if err != nil {
				logger.Errorf("⚠️ Failed to set refresher brief to employee id=%d: %s", training.EmployeeID, err)
				continue
			}

			if err := rcw.processRefresherBrief(ctx, training); err != nil {
				logger.Error("⚠️ Failed to process refresher brief for employee id=%d: %s", training.EmployeeID, err)
				continue
			}

			continue
		}

		if training.TrainingType == domain.TrainingTypeRefresher {
			if err := rcw.processRefresherBrief(ctx, training); err != nil {
				logger.Errorf("⚠️ Failed to process refresher brief for employee id=%d: %s", training.EmployeeID, err)
				continue
			}

			continue
		}

		if training.TrainingType == domain.TrainingTypeRegular {
			if err := rcw.processRegularTraining(ctx, training); err != nil {
				logger.Errorf("⚠️ Failed to process regular training for employee id=%d: %s", training.EmployeeID, err)
				continue
			}

			continue
		}

		logger.Errorf("⚠️ Unknown or unsupported training type: %s: %s", training.TrainingType)
	}

	return nil
}

func (rcw *retrainingControlWorker) processRefresherBrief(ctx context.Context, brief domain.UpcomingTraining) error {
	executorID, err := rcw.employeesRepo.GetEmployeeLeader(ctx, brief.EmployeeID)
	if err != nil {
		return errors.WithMessage(err, "failed to get employee leader")
	}

	task, err := rcw.service.CreateControlTask(ctx, brief.EmployeeID, brief.TrainingID, executorID)
	if err != nil {
		return errors.WithMessage(err, "failed to create task")
	}

	var notificationType domain.NotificationType

	switch brief.DaysLeft {
	case daysTillRetrainingFirst:
		notificationType = domain.NotificationTypeRefreshBriefFirst

	case daysTillRetrainingSecond:
		notificationType = domain.NotificationTypeRefreshBriefSecond

	default:
		return fmt.Errorf("invalid days left: %d, want: %d or %d", brief.DaysLeft, daysTillRetrainingFirst, daysTillRetrainingSecond)
	}

	if err := rcw.notificationsRepo.AddNotificationToQueue(ctx, brief.EmployeeID, brief.TrainingID, notificationType); err != nil {
		return errors.WithMessage(err, "failed to add notification to queue")
	}

	if err := rcw.tasksRepo.AddTask(ctx, tasksRepo.TaskBaseInfo(*task)); err != nil {
		return errors.WithMessage(err, "failed to add task")
	}

	return nil
}

func (rcw *retrainingControlWorker) processRegularTraining(ctx context.Context, training domain.UpcomingTraining) error {
	if training.DaysLeft == daysTillRetrainingFirst {
		task, err := rcw.service.CreateAssignTask(ctx, training.EmployeeID, training.TrainingID)
		if err != nil {
			return errors.WithMessage(err, "failed to create task")
		}

		if err := rcw.tasksRepo.AddTask(ctx, tasksRepo.TaskBaseInfo(*task)); err != nil {
			return errors.WithMessage(err, "failed to add task")
		}
	}

	return nil
}
