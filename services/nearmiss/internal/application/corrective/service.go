package corrective

import (
	"context"

	correctiveDomain "prahari/services/nearmiss/internal/domain/correctiveaction"
)

// Repository query plans persistence.
type Repository interface {
	Create(ctx context.Context, ca *correctiveDomain.CorrectiveAction) error
	FindByNearMissID(ctx context.Context, nearmissID string) ([]*correctiveDomain.CorrectiveAction, error)
}

// Service manages CAPA actions.
type Service struct {
	repo Repository
}

// NewService instantiates Corrective Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateCorrectiveAction persists plan data.
func (s *Service) CreateCorrectiveAction(ctx context.Context, ca *correctiveDomain.CorrectiveAction) (*correctiveDomain.CorrectiveAction, error) {
	if err := ca.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, ca); err != nil {
		return nil, err
	}
	return ca, nil
}
