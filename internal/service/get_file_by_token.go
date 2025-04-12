package service

import (
	"context"
	"database/sql"
	"io"
	"time"

	"github.com/DarYur13/learn-control/internal/domain"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (s *Service) GetFileByToken(ctx context.Context, token uuid.UUID) (io.Reader, error) {
	var registrationSheetInfo *domain.RegistrationSheetInfo

	s.txManager.Do(ctx, func(tx *sql.Tx) error {
		tokenInfo, err := s.downloadTokensStorage.GetTokenInfoTx(ctx, tx, token)
		if err != nil {
			return errors.WithMessage(err, "get token info")
		}

		if _, err := s.employeesStorage.UpdateEmployeeTrainingDateTx(
			ctx,
			tx,
			tokenInfo.EmployeeID,
			tokenInfo.TrainingID,
			time.Now(),
		); err != nil {
			return errors.WithMessage(err, "update employee training date")
		}

		registrationSheetInfo, err = s.downloadTokensStorage.GetRegistrationSheetInfoTx(ctx, tx, token)
		if err != nil {
			return errors.WithMessage(err, "get registration sheet info")
		}

		return nil
	})

	return s.docsGenerator.GenerateRegistrationSheet(ctx, *registrationSheetInfo)
}
