package carbon

import (
	"context"
	"time"

	"prahari/services/esg/internal/domain/audittrail"
	"prahari/services/esg/internal/domain/carboninventory"
	"prahari/services/esg/internal/domain/timeline"
)

type Repository interface {
	SaveInventory(ctx context.Context, i *carboninventory.Inventory) error
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

func (s *Service) CalculateCarbon(ctx context.Context, inv *carboninventory.Inventory) error {
	inv.CreatedAt = time.Now()
	inv.CalculateTotal()

	if err := inv.Validate(); err != nil {
		return err
	}

	if err := s.repo.SaveInventory(ctx, inv); err != nil {
		return err
	}

	_ = s.events.Publish(ctx, "carbon.inventory.calculated", inv)
	_ = s.trail.LogAuditTrail(ctx, audittrail.NewEntry("CALCULATE", "carbon_inventory", inv.ID, "SYSTEM", nil, map[string]string{"total": "calculated"}))
	_ = s.timeline.LogTimeline(ctx, timeline.NewEvent(inv.ID, "CARBON_CALCULATED", "SYSTEM", "Scope 1/2/3 Carbon Inventory Calculated", nil))

	return nil
}
