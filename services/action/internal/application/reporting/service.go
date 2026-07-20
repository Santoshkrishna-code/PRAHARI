package reporting

import (
	"context"

	"prahari/services/action/internal/domain/action"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetActionByID(ctx context.Context, id string) (*action.Action, error)
	ListActions(ctx context.Context, plantID string) ([]*action.Action, error)
	GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAction(ctx context.Context, id string) (*action.Action, error) {
	return s.repo.GetActionByID(ctx, id)
}

func (s *Service) ListActions(ctx context.Context, plantID string) ([]*action.Action, error) {
	return s.repo.ListActions(ctx, plantID)
}

func (s *Service) GetExecutiveMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving continuous improvement CAPA dashboard metrics", prahariLogger.String("plant_id", plantID))
	return s.repo.GetDashboardMetrics(ctx, plantID)
}
