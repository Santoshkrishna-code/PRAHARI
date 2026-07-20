package residual

import (
	"context"

	residualDomain "prahari/services/risk/internal/domain/residualrisk"
)

// Repository persistence residual scores.
type Repository interface {
	Create(ctx context.Context, rr *residualDomain.ResidualRisk) error
	FindByRiskID(ctx context.Context, riskID string) ([]*residualDomain.ResidualRisk, error)
}

// Service manages post-mitigation scores evaluations.
type Service struct {
	repo Repository
}

// NewService instantiates Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CalculateResidual persists calculations.
func (s *Service) CalculateResidual(ctx context.Context, rr *residualDomain.ResidualRisk) (*residualDomain.ResidualRisk, error) {
	if err := rr.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, rr); err != nil {
		return nil, err
	}
	return rr, nil
}
