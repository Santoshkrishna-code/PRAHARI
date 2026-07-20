package emission

import (
	"context"
	"time"

	"prahari/services/environmental/internal/domain/audittrail"
	"prahari/services/environmental/internal/domain/emission"
	"prahari/services/environmental/internal/domain/policy"
	"prahari/services/environmental/internal/domain/timeline"
)

type Repository interface {
	SaveEmission(ctx context.Context, e *emission.Emission) error
}

type EventPublisher interface {
	Publish(ctx context.Context, topic string, payload interface{}) error
}

type AuditTrailLogger interface {
	LogAuditTrail(ctx context.Context, entry *audittrail.Entry) error
}

type TimelineLogger interface {
	LogTimeline(ctx context.Context, e *timeline.Event) error
}

type Service struct {
	repo     Repository
	events   EventPublisher
	trail    AuditTrailLogger
	timeline TimelineLogger
}

func NewService(repo Repository, events EventPublisher, trail AuditTrailLogger, timeline TimelineLogger) *Service {
	return &Service{
		repo:     repo,
		events:   events,
		trail:    trail,
		timeline: timeline,
	}
}

func (s *Service) LogEmission(ctx context.Context, e *emission.Emission) error {
	e.CreatedAt = time.Now()

	if err := e.Validate(); err != nil {
		return err
	}

	// Policy evaluation
	if err := policy.EvaluateEmissionLimits(e); err != nil {
		e.IsExceeded = true
		_ = s.events.Publish(ctx, "emission.exceeded", e)
		_ = s.events.Publish(ctx, "compliance.failed", map[string]interface{}{
			"id":     e.ID,
			"source": "EMISSION_POINT_SOURCE",
			"gas":    e.GasType,
			"reason": err.Error(),
		})
	}

	if err := s.repo.SaveEmission(ctx, e); err != nil {
		return err
	}

	_ = s.trail.LogAuditTrail(ctx, audittrail.NewEntry("CREATE", "emission", e.ID, "SYSTEM", nil, map[string]string{"gas": e.GasType}))
	return nil
}
