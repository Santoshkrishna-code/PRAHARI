package optimization

import (
	"context"
	"time"

	"prahari/services/energy/internal/domain/audittrail"
	"prahari/services/energy/internal/domain/optimization"
	"prahari/services/energy/internal/domain/timeline"
)

type Repository interface {
	SaveRecommendation(ctx context.Context, r *optimization.Recommendation) error
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

func (s *Service) RecommendOptimization(ctx context.Context, r *optimization.Recommendation) error {
	r.CreatedAt = time.Now()
	r.Status = "RECOMMENDED"

	if err := r.Validate(); err != nil {
		return err
	}

	if err := s.repo.SaveRecommendation(ctx, r); err != nil {
		return err
	}

	_ = s.events.Publish(ctx, "optimization.recommended", r)
	_ = s.trail.LogAuditTrail(ctx, audittrail.NewEntry("CREATE", "optimization_recommendation", r.ID, "SYSTEM", nil, map[string]string{"title": r.Title}))
	_ = s.timeline.LogTimeline(ctx, timeline.NewEvent(r.ID, "OPTIMIZATION_RECOMMENDED", "SYSTEM", "Energy conservation measure recommended", nil))

	return nil
}
