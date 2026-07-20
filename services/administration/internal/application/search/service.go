package search

import (
	"context"

	"prahari/services/administration/internal/domain/search"
	"prahari/services/administration/internal/domain/tenant"
)

type Repository interface {
	SearchTenants(ctx context.Context, criteria *search.Criteria) ([]*tenant.Tenant, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*tenant.Tenant, int64, error) {
	return s.repo.SearchTenants(ctx, criteria)
}
