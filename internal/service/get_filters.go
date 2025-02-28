package service

// func (s *Service) GetFilters(ctx context.Context) (*domain.Filters, error) {
// 	departments, err := s.storage.GetDepartments(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	positions, err := s.storage.GetPositions(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	trainings, err := s.storage.GetTrainings(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	filters := domain.Filters{
// 		Deparments: departments,
// 		Positions:  positions,
// 	}

// 	for _, training := range trainings {
// 		filters.Trainings = append(filters.Trainings, domain.TrainingBaseInfo(training))
// 	}

// 	return &filters, nil
// }
