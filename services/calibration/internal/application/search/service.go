package search

import (
	"context"

	"prahari/services/calibration/internal/domain/calibration"
	"prahari/services/calibration/internal/domain/search"
)

type Repository interface {
	SearchCalibrations(ctx context.Context, criteria *search.Criteria) ([]*calibration.Record, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*calibration.Record, int64, error) {
	return s.repo.SearchCalibrations(ctx, criteria)
}
