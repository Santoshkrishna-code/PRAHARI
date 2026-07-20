package permit

import (
	"context"
	"time"

	"prahari/services/environmental/internal/domain/audittrail"
	"prahari/services/environmental/internal/domain/environmentalpermit"
	"prahari/services/environmental/internal/domain/timeline"
)

type Repository interface {
	SavePermit(ctx context.Context, p *environmentalpermit.EnvironmentalPermit) error
}

type AuditTrailLogger interface {
	LogAuditTrail(ctx context.Context, entry *audittrail.Entry) error
}

type TimelineLogger interface {
	LogTimeline(ctx context.Context, e *timeline.Event) error
}

type Service struct {
	repo     Repository
	trail    AuditTrailLogger
	timeline TimelineLogger
}

func NewService(repo Repository, trail AuditTrailLogger, timeline TimelineLogger) *Service {
	return &Service{
		repo:     repo,
		trail:    trail,
		timeline: timeline,
	}
}

func (s *Service) RegisterPermit(ctx context.Context, p *environmentalpermit.EnvironmentalPermit) error {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	if err := p.Validate(); err != nil {
		return err
	}

	if err := s.repo.SavePermit(ctx, p); err != nil {
		return err
	}

	_ = s.trail.LogAuditTrail(ctx, audittrail.NewEntry("CREATE", "environmental_permit", p.ID, "SYSTEM", nil, map[string]string{"number": p.PermitNumber}))
	_ = s.timeline.LogTimeline(ctx, timeline.NewEvent(p.ID, "PERMIT_REGISTERED", "SYSTEM", "Regulatory Environmental Permit registered successfully", nil))

	return nil
}
