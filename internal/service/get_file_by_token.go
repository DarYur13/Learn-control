package service

import (
	"context"
	"io"

	"github.com/google/uuid"
)

func (s *Service) GetFileByToken(ctx context.Context, token uuid.UUID) (io.Reader, error) {
	registrationSheetInfo, err := s.downloadTokensStorage.GetRegistrationSheetInfo(ctx, token)
	if err != nil {
		return nil, err
	}

	file, err := s.docsGenerator.GenerateRegistrationSheet(ctx, *registrationSheetInfo)
	if err != nil {
		return nil, err
	}

	return file, nil
}
