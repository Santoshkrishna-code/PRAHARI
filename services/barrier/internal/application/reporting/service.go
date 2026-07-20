package reporting

import (
	"context"

	"prahari/services/barrier/internal/domain/barrier"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetBarrierByID(ctx context.Context, id string) (*barrier.Barrier, error)
	ListBarriers(ctx context.Context, plantID string) ([]*barrier.Barrier, error)
	GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetBarrier(ctx context.Context, id string) (*barrier.Barrier, error) {
	return s.repo.GetBarrierByID(ctx, id)
}

func (s *Service) ListBarriers(ctx context.Context, plantID string) ([]*barrier.Barrier, error) {
	return s.repo.ListBarriers(ctx, plantID)
}

func (s *Service) GetExecutiveMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving executive barrier dashboard metrics", prahariLogger.String("plant_id", plantID))
	return s.repo.GetDashboardMetrics(ctx, plantID)
}
