package reporting

import (
	"context"

	"prahari/services/ppe/internal/domain/ppe"
	"prahari/services/ppe/internal/domain/ppeitem"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetPPEByID(ctx context.Context, id string) (*ppe.PPE, error)
	GetPPEItemByID(ctx context.Context, id string) (*ppeitem.Item, error)
	ListPPEs(ctx context.Context, plantID string) ([]*ppe.PPE, error)
	GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetPPE(ctx context.Context, id string) (*ppe.PPE, error) {
	return s.repo.GetPPEByID(ctx, id)
}

func (s *Service) GetPPEItem(ctx context.Context, id string) (*ppeitem.Item, error) {
	return s.repo.GetPPEItemByID(ctx, id)
}

func (s *Service) ListPPEs(ctx context.Context, plantID string) ([]*ppe.PPE, error) {
	return s.repo.ListPPEs(ctx, plantID)
}

func (s *Service) GetExecutiveMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving protective gear compliance inventory dashboard metrics", prahariLogger.String("plant_id", plantID))
	return s.repo.GetDashboardMetrics(ctx, plantID)
}
