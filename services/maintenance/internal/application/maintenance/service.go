package maintenance

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	prahariLogger "prahari/shared/logger"

	maintenanceDomain "prahari/services/maintenance/internal/domain/maintenance"
	auditDomain "prahari/services/maintenance/internal/domain/audit"
	policyDomain "prahari/services/maintenance/internal/domain/policy"
	statusDomain "prahari/services/maintenance/internal/domain/status"
	timelineDomain "prahari/services/maintenance/internal/domain/timeline"
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

// AssetClient checks physical equipment assets existence.
type AssetClient interface {
	VerifyAsset(ctx context.Context, assetID string) (bool, error)
}

// PermitClient checks safety permit verifications status.
type PermitClient interface {
	ValidatePermitStatus(ctx context.Context, assetID string) (bool, error)
}

// WorkflowClient requests approval flows.
type WorkflowClient interface {
	StartWorkflow(ctx context.Context, maintenanceID, workflowType string) error
}

// Service manages maintenance lifecycle transitions.
type Service struct {
	repo     maintenanceDomain.Repository
	events   EventPublisher
	audit    AuditLogger
	timeline TimelineRecorder
	asset    AssetClient
	permit   PermitClient
	workflow WorkflowClient
}

// NewService instantiates Service.
func NewService(
	repo maintenanceDomain.Repository,
	events EventPublisher,
	audit AuditLogger,
	timeline TimelineRecorder,
	asset AssetClient,
	permit PermitClient,
	workflow WorkflowClient,
) *Service {
	return &Service{
		repo:     repo,
		events:   events,
		audit:    audit,
		timeline: timeline,
		asset:    asset,
		permit:   permit,
		workflow: workflow,
	}
}

// generateMaintenanceNumber code format generator.
func generateMaintenanceNumber() string {
	now := time.Now()
	return fmt.Sprintf("MNT-%d%02d%02d-%s", now.Year(), now.Month(), now.Day(), uuid.New().String()[:8])
}

// CreateMaintenance inserts safety checklist audits request.
func (s *Service) CreateMaintenance(ctx context.Context, cmd CreateMaintenanceCommand, actorID string) (*maintenanceDomain.Maintenance, error) {
	// Verify asset exists in EAM master record
	exists, _ := s.asset.VerifyAsset(ctx, cmd.AssetID)
	if !exists {
		return nil, fmt.Errorf("asset ID %s does not exist in master record", cmd.AssetID)
	}

	now := time.Now()
	m := &maintenanceDomain.Maintenance{
		ID:                 uuid.New().String(),
		MaintenanceNumber:  generateMaintenanceNumber(),
		AssetID:            cmd.AssetID,
		MaintenanceType:    cmd.MaintenanceType,
		Priority:           maintenanceDomain.Priority(cmd.Priority),
		DepartmentID:       cmd.DepartmentID,
		Title:              cmd.Title,
		Description:        cmd.Description,
		StatusCode:         string(statusDomain.CodeDraft),
		TotalEstimatedCost: cmd.EstimatedCost,
		CreatedAt:          now,
		UpdatedAt:          now,
	}

	if err := m.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, m); err != nil {
		return nil, err
	}

	evt := timelineDomain.NewEvent(m.ID, timelineDomain.EventCreated, actorID, "Maintenance draft work order created", nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("maintenance", m.ID, auditDomain.ActionCreated, actorID, nil, m)
	_ = s.audit.Log(ctx, auditLog)

	_ = s.events.Publish(ctx, "maintenance.created", m)

	return m, nil
}

// GetMaintenance returns maintenance aggregate.
func (s *Service) GetMaintenance(ctx context.Context, id string) (*maintenanceDomain.Maintenance, error) {
	return s.repo.FindByID(ctx, id)
}

// TransitionStatus validates machine matrix constraints.
func (s *Service) TransitionStatus(ctx context.Context, cmd TransitionStatusCommand) error {
	m, err := s.repo.FindByID(ctx, cmd.MaintenanceID)
	if err != nil {
		return err
	}

	fromStatus := statusDomain.Code(m.StatusCode)
	toStatus := statusDomain.Code(cmd.TargetCode)

	if err := statusDomain.ValidateTransition(fromStatus, toStatus); err != nil {
		return err
	}

	// Enforce safety permits rules
	if toStatus == statusDomain.CodeInProgress {
		hasPermit, _ := s.permit.ValidatePermitStatus(ctx, m.AssetID)
		if err := policyDomain.ValidateWorkPermitPrecondition(m, hasPermit); err != nil {
			return err
		}
	}

	// Enforce workflow cost approvals limits
	if toStatus == statusDomain.CodePlanned {
		if err := policyDomain.ValidateMaintenanceCost(m); err != nil {
			prahariLogger.Info(ctx, "Maintenance plan exceeds cost limit, starting review workflow", prahariLogger.String("id", m.ID))
			_ = s.workflow.StartWorkflow(ctx, m.ID, "maintenance_cost_approval")
		}
	}

	m.StatusCode = cmd.TargetCode
	m.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, m); err != nil {
		return err
	}

	evt := timelineDomain.NewEvent(m.ID, timelineDomain.EventType(fmt.Sprintf("MAINTENANCE_%s", cmd.TargetCode)), cmd.ActorID,
		fmt.Sprintf("Work order transitioned to %s", cmd.TargetCode), nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("maintenance", m.ID, auditDomain.ActionStatusChanged, cmd.ActorID,
		map[string]string{"status": string(fromStatus)}, map[string]string{"status": cmd.TargetCode})
	_ = s.audit.Log(ctx, auditLog)

	eventTopic := fmt.Sprintf("maintenance.%s", cmd.TargetCode)
	_ = s.events.Publish(ctx, eventTopic, m)

	return nil
}
