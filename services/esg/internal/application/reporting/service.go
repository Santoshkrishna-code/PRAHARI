package reporting

import (
	"context"
	"time"

	"prahari/services/esg/internal/domain/audittrail"
	"prahari/services/esg/internal/domain/sustainabilityreport"
	"prahari/services/esg/internal/domain/timeline"
)

type Repository interface {
	SaveReport(ctx context.Context, r *sustainabilityreport.Report) error
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

func (s *Service) GenerateReport(ctx context.Context, r *sustainabilityreport.Report) error {
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()

	if err := r.Validate(); err != nil {
		return err
	}

	if err := s.repo.SaveReport(ctx, r); err != nil {
		return err
	}

	_ = s.events.Publish(ctx, "sustainability.report.generated", r)
	_ = s.trail.LogAuditTrail(ctx, audittrail.NewEntry("CREATE", "sustainability_report", r.ID, "SYSTEM", nil, map[string]string{"title": r.Title}))
	_ = s.timeline.LogTimeline(ctx, timeline.NewEvent(r.ID, "REPORT_GENERATED", "SYSTEM", "Corporate Sustainability Report compiled", nil))

	return nil
}
