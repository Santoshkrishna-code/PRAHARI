package approval

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	approvalDomain "prahari/services/permit/internal/domain/approval"
	auditDomain "prahari/services/permit/internal/domain/audit"
	timelineDomain "prahari/services/permit/internal/domain/timeline"
)

// Repository defines the port contract for approvals.
type Repository interface {
	Create(ctx context.Context, approval *approvalDomain.Approval) error
	FindByID(ctx context.Context, id string) (*approvalDomain.Approval, error)
	FindByPermitID(ctx context.Context, permitID string) ([]*approvalDomain.Approval, error)
	Update(ctx context.Context, approval *approvalDomain.Approval) error
}

// IdentityClient defines the port for validating authority existence.
type IdentityClient interface {
	UserExists(ctx context.Context, userID string) (bool, error)
}

// EventPublisher defines the port for sending Kafka events.
type EventPublisher interface {
	Publish(ctx context.Context, topic string, payload interface{}) error
}

// AuditLogger defines the port for logging change entries.
type AuditLogger interface {
	Log(ctx context.Context, entry *auditDomain.Entry) error
}

// TimelineRecorder defines the port for registering milestones.
type TimelineRecorder interface {
	Record(ctx context.Context, event *timelineDomain.Event) error
}

// PermitUpdater defines the port for modifying a permit's status.
type PermitUpdater interface {
	UpdateStatus(ctx context.Context, permitID, status string) error
}

// Service manages multi-level workflow signs.
type Service struct {
	repo     Repository
	identity IdentityClient
	events   EventPublisher
	audit    AuditLogger
	timeline TimelineRecorder
	permit   PermitUpdater
}

// NewService instantiates an Approval Application Service.
func NewService(
	repo Repository,
	identity IdentityClient,
	events EventPublisher,
	audit AuditLogger,
	timeline TimelineRecorder,
	permit PermitUpdater,
) *Service {
	return &Service{
		repo:     repo,
		identity: identity,
		events:   events,
		audit:    audit,
		timeline: timeline,
		permit:   permit,
	}
}

// SubmitApproval records an approval decision.
func (s *Service) SubmitApproval(ctx context.Context, permitID, approverID string, role approvalDomain.Role, decision approvalDomain.Decision, comment, signature string) (*approvalDomain.Approval, error) {
	// Verify approver exists
	exists, err := s.identity.UserExists(ctx, approverID)
	if err != nil || !exists {
		return nil, errors.New("approver user does not exist")
	}

	app := &approvalDomain.Approval{
		ID:            uuid.New().String(),
		PermitID:      permitID,
		ApproverID:    approverID,
		ApproverRole:  role,
		Decision:      decision,
		Comments:      comment,
		SequenceOrder: 1,
	}

	if decision == approvalDomain.DecisionApproved {
		app.Approve(signature)
	} else {
		app.Reject(comment)
	}

	if err := app.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, app); err != nil {
		return nil, err
	}

	// Timeline milestone
	var timelineType timelineDomain.EventType
	if decision == approvalDomain.DecisionApproved {
		timelineType = timelineDomain.EventApproved
	} else {
		timelineType = timelineDomain.EventCancelled // Rejections cancel/send back
	}

	evt := timelineDomain.NewEvent(permitID, timelineType, approverID, fmt.Sprintf("Permit signoff outcome: %s", decision), nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("approval", app.ID, auditDomain.ActionApproved, approverID, nil, app)
	_ = s.audit.Log(ctx, auditLog)

	_ = s.events.Publish(ctx, "permit.approved", app)

	return app, nil
}

// GetApprovals lists signs for a permit.
func (s *Service) GetApprovals(ctx context.Context, permitID string) ([]*approvalDomain.Approval, error) {
	return s.repo.FindByPermitID(ctx, permitID)
}
