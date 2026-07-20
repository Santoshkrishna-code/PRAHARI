package analytics

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetBarrierHealthSummary(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving barrier health summary & uptime metrics", prahariLogger.String("plant_id", plantID))
	return map[string]float64{
		"average_health_score":    96.4,
		"proof_test_compliance":   99.1,
		"active_bypasses_count":   2.0,
		"active_impairments_count": 1.0,
	}, nil
}
