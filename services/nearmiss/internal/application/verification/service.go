package verification

import (
	"context"

	verifyDomain "prahari/services/nearmiss/internal/domain/verification"
)

// Repository persistence gate.
type Repository interface {
	Create(ctx context.Context, v *verifyDomain.Verification) error
	FindByNearMissID(ctx context.Context, nearmissID string) ([]*verifyDomain.Verification, error)
}

// Service manages verification audits.
type Service struct {
	repo Repository
}

// NewService instantiates Verification Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// VerifyNearMiss persists audit entry log.
func (s *Service) VerifyNearMiss(ctx context.Context, v *verifyDomain.Verification) (*verifyDomain.Verification, error) {
	if err := v.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, v); err != nil {
		return nil, err
	}
	return v, nil
}
