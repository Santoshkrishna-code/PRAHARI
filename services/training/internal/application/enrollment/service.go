package enrollment

import (
	"context"

	enrollDomain "prahari/services/training/internal/domain/enrollment"
)

// Repository persistence enrollments.
type Repository interface {
	Create(ctx context.Context, e *enrollDomain.Enrollment) error
}

// Service manages trainees enrollments.
type Service struct {
	repo Repository
}

// NewService instantiates Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// EnrollTrainee registers trainee.
func (s *Service) EnrollTrainee(ctx context.Context, e *enrollDomain.Enrollment) (*enrollDomain.Enrollment, error) {
	if err := e.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, e); err != nil {
		return nil, err
	}
	return e, nil
}
