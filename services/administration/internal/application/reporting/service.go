package reporting

import (
	"context"

	"prahari/services/administration/internal/domain/tenant"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetTenantByID(ctx context.Context, id string) (*tenant.Tenant, error)
	ListTenants(ctx context.Context) ([]*tenant.Tenant, error)
	GetDashboardMetrics(ctx context.Context) (map[string]float64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetTenant(ctx context.Context, id string) (*tenant.Tenant, error) {
	return s.repo.GetTenantByID(ctx, id)
}

func (s *Service) ListTenants(ctx context.Context) ([]*tenant.Tenant, error) {
	return s.repo.ListTenants(ctx)
}

func (s *Service) GetExecutiveMetrics(ctx context.Context) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving platform administrative dashboard metrics")
	return s.repo.GetDashboardMetrics(ctx)
}
