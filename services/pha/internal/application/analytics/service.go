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

func (s *Service) GetSILDistribution(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving SIL distribution and risk reduction metrics", prahariLogger.String("plant_id", plantID))
	return map[string]float64{
		"sil_1_count": 28.0,
		"sil_2_count": 14.0,
		"sil_3_count": 3.0,
		"sil_4_count": 0.0,
	}, nil
}
