package service

import (
	"context"
)

func (s *Service) GetPositions(ctx context.Context) ([]string, error) {
	return s.storage.GetPositions(ctx)
}
