package environment

import (
	"context"
	"time"

	"prahari/services/environmental/internal/domain/audittrail"
	"prahari/services/environmental/internal/domain/environment"
	"prahari/services/environmental/internal/domain/timeline"
)

type Repository interface {
	SaveAspect(ctx context.Context, a *environment.EnvironmentalAspect) error
	SaveImpact(ctx context.Context, i *environment.EnvironmentalImpact) error
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

func (s *Service) RegisterAspect(ctx context.Context, a *environment.EnvironmentalAspect) error {
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()

	if err := a.Validate(); err != nil {
		return err
	}

	if err := s.repo.SaveAspect(ctx, a); err != nil {
		return err
	}

	_ = s.trail.LogAuditTrail(ctx, audittrail.NewEntry("CREATE", "environmental_aspect", a.ID, "SYSTEM", nil, map[string]string{"name": a.Name}))
	_ = s.timeline.LogTimeline(ctx, timeline.NewEvent(a.ID, "ASPECT_REGISTERED", "SYSTEM", "Environmental Aspect registered successfully", nil))

	return nil
}

func (s *Service) EvaluateImpact(ctx context.Context, i *environment.EnvironmentalImpact) error {
	i.EvaluateRiskScore()

	if err := s.repo.SaveImpact(ctx, i); err != nil {
		return err
	}

	_ = s.trail.LogAuditTrail(ctx, audittrail.NewEntry("CREATE", "environmental_impact", i.ID, "SYSTEM", nil, map[string]string{"aspect_id": i.AspectID}))
	return nil
}
