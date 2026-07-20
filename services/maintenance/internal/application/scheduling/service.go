package scheduling

import (
	"context"

	scheduleDomain "prahari/services/maintenance/internal/domain/schedule"
)

// Repository persistence schedules operations.
type Repository interface {
	Create(ctx context.Context, s *scheduleDomain.Schedule) error
	FindByID(ctx context.Context, id string) (*scheduleDomain.Schedule, error)
}

// Service manages maintenance scheduling.
type Service struct {
	repo Repository
}

// NewService instantiates Scheduling Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// ScheduleMaintenance persists calendar bookings.
func (s *Service) ScheduleMaintenance(ctx context.Context, sch *scheduleDomain.Schedule) (*scheduleDomain.Schedule, error) {
	if err := sch.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, sch); err != nil {
		return nil, err
	}
	return sch, nil
}
