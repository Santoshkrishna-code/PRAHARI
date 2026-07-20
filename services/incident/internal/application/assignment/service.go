package assignment

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	prahariLogger "prahari/shared/logger"

	assignmentDomain "prahari/services/incident/internal/domain/assignment"
	auditDomain "prahari/services/incident/internal/domain/audit"
	timelineDomain "prahari/services/incident/internal/domain/timeline"
)

// Repository defines the persistence port for assignment operations.
type Repository interface {
	Create(ctx context.Context, assignment *assignmentDomain.Assignment) error
	FindByID(ctx context.Context, id string) (*assignmentDomain.Assignment, error)
	FindByIncidentID(ctx context.Context, incidentID string) ([]*assignmentDomain.Assignment, error)
	Update(ctx context.Context, assignment *assignmentDomain.Assignment) error
}

// IncidentUpdater defines the port for updating incident status during assignment.
type IncidentUpdater interface {
	UpdateAssignee(ctx context.Context, incidentID, assigneeID string) error
	UpdateStatus(ctx context.Context, incidentID, statusCode string) error
}

// IdentityClient defines the port for validating user identity during assignment.
type IdentityClient interface {
	UserExists(ctx context.Context, userID string) (bool, error)
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

// Service orchestrates incident assignment operations.
type Service struct {
	repo     Repository
	incident IncidentUpdater
	identity IdentityClient
	events   EventPublisher
	audit    AuditLogger
	timeline TimelineRecorder
}

// NewService constructs a Service with all required dependencies injected.
func NewService(
	repo Repository,
	incident IncidentUpdater,
	identity IdentityClient,
	events EventPublisher,
	audit AuditLogger,
	timeline TimelineRecorder,
) *Service {
	return &Service{
		repo:     repo,
		incident: incident,
		identity: identity,
		events:   events,
		audit:    audit,
		timeline: timeline,
	}
}

// AssignIncident assigns a user to an incident in a specific role.
// It validates the assignee exists via the Identity Service before persisting.
func (s *Service) AssignIncident(ctx context.Context, incidentID, assigneeID, assignerID string, role assignmentDomain.Role) (*assignmentDomain.Assignment, error) {
	// Validate assignee exists via Identity Service
	exists, err := s.identity.UserExists(ctx, assigneeID)
	if err != nil {
		return nil, fmt.Errorf("failed to validate assignee identity: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("assignee user %s does not exist", assigneeID)
	}

	assignment := &assignmentDomain.Assignment{
		ID:         uuid.New().String(),
		IncidentID: incidentID,
		AssigneeID: assigneeID,
		AssignerID: assignerID,
		Role:       role,
		AssignedAt: time.Now(),
		IsActive:   true,
	}

	if err := assignment.Validate(); err != nil {
		return nil, fmt.Errorf("assignment validation failed: %w", err)
	}

	if err := s.repo.Create(ctx, assignment); err != nil {
		return nil, fmt.Errorf("failed to persist assignment: %w", err)
	}

	// Update incident assignee and status
	if err := s.incident.UpdateAssignee(ctx, incidentID, assigneeID); err != nil {
		prahariLogger.Error(ctx, "Failed to update incident assignee", prahariLogger.Err(err))
	}
	if err := s.incident.UpdateStatus(ctx, incidentID, "ASSIGNED"); err != nil {
		prahariLogger.Error(ctx, "Failed to transition incident to ASSIGNED", prahariLogger.Err(err))
	}

	// Timeline
	evt := timelineDomain.NewEvent(incidentID, timelineDomain.EventAssigned, assignerID,
		fmt.Sprintf("Incident assigned to %s as %s", assigneeID, role),
		map[string]string{"assignee_id": assigneeID, "role": string(role)})
	if err := s.timeline.Record(ctx, evt); err != nil {
		prahariLogger.Error(ctx, "Failed to record timeline event", prahariLogger.Err(err))
	}

	// Audit
	entry := auditDomain.NewEntry("assignment", assignment.ID, auditDomain.ActionAssigned, assignerID, nil, assignment)
	if err := s.audit.Log(ctx, entry); err != nil {
		prahariLogger.Error(ctx, "Failed to write audit log", prahariLogger.Err(err))
	}

	// Publish event
	if err := s.events.Publish(ctx, "incident.assigned", assignment); err != nil {
		prahariLogger.Error(ctx, "Failed to publish incident.assigned event", prahariLogger.Err(err))
	}

	return assignment, nil
}

// AcceptAssignment records that the assignee has acknowledged the assignment.
func (s *Service) AcceptAssignment(ctx context.Context, assignmentID, actorID string) error {
	assignment, err := s.repo.FindByID(ctx, assignmentID)
	if err != nil {
		return fmt.Errorf("failed to retrieve assignment: %w", err)
	}

	assignment.Accept()

	if err := s.repo.Update(ctx, assignment); err != nil {
		return fmt.Errorf("failed to persist assignment acceptance: %w", err)
	}

	prahariLogger.Info(ctx, "Assignment accepted",
		prahariLogger.String("assignment_id", assignmentID),
		prahariLogger.String("actor_id", actorID))

	return nil
}

// GetAssignments retrieves all assignments for an incident.
func (s *Service) GetAssignments(ctx context.Context, incidentID string) ([]*assignmentDomain.Assignment, error) {
	return s.repo.FindByIncidentID(ctx, incidentID)
}
