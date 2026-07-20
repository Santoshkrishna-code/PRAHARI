package assessment

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	prahariLogger "prahari/shared/logger"

	auditDomain "prahari/services/risk/internal/domain/audit"
	policyDomain "prahari/services/risk/internal/domain/policy"
	riskDomain "prahari/services/risk/internal/domain/risk"
	statusDomain "prahari/services/risk/internal/domain/status"
	timelineDomain "prahari/services/risk/internal/domain/timeline"
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
	SubmitForApproval(ctx context.Context, riskID string) error
}

// Service manages risk register lifecycle transitions.
type Service struct {
	repo     riskDomain.Repository
	events   EventPublisher
	audit    AuditLogger
	timeline TimelineRecorder
	workflow WorkflowClient
}

// NewService instantiates Service.
func NewService(
	repo riskDomain.Repository,
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

// generateRiskNumber builds format tags.
func generateRiskNumber() string {
	now := time.Now()
	return fmt.Sprintf("RSK-%d%02d%02d-%s", now.Year(), now.Month(), now.Day(), uuid.New().String()[:8])
}

// CreateRiskAssessment registers a new risk index profile.
func (s *Service) CreateRiskAssessment(ctx context.Context, cmd CreateRiskCommand, actorID string) (*riskDomain.Risk, error) {
	now := time.Now()
	r := &riskDomain.Risk{
		ID:                uuid.New().String(),
		RiskNumber:        generateRiskNumber(),
		AssetID:           cmd.AssetID,
		DepartmentID:      cmd.DepartmentID,
		InherentRiskScore: 0,
		ResidualRiskScore: 0,
		StatusCode:        string(statusDomain.CodeDraft),
		Title:             cmd.Title,
		Description:       cmd.Description,
		CreatedAt:         now,
		UpdatedAt:         now,
	}

	if err := r.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, r); err != nil {
		return nil, err
	}

	evt := timelineDomain.NewEvent(r.ID, timelineDomain.EventCreated, actorID, "Risk assessment register record created", nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("risk", r.ID, auditDomain.ActionCreated, actorID, nil, r)
	_ = s.audit.Log(ctx, auditLog)

	_ = s.events.Publish(ctx, "risk.created", r)

	return r, nil
}

// GetRiskAssessment returns register details.
func (s *Service) GetRiskAssessment(ctx context.Context, id string) (*riskDomain.Risk, error) {
	return s.repo.FindByID(ctx, id)
}

// TransitionStatus manages machine validations.
func (s *Service) TransitionStatus(ctx context.Context, cmd TransitionStatusCommand) error {
	r, err := s.repo.FindByID(ctx, cmd.RiskID)
	if err != nil {
		return err
	}

	fromStatus := statusDomain.Code(r.StatusCode)
	toStatus := statusDomain.Code(cmd.TargetCode)

	if err := statusDomain.ValidateTransition(fromStatus, toStatus); err != nil {
		return err
	}

	// Enforce ISO escalation checks
	if toStatus == statusDomain.CodeReview {
		if err := policyDomain.CheckRiskEscalationRequirements(r); err != nil {
			prahariLogger.Info(ctx, "High risk requires bow-tie evaluation", prahariLogger.String("id", r.ID))
		}
	}

	// Enforce reassessment triggers
	if toStatus == statusDomain.CodeReassess {
		if err := policyDomain.CheckReassessmentParameters(r, 1); err != nil {
			prahariLogger.Info(ctx, "Incidents detected in zone, forcing reassessment", prahariLogger.String("id", r.ID))
		}
	}

	r.StatusCode = cmd.TargetCode
	r.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, r); err != nil {
		return err
	}

	if toStatus == statusDomain.CodeApproval {
		_ = s.workflow.SubmitForApproval(ctx, r.ID)
	}

	evt := timelineDomain.NewEvent(r.ID, timelineDomain.EventType(fmt.Sprintf("RISK_%s", cmd.TargetCode)), cmd.ActorID,
		fmt.Sprintf("Risk assessment transitioned to %s", cmd.TargetCode), nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("risk", r.ID, auditDomain.ActionStatusChanged, cmd.ActorID,
		map[string]string{"status": string(fromStatus)}, map[string]string{"status": cmd.TargetCode})
	_ = s.audit.Log(ctx, auditLog)

	eventTopic := fmt.Sprintf("risk.%s", cmd.TargetCode)
	_ = s.events.Publish(ctx, eventTopic, r)

	return nil
}
}
