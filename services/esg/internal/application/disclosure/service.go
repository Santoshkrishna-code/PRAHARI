package disclosure

import (
	"context"
	"time"

	"prahari/services/esg/internal/domain/audittrail"
	"prahari/services/esg/internal/domain/disclosure"
	"prahari/services/esg/internal/domain/timeline"
)

type Repository interface {
	SaveDisclosure(ctx context.Context, d *disclosure.Disclosure) error
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

func (s *Service) PublishDisclosure(ctx context.Context, d *disclosure.Disclosure) error {
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
	d.Status = "PUBLISHED"

	if err := d.Validate(); err != nil {
		return err
	}

	if err := s.repo.SaveDisclosure(ctx, d); err != nil {
		return err
	}

	_ = s.events.Publish(ctx, "disclosure.published", d)
	_ = s.trail.LogAuditTrail(ctx, audittrail.NewEntry("PUBLISH", "disclosure", d.ID, "SYSTEM", nil, map[string]string{"ref": d.ReferenceCode}))
	_ = s.timeline.LogTimeline(ctx, timeline.NewEvent(d.ID, "DISCLOSURE_PUBLISHED", "SYSTEM", "Regulatory framework disclosure published", nil))

	return nil
}
