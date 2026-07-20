package review

import (
	"context"

	reviewDomain "prahari/services/risk/internal/domain/review"
)

// Repository persistence review sessions.
type Repository interface {
	Create(ctx context.Context, rr *reviewDomain.RiskReview) error
	FindByRiskID(ctx context.Context, riskID string) ([]*reviewDomain.RiskReview, error)
}

// Service manages reassessment reviews.
type Service struct {
	repo Repository
}

// NewService instantiates Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateRiskReview logs validation checks.
func (s *Service) CreateRiskReview(ctx context.Context, rr *reviewDomain.RiskReview) (*reviewDomain.RiskReview, error) {
	if err := rr.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, rr); err != nil {
		return nil, err
	}
	return rr, nil
}
