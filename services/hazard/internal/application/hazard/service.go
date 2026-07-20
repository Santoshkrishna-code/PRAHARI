package hazard

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	prahariLogger "prahari/shared/logger"

	hazardDomain "prahari/services/hazard/internal/domain/hazard"
	auditDomain "prahari/services/hazard/internal/domain/audit"
	policyDomain "prahari/services/hazard/internal/domain/policy"
	statusDomain "prahari/services/hazard/internal/domain/status"
	timelineDomain "prahari/services/hazard/internal/domain/timeline"
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

// IncidentClient escalates high risks to reactive incidents.
type IncidentClient interface {
	EscalateHazard(ctx context.Context, hazardID, title, description string) error
}

// WorkflowClient requests approval flows.
type WorkflowClient interface {
	StartWorkflow(ctx context.Context, hazardID, workflowType string) error
}

// Service manages hazard lifecycle transitions.
type Service struct {
	repo     hazardDomain.Repository
	events   EventPublisher
	audit    AuditLogger
	timeline TimelineRecorder
	incident IncidentClient
	workflow WorkflowClient
}

// NewService instantiates Service.
func NewService(
	repo hazardDomain.Repository,
	events EventPublisher,
	audit AuditLogger,
	timeline TimelineRecorder,
	incident IncidentClient,
	workflow WorkflowClient,
) *Service {
	return &Service{
		repo:     repo,
		events:   events,
		audit:    audit,
		timeline: timeline,
		incident: incident,
		workflow: workflow,
	}
}

// generateHazardNumber code format generator.
func generateHazardNumber() string {
	now := time.Now()
	return fmt.Sprintf("HAZ-%d%02d%02d-%s", now.Year(), now.Month(), now.Day(), uuid.New().String()[:8])
}

// CreateHazard inserts proactive hazard safety log.
func (s *Service) CreateHazard(ctx context.Context, cmd CreateHazardCommand, actorID string) (*hazardDomain.Hazard, error) {
	now := time.Now()
	h := &hazardDomain.Hazard{
		ID:           uuid.New().String(),
		HazardNumber: generateHazardNumber(),
		AssetID:      cmd.AssetID,
		ContractorID: cmd.ContractorID,
		HazardType:   cmd.HazardType,
		StatusCode:   string(statusDomain.CodeIdentified),
		DepartmentID: cmd.DepartmentID,
		Title:        cmd.Title,
		Description:  cmd.Description,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := h.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, h); err != nil {
		return nil, err
	}

	evt := timelineDomain.NewEvent(h.ID, timelineDomain.EventCreated, actorID, "Proactive hazard identified", nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("hazard", h.ID, auditDomain.ActionCreated, actorID, nil, h)
	_ = s.audit.Log(ctx, auditLog)

	_ = s.events.Publish(ctx, "hazard.created", h)

	return h, nil
}

// GetHazard returns hazard aggregate.
func (s *Service) GetHazard(ctx context.Context, id string) (*hazardDomain.Hazard, error) {
	return s.repo.FindByID(ctx, id)
}

// TransitionStatus validates machine matrix constraints.
func (s *Service) TransitionStatus(ctx context.Context, cmd TransitionStatusCommand) error {
	h, err := s.repo.FindByID(ctx, cmd.HazardID)
	if err != nil {
		return err
	}

	fromStatus := statusDomain.Code(h.StatusCode)
	toStatus := statusDomain.Code(cmd.TargetCode)

	if err := statusDomain.ValidateTransition(fromStatus, toStatus); err != nil {
		return err
	}

	// Enforce residual risk check rules
	if toStatus == statusDomain.CodeVerified {
		if err := policyDomain.ValidateResidualRisk(h); err != nil {
			return err
		}
	}

	// Enforce risk escalation rules
	if toStatus == statusDomain.CodeAssessed {
		if err := policyDomain.CheckEscalationRequirement(h); err != nil {
			prahariLogger.Info(ctx, "Hazard risk score exceeds threshold, auto-escalating to Incident Service", prahariLogger.String("id", h.ID))
			_ = s.incident.EscalateHazard(ctx, h.ID, h.Title, h.Description)
		}
	}

	h.StatusCode = cmd.TargetCode
	h.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, h); err != nil {
		return err
	}

	evt := timelineDomain.NewEvent(h.ID, timelineDomain.EventType(fmt.Sprintf("HAZARD_%s", cmd.TargetCode)), cmd.ActorID,
		fmt.Sprintf("Hazard transitioned to %s", cmd.TargetCode), nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("hazard", h.ID, auditDomain.ActionStatusChanged, cmd.ActorID,
		map[string]string{"status": string(fromStatus)}, map[string]string{"status": cmd.TargetCode})
	_ = s.audit.Log(ctx, auditLog)

	eventTopic := fmt.Sprintf("hazard.%s", cmd.TargetCode)
	_ = s.events.Publish(ctx, eventTopic, h)

	return nil
}
