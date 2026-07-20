package coaching

import (
	"context"

	coachingDomain "prahari/services/observation/internal/domain/coaching"
)

// Repository persistence coaching sessions definitions.
type Repository interface {
	Create(ctx context.Context, cs *coachingDomain.CoachingSession) error
	FindByObservationID(ctx context.Context, observationID string) ([]*coachingDomain.CoachingSession, error)
}

// Service manages BBS coaching dialogues.
type Service struct {
	repo Repository
}

// NewService instantiates Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateCoachingSession persists session data.
func (s *Service) CreateCoachingSession(ctx context.Context, cs *coachingDomain.CoachingSession) (*coachingDomain.CoachingSession, error) {
	if err := cs.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, cs); err != nil {
		return nil, err
	}
	return cs, nil
}
