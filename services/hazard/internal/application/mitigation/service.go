package mitigation

import (
	"context"

	mitigationDomain "prahari/services/hazard/internal/domain/mitigation"
)

// Repository query plans persistence.
type Repository interface {
	Create(ctx context.Context, m *mitigationDomain.Mitigation) error
	FindByHazardID(ctx context.Context, hazardID string) ([]*mitigationDomain.Mitigation, error)
}

// Service manages mitigation plans.
type Service struct {
	repo Repository
}

// NewService instantiates Mitigation Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateMitigation persists plan data.
func (s *Service) CreateMitigation(ctx context.Context, m *mitigationDomain.Mitigation) (*mitigationDomain.Mitigation, error) {
	if err := m.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}
