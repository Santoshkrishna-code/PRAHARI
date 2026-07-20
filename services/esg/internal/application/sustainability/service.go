package sustainability

import (
	"context"
	"time"

	"prahari/services/esg/internal/domain/audittrail"
	"prahari/services/esg/internal/domain/esgobjective"
	"prahari/services/esg/internal/domain/timeline"
)

type Repository interface {
	SaveObjective(ctx context.Context, o *esgobjective.Objective) error
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

func (s *Service) CreateObjective(ctx context.Context, o *esgobjective.Objective) error {
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
	o.Status = "OBJECTIVE_DEFINED"

	if err := o.Validate(); err != nil {
		return err
	}

	if err := s.repo.SaveObjective(ctx, o); err != nil {
		return err
	}

	_ = s.trail.LogAuditTrail(ctx, audittrail.NewEntry("CREATE", "sustainability_objective", o.ID, "SYSTEM", nil, map[string]string{"title": o.Title}))
	_ = s.timeline.LogTimeline(ctx, timeline.NewEvent(o.ID, "OBJECTIVE_CREATED", "SYSTEM", "Sustainability Objective defined", nil))

	return nil
}
