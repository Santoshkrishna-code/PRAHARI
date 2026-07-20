package benchmarking

import (
	"context"

	"prahari/services/analytics/internal/domain/benchmark"
)

type Repository interface {
	GetBenchmark(ctx context.Context, metricKey, plantID string) (*benchmark.Comparison, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ComparePlantPerformance(ctx context.Context, metricKey, plantID string) (*benchmark.Comparison, error) {
	return s.repo.GetBenchmark(ctx, metricKey, plantID)
}
