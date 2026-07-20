package review

import (
	"context"

	reviewDomain "prahari/services/compliance/internal/domain/review"
)

// Repository persistence reviews checks.
type Repository interface {
	Create(ctx context.Context, cr *reviewDomain.ComplianceReview) error
	FindByComplianceID(ctx context.Context, complianceID string) ([]*reviewDomain.ComplianceReview, error)
}

// Service manages verification checklists reviews.
type Service struct {
	repo Repository
}

// NewService instantiates Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// LogReview registers evaluation.
func (s *Service) LogReview(ctx context.Context, cr *reviewDomain.ComplianceReview) (*reviewDomain.ComplianceReview, error) {
	if err := cr.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, cr); err != nil {
		return nil, err
	}
	return cr, nil
}
