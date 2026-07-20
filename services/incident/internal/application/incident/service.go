package incident

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	prahariLogger "prahari/shared/logger"

	incidentDomain "prahari/services/incident/internal/domain/incident"
	auditDomain "prahari/services/incident/internal/domain/audit"
	statusDomain "prahari/services/incident/internal/domain/status"
	timelineDomain "prahari/services/incident/internal/domain/timeline"
)

// Repository defines the persistence port for the incident service.
type Repository interface {
	Create(ctx context.Context, incident *incidentDomain.Incident) error
	FindByID(ctx context.Context, id string) (*incidentDomain.Incident, error)
	Update(ctx context.Context, incident *incidentDomain.Incident) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, offset, limit int) ([]*incidentDomain.Incident, int, error)
}

// EventPublisher defines the port for publishing domain events.
type EventPublisher interface {
	Publish(ctx context.Context, topic string, payload interface{}) error
}

// AuditLogger defines the port for persisting audit entries.
type AuditLogger interface {
	Log(ctx context.Context, entry *auditDomain.Entry) error
}

// TimelineRecorder defines the port for recording timeline events.
type TimelineRecorder interface {
	Record(ctx context.Context, event *timelineDomain.Event) error
}

// WorkflowClient defines the port for triggering workflows.
type WorkflowClient interface {
	StartWorkflow(ctx context.Context, incidentID, workflowType string) error
}

// Service orchestrates incident lifecycle operations, coordinating between
// the domain layer, infrastructure adapters, and platform service integrations.
type Service struct {
	repo      Repository
	events    EventPublisher
	audit     AuditLogger
	timeline  TimelineRecorder
	workflow  WorkflowClient
}

