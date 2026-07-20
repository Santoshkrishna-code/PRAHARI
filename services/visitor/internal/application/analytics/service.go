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

func (s *Service) GetAccessSummary(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving access control throughput and compliance metrics", prahariLogger.String("plant_id", plantID))
	return map[string]float64{
		"active_visitors_count":        42.0,
		"average_visit_duration_hrs":   4.2,
		"average_checkin_time_sec":     95.0,
		"badge_return_compliance_pct":  99.8,
		"induction_compliance_pct":     100.0,
		"visitor_throughput_daily":     124.0,
		"gate_processing_time_sec":     48.0,
	}, nil
}
