package effectiveness

import (
	"context"

	effectDomain "prahari/services/observation/internal/domain/effectiveness"
)

// Repository persistence gate.
type Repository interface {
	Create(ctx context.Context, e *effectDomain.Effectiveness) error
	FindByObservationID(ctx context.Context, observationID string) ([]*effectDomain.Effectiveness, error)
}

// Service evaluates safety behavior improvements.
type Service struct {
	repo Repository
}

// NewService instantiates Effectiveness Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// EvaluateCoaching persists rating logs.
func (s *Service) EvaluateCoaching(ctx context.Context, e *effectDomain.Effectiveness) (*effectDomain.Effectiveness, error) {
	if err := e.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, e); err != nil {
		return nil, err
	}
	return e, nil
}
