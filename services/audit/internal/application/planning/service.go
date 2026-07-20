package planning

import (
	"context"

	planDomain "prahari/services/audit/internal/domain/auditplan"
)

// Repository persistence audit planning.
type Repository interface {
	Create(ctx context.Context, ap *planDomain.AuditPlan) error
}

// Service manages annual schedules planning.
type Service struct {
	repo Repository
}

// NewService instantiates Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// PlanAudit schedules start and end dates.
func (s *Service) PlanAudit(ctx context.Context, ap *planDomain.AuditPlan) (*planDomain.AuditPlan, error) {
	if err := ap.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, ap); err != nil {
		return nil, err
	}
	return ap, nil
}
