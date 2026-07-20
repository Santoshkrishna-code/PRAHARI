package inspection

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	prahariLogger "prahari/shared/logger"

	inspectionDomain "prahari/services/inspection/internal/domain/inspection"
	auditDomain "prahari/services/inspection/internal/domain/audit"
	statusDomain "prahari/services/inspection/internal/domain/status"
	timelineDomain "prahari/services/inspection/internal/domain/timeline"
)

// EventPublisher defines the Kafka message dispatches port.
type EventPublisher interface {
	Publish(ctx context.Context, topic string, payload interface{}) error
}

// AuditLogger writes immutable snap transaction records.
type AuditLogger interface {
	Log(ctx context.Context, entry *auditDomain.Entry) error
}

// TimelineRecorder registers milestones events logs.
type TimelineRecorder interface {
	Record(ctx context.Context, event *timelineDomain.Event) error
}

// WorkflowClient triggers manager review workflows.
type WorkflowClient interface {
	StartWorkflow(ctx context.Context, inspectionID, workflowType string) error
}

// IncidentClient raises incident tickets on critical/catastrophic NCR failure.
type IncidentClient interface {
	CreateIncidentFromFinding(ctx context.Context, title, description, findingID string) (string, error)
}

// Service manages inspection lifecycle transitions.
type Service struct {
	repo     inspectionDomain.Repository
	events   EventPublisher
	audit    AuditLogger
	timeline TimelineRecorder
	workflow WorkflowClient
	incident IncidentClient
}

// NewService instantiates Service.
func NewService(
	repo inspectionDomain.Repository,
	events EventPublisher,
	audit AuditLogger,
	timeline TimelineRecorder,
	workflow WorkflowClient,
	incident IncidentClient,
) *Service {
	return &Service{
		repo:     repo,
		events:   events,
		audit:    audit,
		timeline: timeline,
		workflow: workflow,
		incident: incident,
	}
}

// generateInspectionNumber constructs unique reference string codes.
func generateInspectionNumber() string {
	now := time.Now()
	return fmt.Sprintf("INSP-%d%02d%02d-%s", now.Year(), now.Month(), now.Day(), uuid.New().String()[:8])
}

// CreateInspection inserts a safety audit inspection record.
func (s *Service) CreateInspection(ctx context.Context, cmd CreateInspectionCommand) (*inspectionDomain.Inspection, error) {
	now := time.Now()
	i := &inspectionDomain.Inspection{
		ID:               uuid.New().String(),
		InspectionNumber: generateInspectionNumber(),
		Title:            cmd.Title,
		Description:      cmd.Description,
		InspectionType:   inspectionDomain.Type(cmd.InspectionType),
		StatusCode:       string(statusDomain.CodeDraft),
		InspectorID:      cmd.InspectorID,
		DepartmentID:     cmd.DepartmentID,
		AssetID:          cmd.AssetID,
		LinkedPermitID:   cmd.LinkedPermitID,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	if err := i.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, i); err != nil {
		return nil, err
	}

	evt := timelineDomain.NewEvent(i.ID, timelineDomain.EventCreated, cmd.InspectorID, "Inspection walkthrough draft created", nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("inspection", i.ID, auditDomain.ActionCreated, cmd.InspectorID, nil, i)
	_ = s.audit.Log(ctx, auditLog)

	_ = s.events.Publish(ctx, "inspection.created", i)

	return i, nil
}

// GetInspection returns inspection aggregate.
func (s *Service) GetInspection(ctx context.Context, id string) (*inspectionDomain.Inspection, error) {
	return s.repo.FindByID(ctx, id)
}

// UpdateInspection saves edits.
func (s *Service) UpdateInspection(ctx context.Context, id string, cmd UpdateInspectionCommand, actorID string) (*inspectionDomain.Inspection, error) {
	i, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	oldSnapshot := *i

	if cmd.Title != "" {
		i.Title = cmd.Title
	}
	if cmd.Description != "" {
		i.Description = cmd.Description
	}
	i.UpdatedAt = time.Now()

	if err := i.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Update(ctx, i); err != nil {
		return nil, err
	}

	auditLog := auditDomain.NewEntry("inspection", i.ID, auditDomain.ActionUpdated, actorID, &oldSnapshot, i)
	_ = s.audit.Log(ctx, auditLog)

	_ = s.events.Publish(ctx, "inspection.updated", i)

	return i, nil
}

// TransitionStatus handles machine matrix validations.
func (s *Service) TransitionStatus(ctx context.Context, cmd TransitionStatusCommand) error {
	i, err := s.repo.FindByID(ctx, cmd.InspectionID)
	if err != nil {
		return err
	}

	fromStatus := statusDomain.Code(i.StatusCode)
	toStatus := statusDomain.Code(cmd.TargetCode)

	if err := statusDomain.ValidateTransition(fromStatus, toStatus); err != nil {
		return err
	}

	i.StatusCode = cmd.TargetCode
	i.UpdatedAt = time.Now()

	// Capture execution date markers
	now := time.Now()
	if toStatus == statusDomain.CodeInProgress {
		i.StartedAt = &now
	} else if toStatus == statusDomain.CodeCompleted {
		i.CompletedAt = &now
	}

	if err := s.repo.Update(ctx, i); err != nil {
		return err
	}

	evt := timelineDomain.NewEvent(i.ID, timelineDomain.EventType(fmt.Sprintf("INSPECTION_%s", cmd.TargetCode)), cmd.ActorID,
		fmt.Sprintf("Status transitioned to %s", cmd.TargetCode), nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("inspection", i.ID, auditDomain.ActionStatusChanged, cmd.ActorID,
		map[string]string{"status": string(fromStatus)}, map[string]string{"status": cmd.TargetCode})
	_ = s.audit.Log(ctx, auditLog)

	eventTopic := fmt.Sprintf("inspection.%s", cmd.TargetCode)
	_ = s.events.Publish(ctx, eventTopic, i)

	// Trigger manager review workflow on completion
	if toStatus == statusDomain.CodeCompleted {
		_ = s.workflow.StartWorkflow(ctx, i.ID, "inspection_review_flow")
	}

	return nil
}

// AutoEscalateCriticalFinding triggers gRPC calls to Incident Service.
func (s *Service) AutoEscalateCriticalFinding(ctx context.Context, inspectionID, findingID, description string) (string, error) {
	prahariLogger.Info(ctx, "Executing auto incident creation workflow for critical finding",
		prahariLogger.String("inspection_id", inspectionID),
		prahariLogger.String("finding_id", findingID))

	title := fmt.Sprintf("Escalated Finding: %s", description)
	incidentID, err := s.incident.CreateIncidentFromFinding(ctx, title, "Automated ticket raised from critical safety audit finding.", findingID)
	if err != nil {
		return "", err
	}

	return incidentID, nil
}
