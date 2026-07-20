package consumption

import (
	"context"
	"time"

	"prahari/services/energy/internal/domain/audittrail"
	"prahari/services/energy/internal/domain/meterreading"
	"prahari/services/energy/internal/domain/timeline"
)

type Repository interface {
	SaveReading(ctx context.Context, r *meterreading.Reading) error
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

func (s *Service) RecordMeterReading(ctx context.Context, r *meterreading.Reading) error {
	r.CreatedAt = time.Now()
	if r.ReadingTime.IsZero() {
		r.ReadingTime = time.Now()
	}

	if err := r.Validate(); err != nil {
		return err
	}

	if err := s.repo.SaveReading(ctx, r); err != nil {
		return err
	}

	_ = s.events.Publish(ctx, "meter.reading.recorded", r)
	_ = s.trail.LogAuditTrail(ctx, audittrail.NewEntry("RECORD", "meter_reading", r.ID, "SYSTEM", nil, map[string]string{"meter_id": r.MeterID}))
	_ = s.timeline.LogTimeline(ctx, timeline.NewEvent(r.ID, "METER_READING_RECORDED", "SYSTEM", "Smart meter reading telemetry logged", nil))

	return nil
}
