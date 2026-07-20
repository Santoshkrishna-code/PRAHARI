package reporting

import (
	"context"

	"prahari/services/chemical/internal/domain/chemical"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetChemicalByID(ctx context.Context, id string) (*chemical.Chemical, error)
	ListChemicals(ctx context.Context, plantID string) ([]*chemical.Chemical, error)
	GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetChemical(ctx context.Context, id string) (*chemical.Chemical, error) {
	return s.repo.GetChemicalByID(ctx, id)
}

func (s *Service) ListChemicals(ctx context.Context, plantID string) ([]*chemical.Chemical, error) {
	return s.repo.ListChemicals(ctx, plantID)
}

func (s *Service) GetExecutiveMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving chemical inventory safety metrics",
		prahariLogger.String("plant_id", plantID))
	return s.repo.GetDashboardMetrics(ctx, plantID)
}
