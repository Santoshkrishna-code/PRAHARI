package nearmiss

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	prahariLogger "prahari/shared/logger"

	auditDomain "prahari/services/nearmiss/internal/domain/audit"
	nearmissDomain "prahari/services/nearmiss/internal/domain/nearmiss"
	policyDomain "prahari/services/nearmiss/internal/domain/policy"
	statusDomain "prahari/services/nearmiss/internal/domain/status"
	timelineDomain "prahari/services/nearmiss/internal/domain/timeline"
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

// IncidentClient escalates serious potential issues to incident ticketing.
type IncidentClient interface {
	EscalateNearMiss(ctx context.Context, nearmissID, title, description string) error
}

// HazardClient creates proactive hazards.
type HazardClient interface {
	CreateHazard(ctx context.Context, nearmissID, title, description string) error
}

// Service manages nearmiss lifecycle transitions.
type Service struct {
	repo     nearmissDomain.Repository
	events   EventPublisher
	audit    AuditLogger
	timeline TimelineRecorder
	incident IncidentClient
	hazard   HazardClient
}

// NewService instantiates Service.
func NewService(
	repo nearmissDomain.Repository,
	events EventPublisher,
	audit AuditLogger,
	timeline TimelineRecorder,
	incident IncidentClient,
	hazard HazardClient,
) *Service {
	return &Service{
		repo:     repo,
		events:   events,
		audit:    audit,
		timeline: timeline,
		incident: incident,
		hazard:   hazard,
	}
}

// generateNearMissNumber code format generator.
func generateNearMissNumber() string {
	now := time.Now()
	return fmt.Sprintf("NM-%d%02d%02d-%s", now.Year(), now.Month(), now.Day(), uuid.New().String()[:8])
}

// CreateNearMiss inserts near miss safety log.
func (s *Service) CreateNearMiss(ctx context.Context, cmd CreateNearMissCommand, actorID string) (*nearmissDomain.NearMiss, error) {
	now := time.Now()
	nm := &nearmissDomain.NearMiss{
		ID:             uuid.New().String(),
		NearMissNumber: generateNearMissNumber(),
		AssetID:        cmd.AssetID,
		ContractorID:   cmd.ContractorID,
		Classification: cmd.Classification,
		SeverityLevel:  cmd.SeverityLevel,
		StatusCode:     string(statusDomain.CodeReported),
		DepartmentID:   cmd.DepartmentID,
		Title:          cmd.Title,
		Description:    cmd.Description,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := nm.Validate(); err != nil {
		return nil, err
	}

	// Enforce reporter privacy anonymity checklist checks
	if err := policyDomain.ValidateReporterPrivacy(nm, cmd.IsAnonymous, actorID); err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, nm); err != nil {
		return nil, err
	}

	evt := timelineDomain.NewEvent(nm.ID, timelineDomain.EventCreated, actorID, "Near miss reported", nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("nearmiss", nm.ID, auditDomain.ActionCreated, actorID, nil, nm)
	_ = s.audit.Log(ctx, auditLog)

	_ = s.events.Publish(ctx, "nearmiss.created", nm)

	return nm, nil
}

// GetNearMiss returns nearmiss aggregate.
func (s *Service) GetNearMiss(ctx context.Context, id string) (*nearmissDomain.NearMiss, error) {
	return s.repo.FindByID(ctx, id)
}

// TransitionStatus validates machine matrix constraints.
func (s *Service) TransitionStatus(ctx context.Context, cmd TransitionStatusCommand) error {
	nm, err := s.repo.FindByID(ctx, cmd.NearMissID)
	if err != nil {
		return err
	}

	fromStatus := statusDomain.Code(nm.StatusCode)
	toStatus := statusDomain.Code(cmd.TargetCode)

	if err := statusDomain.ValidateTransition(fromStatus, toStatus); err != nil {
		return err
	}

	// Enforce escalation checking policy
	if toStatus == statusDomain.CodeEscalated {
		if err := policyDomain.ValidateNearMissEscalation(nm); err != nil {
			prahariLogger.Info(ctx, "Near miss requires escalation, triggering incident ticket creation", prahariLogger.String("id", nm.ID))
			_ = s.incident.EscalateNearMiss(ctx, nm.ID, nm.Title, nm.Description)
		}
	}

	nm.StatusCode = cmd.TargetCode
	nm.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, nm); err != nil {
		return err
	}

	evt := timelineDomain.NewEvent(nm.ID, timelineDomain.EventType(fmt.Sprintf("NEARMISS_%s", cmd.TargetCode)), cmd.ActorID,
		fmt.Sprintf("Near miss transitioned to %s", cmd.TargetCode), nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("nearmiss", nm.ID, auditDomain.ActionStatusChanged, cmd.ActorID,
		map[string]string{"status": string(fromStatus)}, map[string]string{"status": cmd.TargetCode})
	_ = s.audit.Log(ctx, auditLog)

	eventTopic := fmt.Sprintf("nearmiss.%s", cmd.TargetCode)
	_ = s.events.Publish(ctx, eventTopic, nm)

	return nil
}