// NewService constructs a Service with all required dependencies injected.
func NewService(
	repo Repository,
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

// generateIncidentNumber creates a unique human-readable incident number.
func generateIncidentNumber() string {
	now := time.Now()
	return fmt.Sprintf("INC-%d%02d%02d-%s", now.Year(), now.Month(), now.Day(), uuid.New().String()[:8])
}

// CreateIncident registers a new incident in the system. It validates domain invariants,
// persists the aggregate, records a timeline event, publishes a Kafka event, and
// triggers the incident workflow in the Workflow Engine.
func (s *Service) CreateIncident(ctx context.Context, cmd CreateIncidentCommand) (*incidentDomain.Incident, error) {
	inc := &incidentDomain.Incident{
		ID:             uuid.New().String(),
		IncidentNumber: generateIncidentNumber(),
		Title:          cmd.Title,
		Description:    cmd.Description,
		Type:           incidentDomain.Type(cmd.Type),
		CategoryID:     cmd.CategoryID,
		SeverityLevel:  cmd.SeverityLevel,
		PriorityLevel:  cmd.PriorityLevel,
		StatusCode:     string(statusDomain.CodeDraft),
		ReporterID:     cmd.ReporterID,
		DepartmentID:   cmd.DepartmentID,
		LocationID:     cmd.LocationID,
		LocationDetail: cmd.LocationDetail,
		OccurredAt:     cmd.OccurredAt,
		ReportedAt:     time.Now(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := inc.Validate(); err != nil {
		return nil, fmt.Errorf("incident validation failed: %w", err)
	}

	if err := s.repo.Create(ctx, inc); err != nil {
		return nil, fmt.Errorf("failed to persist incident: %w", err)
	}

	// Record timeline event
	evt := timelineDomain.NewEvent(inc.ID, timelineDomain.EventCreated, cmd.ReporterID,
		fmt.Sprintf("Incident %s created: %s", inc.IncidentNumber, inc.Title), nil)
	if err := s.timeline.Record(ctx, evt); err != nil {
		prahariLogger.Error(ctx, "Failed to record timeline event", prahariLogger.Err(err))
	}

	// Audit log
	entry := auditDomain.NewEntry("incident", inc.ID, auditDomain.ActionCreated, cmd.ReporterID, nil, inc)
	if err := s.audit.Log(ctx, entry); err != nil {
		prahariLogger.Error(ctx, "Failed to write audit log", prahariLogger.Err(err))
	}

	// Publish domain event
	if err := s.events.Publish(ctx, "incident.created", inc); err != nil {
		prahariLogger.Error(ctx, "Failed to publish incident.created event", prahariLogger.Err(err))
	}

	// Trigger workflow
	if err := s.workflow.StartWorkflow(ctx, inc.ID, "incident_review"); err != nil {
		prahariLogger.Error(ctx, "Failed to trigger incident workflow", prahariLogger.Err(err))
	}

	prahariLogger.Info(ctx, "Incident created successfully",
		prahariLogger.String("incident_id", inc.ID),
		prahariLogger.String("incident_number", inc.IncidentNumber))

	return inc, nil
}

// GetIncident retrieves an incident by its unique identifier.
func (s *Service) GetIncident(ctx context.Context, id string) (*incidentDomain.Incident, error) {
	inc, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve incident: %w", err)
	}
	return inc, nil
}

// UpdateIncident modifies an existing incident and publishes an update event.
func (s *Service) UpdateIncident(ctx context.Context, id string, cmd UpdateIncidentCommand, actorID string) (*incidentDomain.Incident, error) {
	inc, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve incident for update: %w", err)
	}

	oldSnapshot := *inc

	if cmd.Title != "" {
		inc.Title = cmd.Title
	}
	if cmd.Description != "" {
		inc.Description = cmd.Description
	}
	if cmd.CategoryID != "" {
		inc.CategoryID = cmd.CategoryID
	}
	if cmd.SeverityLevel != "" {
		inc.SeverityLevel = cmd.SeverityLevel
	}
	if cmd.PriorityLevel != "" {
		inc.PriorityLevel = cmd.PriorityLevel
	}
	if cmd.LocationID != "" {
		inc.LocationID = cmd.LocationID
	}
	if cmd.LocationDetail != "" {
		inc.LocationDetail = cmd.LocationDetail
	}
	inc.UpdatedAt = time.Now()

	if err := inc.Validate(); err != nil {
		return nil, fmt.Errorf("incident validation failed after update: %w", err)
	}

	if err := s.repo.Update(ctx, inc); err != nil {
		return nil, fmt.Errorf("failed to persist incident update: %w", err)
	}

	// Audit log with before/after
	entry := auditDomain.NewEntry("incident", inc.ID, auditDomain.ActionUpdated, actorID, &oldSnapshot, inc)
	if err := s.audit.Log(ctx, entry); err != nil {
		prahariLogger.Error(ctx, "Failed to write audit log", prahariLogger.Err(err))
	}

	// Publish domain event
	if err := s.events.Publish(ctx, "incident.updated", inc); err != nil {
		prahariLogger.Error(ctx, "Failed to publish incident.updated event", prahariLogger.Err(err))
	}

	return inc, nil
}

// TransitionStatus changes the incident status using the domain state machine,
// rejecting invalid transitions.
func (s *Service) TransitionStatus(ctx context.Context, cmd TransitionStatusCommand) error {
	inc, err := s.repo.FindByID(ctx, cmd.IncidentID)
	if err != nil {
		return fmt.Errorf("failed to retrieve incident for transition: %w", err)
	}

	fromCode := statusDomain.Code(inc.StatusCode)
	toCode := statusDomain.Code(cmd.TargetCode)

	if err := statusDomain.ValidateTransition(fromCode, toCode); err != nil {
		return fmt.Errorf("status transition rejected: %w", err)
	}

	oldStatus := inc.StatusCode
	inc.StatusCode = cmd.TargetCode
	inc.UpdatedAt = time.Now()

	// Apply special lifecycle side effects
	switch toCode {
	case statusDomain.CodeResolved:
		inc.MarkResolved()
	case statusDomain.CodeClosed:
		inc.MarkClosed()
	}

	if err := s.repo.Update(ctx, inc); err != nil {
		return fmt.Errorf("failed to persist status transition: %w", err)
	}

	// Timeline
	evt := timelineDomain.NewEvent(inc.ID, timelineDomain.EventStatusChanged, cmd.ActorID,
		fmt.Sprintf("Status changed: %s → %s", oldStatus, cmd.TargetCode),
		map[string]string{"from": oldStatus, "to": cmd.TargetCode, "reason": cmd.Reason})
	if err := s.timeline.Record(ctx, evt); err != nil {
		prahariLogger.Error(ctx, "Failed to record timeline event", prahariLogger.Err(err))
	}

	// Audit
	entry := auditDomain.NewEntry("incident", inc.ID, auditDomain.ActionStatusChanged, cmd.ActorID,
		map[string]string{"status": oldStatus}, map[string]string{"status": cmd.TargetCode})
	if err := s.audit.Log(ctx, entry); err != nil {
		prahariLogger.Error(ctx, "Failed to write audit log", prahariLogger.Err(err))
	}

	// Publish appropriate domain event based on target status
	eventTopic := fmt.Sprintf("incident.%s", cmd.TargetCode)
	if err := s.events.Publish(ctx, eventTopic, inc); err != nil {
		prahariLogger.Error(ctx, "Failed to publish status change event", prahariLogger.Err(err))
	}

	return nil
}

// DeleteIncident performs a soft delete on an incident.
func (s *Service) DeleteIncident(ctx context.Context, id, actorID string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete incident: %w", err)
	}

	entry := auditDomain.NewEntry("incident", id, auditDomain.ActionDeleted, actorID, nil, nil)
	if err := s.audit.Log(ctx, entry); err != nil {
		prahariLogger.Error(ctx, "Failed to write audit log", prahariLogger.Err(err))
	}

	if err := s.events.Publish(ctx, "incident.deleted", map[string]string{"id": id}); err != nil {
		prahariLogger.Error(ctx, "Failed to publish incident.deleted event", prahariLogger.Err(err))
	}

	return nil
}

// ListIncidents retrieves a paginated list of incidents.
func (s *Service) ListIncidents(ctx context.Context, page, pageSize int) ([]*incidentDomain.Incident, int, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	return s.repo.List(ctx, offset, pageSize)
}
