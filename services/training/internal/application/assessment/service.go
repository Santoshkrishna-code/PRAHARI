package assessment

import (
	"context"

	assessDomain "prahari/services/training/internal/domain/assessment"
)

// Repository persistence assessments.
type Repository interface {
	Create(ctx context.Context, a *assessDomain.Assessment) error
}

// Service manages score checks.
type Service struct {
	repo Repository
}

// NewService instantiates Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// RecordAssessment registers score metrics.
func (s *Service) RecordAssessment(ctx context.Context, a *assessDomain.Assessment) (*assessDomain.Assessment, error) {
	if err := a.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, a); err != nil {
		return nil, err
	}
	return a, nil
}
