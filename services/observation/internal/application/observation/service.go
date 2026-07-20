package observation

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	prahariLogger "prahari/shared/logger"

	auditDomain "prahari/services/observation/internal/domain/audit"
	observationDomain "prahari/services/observation/internal/domain/observation"
	policyDomain "prahari/services/observation/internal/domain/policy"
	statusDomain "prahari/services/observation/internal/domain/status"
	timelineDomain "prahari/services/observation/internal/domain/timeline"
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

// HazardClient escalates unsafe behaviors to hazards.
type HazardClient interface {
	CreateHazard(ctx context.Context, observationID, title, description string) error
}

// IncidentClient escalates critical unsafe acts to incidents.
type IncidentClient interface {
	EscalateObservation(ctx context.Context, observationID, title, description string) error
}

// Service manages observation lifecycle transitions.
type Service struct {
	repo     observationDomain.Repository
	events   EventPublisher
	audit    AuditLogger
	timeline TimelineRecorder
	hazard   HazardClient
	incident IncidentClient
}

// NewService instantiates Service.
func NewService(
	repo observationDomain.Repository,
	events EventPublisher,
	audit AuditLogger,
	timeline TimelineRecorder,
	hazard HazardClient,
	incident IncidentClient,
) *Service {
	return &Service{
		repo:     repo,
		events:   events,
		audit:    audit,
		timeline: timeline,
		hazard:   hazard,
		incident: incident,
	}
}

// generateObservationNumber code format generator.
func generateObservationNumber() string {
	now := time.Now()
	return fmt.Sprintf("OB-%d%02d%02d-%s", now.Year(), now.Month(), now.Day(), uuid.New().String()[:8])
}

// CreateObservation inserts BBS observation log.
func (s *Service) CreateObservation(ctx context.Context, cmd CreateObservationCommand, actorID string) (*observationDomain.Observation, error) {
	now := time.Now()
	o := &observationDomain.Observation{
		ID:                uuid.New().String(),
		ObservationNumber: generateObservationNumber(),
		AssetID:           cmd.AssetID,
		ContractorID:      cmd.ContractorID,
		ObservationType:   cmd.ObservationType,
		StatusCode:        string(statusDomain.CodeRecorded),
		DepartmentID:      cmd.DepartmentID,
		Title:             cmd.Title,
		Description:       cmd.Description,
		CreatedAt:         now,
		UpdatedAt:         now,
	}

	if err := o.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, o); err != nil {
		return nil, err
	}

	evt := timelineDomain.NewEvent(o.ID, timelineDomain.EventCreated, actorID, "Behavior safety observation recorded", nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("observation", o.ID, auditDomain.ActionCreated, actorID, nil, o)
	_ = s.audit.Log(ctx, auditLog)

	_ = s.events.Publish(ctx, "observation.created", o)

	return o, nil
}

// GetObservation returns observation aggregate.
func (s *Service) GetObservation(ctx context.Context, id string) (*observationDomain.Observation, error) {
	return s.repo.FindByID(ctx, id)
}

// TransitionStatus validates machine matrix constraints.
func (s *Service) TransitionStatus(ctx context.Context, cmd TransitionStatusCommand) error {
	o, err := s.repo.FindByID(ctx, cmd.ObservationID)
	if err != nil {
		return err
	}

	fromStatus := statusDomain.Code(o.StatusCode)
	toStatus := statusDomain.Code(cmd.TargetCode)

	if err := statusDomain.ValidateTransition(fromStatus, toStatus); err != nil {
		return err
	}

	// Enforce BBS policy checks
	if toStatus == statusDomain.CodeCoaching {
		if err := policyDomain.CheckBBSCoachingRequirements(o); err != nil {
			prahariLogger.Info(ctx, "Unsafe behavior triggers mandatory BBS coaching", prahariLogger.String("id", o.ID))
		}
	}

	// Enforce escalation checks
	if toStatus == statusDomain.CodeReviewed {
		if err := policyDomain.VerifyEscalationParameters(o); err != nil {
			prahariLogger.Info(ctx, "Critical unsafe behavior observed, escalating to Hazard Service", prahariLogger.String("id", o.ID))
			_ = s.hazard.CreateHazard(ctx, o.ID, o.Title, o.Description)
		}
	}

	o.StatusCode = cmd.TargetCode
	o.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, o); err != nil {
		return err
	}

	evt := timelineDomain.NewEvent(o.ID, timelineDomain.EventType(fmt.Sprintf("OBSERVATION_%s", cmd.TargetCode)), cmd.ActorID,
		fmt.Sprintf("Observation transitioned to %s", cmd.TargetCode), nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("observation", o.ID, auditDomain.ActionStatusChanged, cmd.ActorID,
		map[string]string{"status": string(fromStatus)}, map[string]string{"status": cmd.TargetCode})
	_ = s.audit.Log(ctx, auditLog)

	eventTopic := fmt.Sprintf("observation.%s", cmd.TargetCode)
	_ = s.events.Publish(ctx, eventTopic, o)

	return nil
}
