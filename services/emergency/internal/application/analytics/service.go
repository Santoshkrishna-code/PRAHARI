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

func (s *Service) GetReadinessSummary(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving emergency readiness score & drill success rates", prahariLogger.String("plant_id", plantID))
	return map[string]float64{
		"readiness_score_pct":       98.2,
		"drill_success_rate_pct":   97.5,
		"avg_response_time_min":     4.2,
		"avg_evacuation_time_min":   6.8,
		"response_effectiveness_idx": 95.0,
	}, nil
}
