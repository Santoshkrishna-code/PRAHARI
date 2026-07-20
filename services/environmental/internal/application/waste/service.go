package waste

import (
	"context"
	"time"

	"prahari/services/environmental/internal/domain/audittrail"
	"prahari/services/environmental/internal/domain/hazardouswaste"
	"prahari/services/environmental/internal/domain/timeline"
	"prahari/services/environmental/internal/domain/waste"
)

type Repository interface {
	SaveSolidWaste(ctx context.Context, w *waste.SolidWaste) error
	SaveHazardousWaste(ctx context.Context, h *hazardouswaste.HazardousWaste) error
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

func (s *Service) LogSolidWaste(ctx context.Context, w *waste.SolidWaste) error {
	w.CreatedAt = time.Now()

	if err := w.Validate(); err != nil {
		return err
	}

	if err := s.repo.SaveSolidWaste(ctx, w); err != nil {
		return err
	}

	_ = s.events.Publish(ctx, "waste.disposed", w)
	_ = s.trail.LogAuditTrail(ctx, audittrail.NewEntry("CREATE", "solid_waste", w.ID, "SYSTEM", nil, map[string]string{"category": w.WasteCategory}))
	_ = s.timeline.LogTimeline(ctx, timeline.NewEvent(w.ID, "SOLID_WASTE_DISPOSED", "SYSTEM", "Solid waste disposal logged successfully", nil))

	return nil
}

func (s *Service) LogHazardousWaste(ctx context.Context, h *hazardouswaste.HazardousWaste) error {
	h.CreatedAt = time.Now()

	if err := h.Validate(); err != nil {
		return err
	}

	if err := s.repo.SaveHazardousWaste(ctx, h); err != nil {
		return err
	}

	_ = s.events.Publish(ctx, "waste.disposed", h)
	_ = s.trail.LogAuditTrail(ctx, audittrail.NewEntry("CREATE", "hazardous_waste", h.ID, "SYSTEM", nil, map[string]string{"chemical": h.ChemicalName}))
	_ = s.timeline.LogTimeline(ctx, timeline.NewEvent(h.ID, "HAZARDOUS_WASTE_DISPOSED", "SYSTEM", "Hazardous waste disposal logged successfully", nil))

	return nil
}
