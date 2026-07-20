package review

import (
	"context"

	reviewDomain "prahari/services/audit/internal/domain/review"
)

// Repository persistence reviews checks.
type Repository interface {
	Create(ctx context.Context, ar *reviewDomain.AuditReview) error
}

// Service manages checklist evaluations verification reviews.
type Service struct {
	repo Repository
}

// NewService instantiates Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// LogReview registers evaluation.
func (s *Service) LogReview(ctx context.Context, ar *reviewDomain.AuditReview) (*reviewDomain.AuditReview, error) {
	if err := ar.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, ar); err != nil {
		return nil, err
	}
	return ar, nil
}
