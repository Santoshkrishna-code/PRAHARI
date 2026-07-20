package obligation

import (
	"context"

	obDomain "prahari/services/compliance/internal/domain/obligation"
)

// Repository persistence obligations checklist items.
type Repository interface {
	Create(ctx context.Context, o *obDomain.Obligation) error
	FindByComplianceID(ctx context.Context, complianceID string) ([]*obDomain.Obligation, error)
}

// Service manages statutory checklist obligations.
type Service struct {
	repo Repository
}

// NewService instantiates Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateObligation registers statutory items.
func (s *Service) CreateObligation(ctx context.Context, o *obDomain.Obligation) (*obDomain.Obligation, error) {
	if err := o.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, o); err != nil {
		return nil, err
	}
	return o, nil
}
