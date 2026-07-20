package planning

import (
	"context"

	planDomain "prahari/services/maintenance/internal/domain/maintenanceplan"
)

// Repository query preventive recurrence.
type Repository interface {
	Create(ctx context.Context, p *planDomain.MaintenancePlan) error
	FindByID(ctx context.Context, id string) (*planDomain.MaintenancePlan, error)
}

// Service manages plans setups.
type Service struct {
	repo Repository
}

// NewService instantiates Planning Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreatePlan inserts recurrence template rules.
func (s *Service) CreatePlan(ctx context.Context, plan *planDomain.MaintenancePlan) (*planDomain.MaintenancePlan, error) {
	plan.CalculateNextRun()
	if err := plan.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, plan); err != nil {
		return nil, err
	}
	return plan, nil
}
