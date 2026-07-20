package reporting

import (
	"context"

	"prahari/services/moc/internal/domain/changerequest"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetRequestByID(ctx context.Context, id string) (*changerequest.Request, error)
	ListRequests(ctx context.Context, plantID string) ([]*changerequest.Request, error)
	GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetRequest(ctx context.Context, id string) (*changerequest.Request, error) {
	return s.repo.GetRequestByID(ctx, id)
}

func (s *Service) ListRequests(ctx context.Context, plantID string) ([]*changerequest.Request, error) {
	return s.repo.ListRequests(ctx, plantID)
}

func (s *Service) GetExecutiveMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving executive MOC dashboard metrics", prahariLogger.String("plant_id", plantID))
	return s.repo.GetDashboardMetrics(ctx, plantID)
}
