package analytics

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type MetricsRepository interface {
	GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error)
}

type Service struct {
	repo MetricsRepository
}

func NewService(repo MetricsRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetExecutiveMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving executive water dashboard metrics", prahariLogger.String("plant_id", plantID))
	return s.repo.GetDashboardMetrics(ctx, plantID)
}
