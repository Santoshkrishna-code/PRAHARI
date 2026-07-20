package permit

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	prahariLogger "prahari/shared/logger"

	permitDomain "prahari/services/permit/internal/domain/permit"
	auditDomain "prahari/services/permit/internal/domain/audit"
	statusDomain "prahari/services/permit/internal/domain/status"
	timelineDomain "prahari/services/permit/internal/domain/timeline"
)

// EventPublisher defines the port for sending Kafka events.
type EventPublisher interface {
	Publish(ctx context.Context, topic string, payload interface{}) error
}

// AuditLogger defines the port for logging change entries.
type AuditLogger interface {
	Log(ctx context.Context, entry *auditDomain.Entry) error
}

// TimelineRecorder defines the port for registering milestones.
type TimelineRecorder interface {
	Record(ctx context.Context, event *timelineDomain.Event) error
}

// WorkflowClient defines the port for orchestrating approval chains.
type WorkflowClient interface {
	StartWorkflow(ctx context.Context, permitID, workflowType string) error
}

// Service manages permit lifecycle modifications.
type Service struct {
	repo     permitDomain.Repository
	events   EventPublisher
	audit    AuditLogger
	timeline TimelineRecorder
	workflow WorkflowClient
}

// NewService instantiates a Service.
func NewService(
	repo permitDomain.Repository,
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

// generatePermitNumber constructs a unique human-readable code.
func generatePermitNumber() string {
	now := time.Now()
	return fmt.Sprintf("PTW-%d%02d%02d-%s", now.Year(), now.Month(), now.Day(), uuid.New().String()[:8])
}

// CreatePermit constructs and saves a draft Permit.
func (s *Service) CreatePermit(ctx context.Context, cmd CreatePermitCommand) (*permitDomain.Permit, error) {
	p := &permitDomain.Permit{
		ID:              uuid.New().String(),
		PermitNumber:    generatePermitNumber(),
		Title:           cmd.Title,
		Description:     cmd.Description,
		PermitTypeID:    cmd.PermitTypeID,
		StatusCode:      string(statusDomain.CodeDraft),
		RiskLevel:       permitDomain.RiskLevelLow,
		ApplicantID:     cmd.ApplicantID,
		SupervisorID:    cmd.SupervisorID,
		DepartmentID:    cmd.DepartmentID,
		ContractorID:    cmd.ContractorID,
		WorkAreaID:      cmd.WorkAreaID,
		WorkDescription: cmd.WorkDescription,
		PlannedStartAt:  cmd.PlannedStartAt,
		PlannedEndAt:    cmd.PlannedEndAt,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := p.Validate(); err != nil {
		return nil, fmt.Errorf("permit validation failed: %w", err)
	}

	if err := s.repo.Create(ctx, p); err != nil {
		return nil, fmt.Errorf("failed to save permit: %w", err)
	}

	// Logs
	evt := timelineDomain.NewEvent(p.ID, timelineDomain.EventCreated, cmd.ApplicantID, "Permit draft created", nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("permit", p.ID, auditDomain.ActionCreated, cmd.ApplicantID, nil, p)
	_ = s.audit.Log(ctx, auditLog)

	_ = s.events.Publish(ctx, "permit.created", p)

	return p, nil
}

// GetPermit retrieves a permit by identifier.
func (s *Service) GetPermit(ctx context.Context, id string) (*permitDomain.Permit, error) {
	return s.repo.FindByID(ctx, id)
}

// UpdatePermit modifies editable attributes.
func (s *Service) UpdatePermit(ctx context.Context, id string, cmd UpdatePermitCommand, actorID string) (*permitDomain.Permit, error) {
	p, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	oldSnapshot := *p

	if cmd.Title != "" {
		p.Title = cmd.Title
	}
	if cmd.Description != "" {
		p.Description = cmd.Description
	}
	if cmd.WorkDescription != "" {
		p.WorkDescription = cmd.WorkDescription
	}
	if !cmd.PlannedStartAt.IsZero() {
		p.PlannedStartAt = cmd.PlannedStartAt
	}
	if !cmd.PlannedEndAt.IsZero() {
		p.PlannedEndAt = cmd.PlannedEndAt
	}
	p.UpdatedAt = time.Now()

	if err := p.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Update(ctx, p); err != nil {
		return nil, err
	}

	auditLog := auditDomain.NewEntry("permit", p.ID, auditDomain.ActionUpdated, actorID, &oldSnapshot, p)
	_ = s.audit.Log(ctx, auditLog)

	_ = s.events.Publish(ctx, "permit.updated", p)

	return p, nil
}

// TransitionStatus transitions status through the state machine.
func (s *Service) TransitionStatus(ctx context.Context, cmd TransitionStatusCommand) error {
	p, err := s.repo.FindByID(ctx, cmd.PermitID)
	if err != nil {
		return err
	}

	fromStatus := statusDomain.Code(p.StatusCode)
	toStatus := statusDomain.Code(cmd.TargetCode)

	if err := statusDomain.ValidateTransition(fromStatus, toStatus); err != nil {
		return err
	}

	p.StatusCode = cmd.TargetCode
	p.UpdatedAt = time.Now()

	// Apply side effects
	now := time.Now()
	if toStatus == statusDomain.CodeActive {
		p.ActualStartAt = &now
		validTime := now.Add(8 * time.Hour) // Default 8-hour permit life
		p.ValidUntil = &validTime
	} else if toStatus == statusDomain.CodeCompleted {
		p.ActualEndAt = &now
	}

	if err := s.repo.Update(ctx, p); err != nil {
		return err
	}

	// Recording
	evt := timelineDomain.NewEvent(p.ID, timelineDomain.EventType(fmt.Sprintf("PERMIT_%s", cmd.TargetCode)), cmd.ActorID,
		fmt.Sprintf("Transitioned to %s: %s", cmd.TargetCode, cmd.Reason), nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("permit", p.ID, auditDomain.ActionStatusChanged, cmd.ActorID,
		map[string]string{"status": string(fromStatus)}, map[string]string{"status": cmd.TargetCode})
	_ = s.audit.Log(ctx, auditLog)

	eventTopic := fmt.Sprintf("permit.%s", cmd.TargetCode)
	_ = s.events.Publish(ctx, eventTopic, p)

	// Trigger workflow on submission
	if toStatus == statusDomain.CodeSubmitted {
		_ = s.workflow.StartWorkflow(ctx, p.ID, "permit_approval_flow")
	}

	return nil
}
