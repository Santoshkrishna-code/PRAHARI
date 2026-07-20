package followup

import (
	"context"

	followupDomain "prahari/services/observation/internal/domain/followup"
)

// Repository persistence gate.
type Repository interface {
	Create(ctx context.Context, f *followupDomain.FollowUp) error
	FindByObservationID(ctx context.Context, observationID string) ([]*followupDomain.FollowUp, error)
}

// Service manages follow-up reviews.
type Service struct {
	repo Repository
}

// NewService instantiates FollowUp Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// LogFollowUp persists audit entry.
func (s *Service) LogFollowUp(ctx context.Context, f *followupDomain.FollowUp) (*followupDomain.FollowUp, error) {
	if err := f.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, f); err != nil {
		return nil, err
	}
	return f, nil
}
