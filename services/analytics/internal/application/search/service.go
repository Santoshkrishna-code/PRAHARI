package search

import (
	"context"

	"prahari/services/analytics/internal/domain/metric"
	"prahari/services/analytics/internal/domain/search"
)

type Repository interface {
	SearchMetrics(ctx context.Context, criteria *search.Criteria) ([]*metric.Metric, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*metric.Metric, int64, error) {
	return s.repo.SearchMetrics(ctx, criteria)
}
