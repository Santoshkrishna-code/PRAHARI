package schedule

import (
	"context"
	"time"

	"github.com/google/uuid"

	scheduleDomain "prahari/services/inspection/internal/domain/schedule"
)

// Repository defines schedule query ports.
type Repository interface {
	Create(ctx context.Context, s *scheduleDomain.Schedule) error
	FindByID(ctx context.Context, id string) (*scheduleDomain.Schedule, error)
	ListActive(ctx context.Context) ([]*scheduleDomain.Schedule, error)
	Update(ctx context.Context, s *scheduleDomain.Schedule) error
}

// Service schedules upcoming safety walkthroughs automatically.
type Service struct {
	repo Repository
}

// NewService instantiates Schedule Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateSchedule inserts recurrence blueprints.
func (s *Service) CreateSchedule(ctx context.Context, sch *scheduleDomain.Schedule) (*scheduleDomain.Schedule, error) {
	sch.ID = uuid.New().String()
	sch.IsActive = true
	sch.LastExecutionDate = time.Now()
	sch.CalculateNextDate()

	if err := sch.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, sch); err != nil {
		return nil, err
	}

	return sch, nil
}

// TriggerExecution advances schedule dates.
func (s *Service) TriggerExecution(ctx context.Context, id string) error {
	sch, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	sch.CalculateNextDate()
	return s.repo.Update(ctx, sch)
}
