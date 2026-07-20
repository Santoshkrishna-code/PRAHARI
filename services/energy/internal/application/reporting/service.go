package reporting

import (
	"context"
	"time"

	"prahari/services/energy/internal/domain/audittrail"
	"prahari/services/energy/internal/domain/energyprofile"
	"prahari/services/energy/internal/domain/timeline"
)

type Repository interface {
	SaveProfile(ctx context.Context, p *energyprofile.Profile) error
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

func (s *Service) CreateProfile(ctx context.Context, p *energyprofile.Profile) error {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	if err := p.Validate(); err != nil {
		return err
	}

	if err := s.repo.SaveProfile(ctx, p); err != nil {
		return err
	}

	_ = s.events.Publish(ctx, "energy.profile.created", p)
	_ = s.trail.LogAuditTrail(ctx, audittrail.NewEntry("CREATE", "energy_profile", p.ID, "SYSTEM", nil, map[string]string{"facility": p.FacilityName}))
	_ = s.timeline.LogTimeline(ctx, timeline.NewEvent(p.ID, "PROFILE_CREATED", "SYSTEM", "Corporate energy facility profile defined", nil))

	return nil
}
