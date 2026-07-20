package execution

import (
	"context"

	evidenceDomain "prahari/services/audit/internal/domain/evidence"
)

// Repository persistence evidence.
type Repository interface {
	Create(ctx context.Context, e *evidenceDomain.Evidence) error
}

// Service collects walkthrough documents evidence.
type Service struct {
	repo Repository
}

// NewService instantiates Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// SaveEvidence records S3 reference pointers keys.
func (s *Service) SaveEvidence(ctx context.Context, e *evidenceDomain.Evidence) (*evidenceDomain.Evidence, error) {
	if err := e.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, e); err != nil {
		return nil, err
	}
	return e, nil
}
