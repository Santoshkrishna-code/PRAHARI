package asset

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	prahariLogger "prahari/shared/logger"

	assetDomain "prahari/services/asset/internal/domain/asset"
	auditDomain "prahari/services/asset/internal/domain/audit"
	policyDomain "prahari/services/asset/internal/domain/policy"
	statusDomain "prahari/services/asset/internal/domain/status"
	timelineDomain "prahari/services/asset/internal/domain/timeline"
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

// InspectionClient checks safety verification status.
type InspectionClient interface {
	HasPSSRPassed(ctx context.Context, assetID string) (bool, error)
}

// WorkflowClient requests approval flows.
type WorkflowClient interface {
	StartWorkflow(ctx context.Context, assetID, workflowType string) error
}

// Service manages asset registration, status changes and lifecycle enforcement.
type Service struct {
	repo       assetDomain.Repository
	events     EventPublisher
	audit      AuditLogger
	timeline   TimelineRecorder
	inspection InspectionClient
	workflow   WorkflowClient
}

// NewService instantiates Service.
func NewService(
	repo assetDomain.Repository,
	events EventPublisher,
	audit AuditLogger,
	timeline TimelineRecorder,
	inspection InspectionClient,
	workflow WorkflowClient,
) *Service {
	return &Service{
		repo:       repo,
		events:     events,
		audit:      audit,
		timeline:   timeline,
		inspection: inspection,
		workflow:   workflow,
	}
}

// generateAssetNumber code format generator.
func generateAssetNumber() string {
	now := time.Now()
	return fmt.Sprintf("AST-%d%02d%02d-%s", now.Year(), now.Month(), now.Day(), uuid.New().String()[:8])
}

// RegisterAsset registers new equipment in the system of record.
func (s *Service) RegisterAsset(ctx context.Context, cmd RegisterAssetCommand, actorID string) (*assetDomain.Asset, error) {
	now := time.Now()
	a := &assetDomain.Asset{
		ID:                uuid.New().String(),
		AssetNumber:       generateAssetNumber(),
		Name:              cmd.Name,
		Description:       cmd.Description,
		SerialNumber:      cmd.SerialNumber,
		LifecycleStatus:   string(statusDomain.CodeRegistered),
		OperationalStatus: assetDomain.OpStopped,
		DepartmentID:      cmd.DepartmentID,
		LocationID:        cmd.LocationID,
		CategoryID:        cmd.CategoryID,
		TypeID:            cmd.TypeID,
		PurchaseDate:      now,
		HealthScore:       100.0,
		ConditionScore:    100.0,
		RemainingUsefulLife: 87600.0, // default 10 years in hours
		FailureProbability: 0.0,
		CreatedAt:         now,
		UpdatedAt:         now,
	}

	if err := a.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, a); err != nil {
		return nil, err
	}

	evt := timelineDomain.NewEvent(a.ID, timelineDomain.EventCreated, actorID, "Asset profiles registered in safety inventory", nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("asset", a.ID, auditDomain.ActionCreated, actorID, nil, a)
	_ = s.audit.Log(ctx, auditLog)

	_ = s.events.Publish(ctx, "asset.created", a)

	return a, nil
}

// GetAsset returns asset profile details.
func (s *Service) GetAsset(ctx context.Context, id string) (*assetDomain.Asset, error) {
	return s.repo.FindByID(ctx, id)
}

// TransitionLifecycle status checks transition validations.
func (s *Service) TransitionLifecycle(ctx context.Context, cmd TransitionStatusCommand) error {
	a, err := s.repo.FindByID(ctx, cmd.AssetID)
	if err != nil {
		return err
	}

	fromStatus := statusDomain.Code(a.LifecycleStatus)
	toStatus := statusDomain.Code(cmd.TargetCode)

	if err := statusDomain.ValidateTransition(fromStatus, toStatus); err != nil {
		return err
	}

	// Enforce commissioning safety rules
	if toStatus == statusDomain.CodeCommissioned {
		hasPassed, _ := s.inspection.HasPSSRPassed(ctx, a.ID)
		if err := policyDomain.ValidateCommissionPreconditions(a, hasPassed); err != nil {
			return err
		}
	}

	a.LifecycleStatus = cmd.TargetCode
	a.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, a); err != nil {
		return err
	}

	evt := timelineDomain.NewEvent(a.ID, timelineDomain.EventType(fmt.Sprintf("ASSET_%s", cmd.TargetCode)), cmd.ActorID,
		fmt.Sprintf("Lifecycle status changed to %s", cmd.TargetCode), nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("asset", a.ID, auditDomain.ActionStatusChanged, cmd.ActorID,
		map[string]string{"status": string(fromStatus)}, map[string]string{"status": cmd.TargetCode})
	_ = s.audit.Log(ctx, auditLog)

	eventTopic := fmt.Sprintf("asset.%s", cmd.TargetCode)
	_ = s.events.Publish(ctx, eventTopic, a)

	return nil
}
