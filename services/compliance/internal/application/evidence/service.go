package evidence

import (
	"context"

	evidenceDomain "prahari/services/compliance/internal/domain/evidence"
)

// Repository persistence evidence logs.
type Repository interface {
	Create(ctx context.Context, e *evidenceDomain.Evidence) error
	FindByObligationID(ctx context.Context, obligationID string) ([]*evidenceDomain.Evidence, error)
}

// Service collects upload documents references.
type Service struct {
	repo Repository
}

// NewService instantiates Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// SaveEvidence logs upload files reference keys.
func (s *Service) SaveEvidence(ctx context.Context, e *evidenceDomain.Evidence) (*evidenceDomain.Evidence, error) {
	if err := e.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, e); err != nil {
		return nil, err
	}
	return e, nil
}
