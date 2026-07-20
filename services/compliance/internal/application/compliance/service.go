package compliance

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	prahariLogger "prahari/shared/logger"

	auditDomain "prahari/services/compliance/internal/domain/audit"
	complianceDomain "prahari/services/compliance/internal/domain/compliance"
	policyDomain "prahari/services/compliance/internal/domain/policy"
	statusDomain "prahari/services/compliance/internal/domain/status"
	timelineDomain "prahari/services/compliance/internal/domain/timeline"
)

// EventPublisher dispatches Kafka events.
type EventPublisher interface {
	Publish(ctx context.Context, topic string, payload interface{}) error
}

// AuditLogger writes change trace snaps.
type AuditLogger interface {
	Log(ctx context.Context, entry *auditDomain.Entry) error
}

// TimelineRecorder writes timeline events.
type TimelineRecorder interface {
	Record(ctx context.Context, event *timelineDomain.Event) error
}

// WorkflowClient calls external approval engine.
type WorkflowClient interface {
	VerifyComplianceReviews(ctx context.Context, complianceID string) error
}

// Service manages compliance register lifecycle transitions.
type Service struct {
	repo     complianceDomain.Repository
	events   EventPublisher
	audit    AuditLogger
	timeline TimelineRecorder
	workflow WorkflowClient
}

// NewService instantiates Service.
func NewService(
	repo complianceDomain.Repository,
	events EventPublisher,
	audit AuditLogger,
	timeline TimelineRecorder,
	workflow WorkflowClient,
) *Service {
	return &Service{
		repo:     repo,
		events:   events,
		audit:    audit,
		timeline: timeline,
		workflow: workflow,
	}
}

// generateComplianceNumber tag formats.
func generateComplianceNumber() string {
	now := time.Now()
	return fmt.Sprintf("CMP-%d%02d%02d-%s", now.Year(), now.Month(), now.Day(), uuid.New().String()[:8])
}

// CreateCompliance registers governance controls.
func (s *Service) CreateCompliance(ctx context.Context, cmd CreateComplianceCommand, actorID string) (*complianceDomain.Compliance, error) {
	now := time.Now()
	c := &complianceDomain.Compliance{
		ID:               uuid.New().String(),
		ComplianceNumber: generateComplianceNumber(),
		AssetID:          cmd.AssetID,
		DepartmentID:     cmd.DepartmentID,
		ComplianceScore:  100.0,
		StatusCode:       string(statusDomain.CodeDraft),
		Title:            cmd.Title,
		Description:      cmd.Description,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	if err := c.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, c); err != nil {
		return nil, err
	}

	evt := timelineDomain.NewEvent(c.ID, timelineDomain.EventCreated, actorID, "Governance compliance record created", nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("compliance", c.ID, auditDomain.ActionCreated, actorID, nil, c)
	_ = s.audit.Log(ctx, auditLog)

	_ = s.events.Publish(ctx, "compliance.created", c)

	return c, nil
}

// GetCompliance returns register details.
func (s *Service) GetCompliance(ctx context.Context, id string) (*complianceDomain.Compliance, error) {
	return s.repo.FindByID(ctx, id)
}

// TransitionStatus manages machine validations.
func (s *Service) TransitionStatus(ctx context.Context, cmd TransitionStatusCommand) error {
	c, err := s.repo.FindByID(ctx, cmd.ComplianceID)
	if err != nil {
		return err
	}

	fromStatus := statusDomain.Code(c.StatusCode)
	toStatus := statusDomain.Code(cmd.TargetCode)

	if err := statusDomain.ValidateTransition(fromStatus, toStatus); err != nil {
		return err
	}

	// Enforce document files checks
	if toStatus == statusDomain.CodeReview {
		if err := policyDomain.CheckEvidenceRequirements(c, 1); err != nil {
			prahariLogger.Info(ctx, "Obligation requires verified evidence document", prahariLogger.String("id", c.ID))
		}
	}

	c.StatusCode = cmd.TargetCode
	c.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, c); err != nil {
		return err
	}

	if toStatus == statusDomain.CodeReview {
		_ = s.workflow.VerifyComplianceReviews(ctx, c.ID)
	}

	evt := timelineDomain.NewEvent(c.ID, timelineDomain.EventType(fmt.Sprintf("COMPLIANCE_%s", cmd.TargetCode)), cmd.ActorID,
		fmt.Sprintf("Compliance transitioned to %s", cmd.TargetCode), nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("compliance", c.ID, auditDomain.ActionStatusChanged, cmd.ActorID,
		map[string]string{"status": string(fromStatus)}, map[string]string{"status": cmd.TargetCode})
	_ = s.audit.Log(ctx, auditLog)

	eventTopic := fmt.Sprintf("compliance.%s", cmd.TargetCode)
	_ = s.events.Publish(ctx, eventTopic, c)

	return nil
}
