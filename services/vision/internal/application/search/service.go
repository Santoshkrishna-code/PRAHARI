package search

import (
	"context"

	"prahari/services/vision/internal/domain/detection"
	"prahari/services/vision/internal/domain/search"
)

type Repository interface {
	SearchDetections(ctx context.Context, criteria *search.Criteria) ([]*detection.Detection, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*detection.Detection, int64, error) {
	return s.repo.SearchDetections(ctx, criteria)
}
