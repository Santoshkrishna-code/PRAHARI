package reporting

import (
	"context"

	"prahari/services/bcm/internal/domain/continuityplan"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetPlanByID(ctx context.Context, id string) (*continuityplan.Plan, error)
	ListPlans(ctx context.Context, plantID string) ([]*continuityplan.Plan, error)
	GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetPlan(ctx context.Context, id string) (*continuityplan.Plan, error) {
	return s.repo.GetPlanByID(ctx, id)
}

func (s *Service) ListPlans(ctx context.Context, plantID string) ([]*continuityplan.Plan, error) {
	return s.repo.ListPlans(ctx, plantID)
}

func (s *Service) GetExecutiveMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving executive BCM dashboard metrics", prahariLogger.String("plant_id", plantID))
	return s.repo.GetDashboardMetrics(ctx, plantID)
}
