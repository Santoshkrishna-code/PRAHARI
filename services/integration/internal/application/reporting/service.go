package reporting

import (
	"context"

	"prahari/services/integration/internal/domain/connector"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetConnectorByID(ctx context.Context, id string) (*connector.Connector, error)
	ListConnectors(ctx context.Context, plantID string) ([]*connector.Connector, error)
	GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetConnector(ctx context.Context, id string) (*connector.Connector, error) {
	return s.repo.GetConnectorByID(ctx, id)
}

func (s *Service) ListConnectors(ctx context.Context, plantID string) ([]*connector.Connector, error) {
	return s.repo.ListConnectors(ctx, plantID)
}

func (s *Service) GetExecutiveMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving Integration Hub metrics",
		prahariLogger.String("plant_id", plantID))
	return s.repo.GetDashboardMetrics(ctx, plantID)
}
