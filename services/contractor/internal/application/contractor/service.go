package contractor

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	prahariLogger "prahari/shared/logger"

	contractorDomain "prahari/services/contractor/internal/domain/contractor"
	auditDomain "prahari/services/contractor/internal/domain/audit"
	policyDomain "prahari/services/contractor/internal/domain/policy"
	statusDomain "prahari/services/contractor/internal/domain/status"
	timelineDomain "prahari/services/contractor/internal/domain/timeline"
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

// IdentityClient checks supervisor credentials.
type IdentityClient interface {
	VerifySupervisor(ctx context.Context, userID string) (bool, error)
}

// WorkflowClient requests approval flows.
type WorkflowClient interface {
	StartWorkflow(ctx context.Context, contractorID, workflowType string) error
}

// Service manages contractor lifecycle transitions.
type Service struct {
	repo     contractorDomain.Repository
	events   EventPublisher
	audit    AuditLogger
	timeline TimelineRecorder
	identity IdentityClient
	workflow WorkflowClient
}

// NewService instantiates Service.
func NewService(
	repo contractorDomain.Repository,
	events EventPublisher,
	audit AuditLogger,
	timeline TimelineRecorder,
	identity IdentityClient,
	workflow WorkflowClient,
) *Service {
	return &Service{
		repo:     repo,
		events:   events,
		audit:    audit,
		timeline: timeline,
		identity: identity,
		workflow: workflow,
	}
}

// generateContractorNumber code format generator.
func generateContractorNumber() string {
	now := time.Now()
	return fmt.Sprintf("CON-%d%02d%02d-%s", now.Year(), now.Month(), now.Day(), uuid.New().String()[:8])
}

// CreateContractor inserts contractor registry drafts.
func (s *Service) CreateContractor(ctx context.Context, cmd RegisterContractorCommand, actorID string) (*contractorDomain.Contractor, error) {
	now := time.Now()
	c := &contractorDomain.Contractor{
		ID:               uuid.New().String(),
		ContractorNumber: generateContractorNumber(),
		CompanyName:      cmd.CompanyName,
		TaxID:            cmd.TaxID,
		StatusCode:       string(statusDomain.CodeRegistered),
		DepartmentID:     cmd.DepartmentID,
		RegistrationDate: now,
		InsuranceExpiry:  cmd.InsuranceExpiry,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	if err := c.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, c); err != nil {
		return nil, err
	}

	evt := timelineDomain.NewEvent(c.ID, timelineDomain.EventCreated, actorID, "Contractor organization registration draft created", nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("contractor", c.ID, auditDomain.ActionCreated, actorID, nil, c)
	_ = s.audit.Log(ctx, auditLog)

	_ = s.events.Publish(ctx, "contractor.created", c)

	return c, nil
}

// GetContractor returns contractor aggregate.
func (s *Service) GetContractor(ctx context.Context, id string) (*contractorDomain.Contractor, error) {
	return s.repo.FindByID(ctx, id)
}

// TransitionStatus validates machine matrix constraints.
func (s *Service) TransitionStatus(ctx context.Context, cmd TransitionStatusCommand) error {
	c, err := s.repo.FindByID(ctx, cmd.ContractorID)
	if err != nil {
		return err
	}

	fromStatus := statusDomain.Code(c.StatusCode)
	toStatus := statusDomain.Code(cmd.TargetCode)

	if err := statusDomain.ValidateTransition(fromStatus, toStatus); err != nil {
		return err
	}

	// Enforce safety insurance checks
	if toStatus == statusDomain.CodeActive {
		if err := policyDomain.ValidateContractorInsurance(c); err != nil {
			return err
		}
	}

	// On approval trigger onboarding workflow reviews
	if toStatus == statusDomain.CodeApproved {
		prahariLogger.Info(ctx, "Contractor company approved, triggering onboarding review workflow", prahariLogger.String("id", c.ID))
		_ = s.workflow.StartWorkflow(ctx, c.ID, "contractor_onboarding_approval")
	}

	c.StatusCode = cmd.TargetCode
	c.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, c); err != nil {
		return err
	}

	evt := timelineDomain.NewEvent(c.ID, timelineDomain.EventType(fmt.Sprintf("CONTRACTOR_%s", cmd.TargetCode)), cmd.ActorID,
		fmt.Sprintf("Contractor transitioned to %s", cmd.TargetCode), nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("contractor", c.ID, auditDomain.ActionStatusChanged, cmd.ActorID,
		map[string]string{"status": string(fromStatus)}, map[string]string{"status": cmd.TargetCode})
	_ = s.audit.Log(ctx, auditLog)

	eventTopic := fmt.Sprintf("contractor.%s", cmd.TargetCode)
	_ = s.events.Publish(ctx, eventTopic, c)

	return nil
}
