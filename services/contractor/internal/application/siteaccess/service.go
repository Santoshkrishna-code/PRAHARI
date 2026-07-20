package siteaccess

import (
	"context"

	accessDomain "prahari/services/contractor/internal/domain/siteaccess"
)

// Repository persistence gates entries.
type Repository interface {
	Create(ctx context.Context, sa *accessDomain.SiteAccess) error
	FindByWorkerID(ctx context.Context, workerID string) ([]*accessDomain.SiteAccess, error)
}

// Service manages access authorizations.
type Service struct {
	repo Repository
}

// NewService instantiates SiteAccess Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// GrantAccess persists entry check stamps.
func (s *Service) GrantAccess(ctx context.Context, sa *accessDomain.SiteAccess) (*accessDomain.SiteAccess, error) {
	if err := sa.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, sa); err != nil {
		return nil, err
	}
	return sa, nil
}
