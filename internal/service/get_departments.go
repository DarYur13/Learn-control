package service

import (
	"context"
)

func (s *Service) GetDepartments(ctx context.Context) ([]string, error) {
	return s.positionsStorage.GetPositionsDepartments(ctx)
}
