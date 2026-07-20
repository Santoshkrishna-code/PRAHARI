package action

import (
	"context"
	"time"

	"github.com/google/uuid"

	capaDomain "prahari/services/inspection/internal/domain/capa"
	auditDomain "prahari/services/inspection/internal/domain/audit"
	timelineDomain "prahari/services/inspection/internal/domain/timeline"
)

// Repository defines action persistent ports.
type Repository interface {
	Create(ctx context.Context, c *capaDomain.CAPA) error
	FindByID(ctx context.Context, id string) (*capaDomain.CAPA, error)
	FindByInspectionID(ctx context.Context, inspectionID string) ([]*capaDomain.CAPA, error)
	Update(ctx context.Context, c *capaDomain.CAPA) error
}

// AuditLogger registers snaps.
type AuditLogger interface {
	Log(ctx context.Context, entry *auditDomain.Entry) error
}

// TimelineRecorder registers milestones.
type TimelineRecorder interface {
	Record(ctx context.Context, event *timelineDomain.Event) error
}

// Service manages CAPA workflows.
type Service struct {
	repo     Repository
	audit    AuditLogger
	timeline TimelineRecorder
}

// NewService instantiates CAPA Service.
func NewService(
	repo Repository,
	audit AuditLogger,
	timeline TimelineRecorder,
) *Service {
	return &Service{
		repo:     repo,
		audit:    audit,
		timeline: timeline,
	}
}

// CreateCAPA inserts action compliance items.
func (s *Service) CreateCAPA(ctx context.Context, c *capaDomain.CAPA) (*capaDomain.CAPA, error) {
	c.ID = uuid.New().String()
	c.CreatedAt = time.Now()
	c.Status = capaDomain.StatusOpen

	if err := c.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, c); err != nil {
		return nil, err
	}

	evt := timelineDomain.NewEvent(c.InspectionID, timelineDomain.EventCAPACreated, c.AssigneeID, "CAPA action generated from failed walkthrough findings", nil)
	_ = s.timeline.Record(ctx, evt)

	return c, nil
}

// CompleteCAPA marks completed task lists status.
func (s *Service) CompleteCAPA(ctx context.Context, id string) error {
	c, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	c.Complete()
	return s.repo.Update(ctx, c)
}

// VerifyCAPA registers manager signoffs.
func (s *Service) VerifyCAPA(ctx context.Context, id, verifier string) error {
	c, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	c.Verify(verifier)
	return s.repo.Update(ctx, c)
}
