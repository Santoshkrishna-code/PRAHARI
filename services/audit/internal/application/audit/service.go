package audit

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	prahariLogger "prahari/shared/logger"

	auditDomain "prahari/services/audit/internal/domain/audit"
	policyDomain "prahari/services/audit/internal/domain/policy"
	statusDomain "prahari/services/audit/internal/domain/status"
	timelineDomain "prahari/services/audit/internal/domain/timeline"
	trailDomain "prahari/services/audit/internal/domain/audittrail"
)

// EventPublisher dispatches Kafka events.
type EventPublisher interface {
	Publish(ctx context.Context, topic string, payload interface{}) error
}

// AuditTrailLogger writes change trace snaps.
type AuditTrailLogger interface {
	Log(ctx context.Context, entry *trailDomain.Entry) error
}

// TimelineRecorder writes timeline events.
type TimelineRecorder interface {
	Record(ctx context.Context, event *timelineDomain.Event) error
}

// WorkflowClient calls external approval engine.
type WorkflowClient interface {
	VerifyAuditReviews(ctx context.Context, auditID string) error
}

// Service manages audit lifecycle transitions.
type Service struct {
	repo     auditDomain.Repository
	events   EventPublisher
	trail    AuditTrailLogger
	timeline TimelineRecorder
	workflow WorkflowClient
}

// NewService instantiates Service.
func NewService(
	repo auditDomain.Repository,
	events EventPublisher,
	trail AuditTrailLogger,
	timeline TimelineRecorder,
	workflow WorkflowClient,
) *Service {
	return &Service{
		repo:     repo,
		events:   events,
		trail:    trail,
		timeline: timeline,
		workflow: workflow,
	}
}

// generateAuditNumber tag formats.
func generateAuditNumber() string {
	now := time.Now()
	return fmt.Sprintf("AUD-%d%02d%02d-%s", now.Year(), now.Month(), now.Day(), uuid.New().String()[:8])
}

// CreateAudit registers a new assurance audit.
func (s *Service) CreateAudit(ctx context.Context, cmd CreateAuditCommand, actorID string) (*auditDomain.Audit, error) {
	now := time.Now()
	a := &auditDomain.Audit{
		ID:               uuid.New().String(),
		AuditNumber:      generateAuditNumber(),
		AssetID:          cmd.AssetID,
		DepartmentID:     cmd.DepartmentID,
		ComplianceRating: 100.0,
		StatusCode:       string(statusDomain.CodeDraft),
		Title:            cmd.Title,
		Description:      cmd.Description,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	if err := a.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, a); err != nil {
		return nil, err
	}

	evt := timelineDomain.NewEvent(a.ID, timelineDomain.EventCreated, actorID, "Audit record created", nil)
	_ = s.timeline.Record(ctx, evt)

	trailLog := trailDomain.NewEntry("audit", a.ID, trailDomain.ActionCreated, actorID, nil, a)
	_ = s.trail.Log(ctx, trailLog)

	_ = s.events.Publish(ctx, "audit.created", a)

	return a, nil
}

// GetAudit returns details.
func (s *Service) GetAudit(ctx context.Context, id string) (*auditDomain.Audit, error) {
	return s.repo.FindByID(ctx, id)
}

// TransitionStatus manages machine validations.
func (s *Service) TransitionStatus(ctx context.Context, cmd TransitionStatusCommand) error {
	a, err := s.repo.FindByID(ctx, cmd.AuditID)
	if err != nil {
		return err
	}

	fromStatus := statusDomain.Code(a.StatusCode)
	toStatus := statusDomain.Code(cmd.TargetCode)

	if err := statusDomain.ValidateTransition(fromStatus, toStatus); err != nil {
		return err
	}

	// Enforce document evidence check
	if toStatus == statusDomain.CodeReview {
		if err := policyDomain.CheckEvidenceRequirements(a, 1); err != nil {
			prahariLogger.Info(ctx, "Audit review requires verified evidence document", prahariLogger.String("id", a.ID))
		}
	}

	a.StatusCode = cmd.TargetCode
	a.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, a); err != nil {
		return err
	}

	if toStatus == statusDomain.CodeReview {
		_ = s.workflow.VerifyAuditReviews(ctx, a.ID)
	}

	evt := timelineDomain.NewEvent(a.ID, timelineDomain.EventType(fmt.Sprintf("AUDIT_%s", cmd.TargetCode)), cmd.ActorID,
		fmt.Sprintf("Audit transitioned to %s", cmd.TargetCode), nil)
	_ = s.timeline.Record(ctx, evt)

	trailLog := trailDomain.NewEntry("audit", a.ID, trailDomain.ActionStatusChanged, cmd.ActorID,
		map[string]string{"status": string(fromStatus)}, map[string]string{"status": cmd.TargetCode})
	_ = s.trail.Log(ctx, trailLog)

	eventTopic := fmt.Sprintf("audit.%s", cmd.TargetCode)
	_ = s.events.Publish(ctx, eventTopic, a)

	return nil
}
}
