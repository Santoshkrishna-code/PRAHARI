package clearance

import (
	"context"
	"errors"
	"time"

	"prahari/services/occupational-health/internal/domain/audittrail"
	"prahari/services/occupational-health/internal/domain/fitnessassessment"
	"prahari/services/occupational-health/internal/domain/healthprofile"
	"prahari/services/occupational-health/internal/domain/medicalclearance"
	"prahari/services/occupational-health/internal/domain/restriction"
	"prahari/services/occupational-health/internal/domain/status"
	"prahari/services/occupational-health/internal/domain/timeline"
)

// Repository manages main profile states.
type Repository interface {
	SaveProfile(ctx context.Context, p *healthprofile.HealthProfile) error
	GetProfile(ctx context.Context, id string) (*healthprofile.HealthProfile, error)
	GetProfileByWorker(ctx context.Context, workerID string) (*healthprofile.HealthProfile, error)
	SaveClearance(ctx context.Context, c *medicalclearance.MedicalClearance) error
	SaveAssessment(ctx context.Context, a *fitnessassessment.FitnessAssessment) error
	SaveRestriction(ctx context.Context, r *restriction.MedicalRestriction) error
	GetActiveRestrictions(ctx context.Context, profileID string) ([]restriction.MedicalRestriction, error)
}

// EventPublisher defines event outbound port.
type EventPublisher interface {
	Publish(ctx context.Context, topic string, payload interface{}) error
}

// AuditTrailLogger captures structural audit logs.
type AuditTrailLogger interface {
	LogAuditTrail(ctx context.Context, entry *audittrail.Entry) error
}

// TimelineLogger tracks history.
type TimelineLogger interface {
	LogTimeline(ctx context.Context, e *timeline.Event) error
}

// WorkflowEngine defines gRPC proxy.
type WorkflowEngine interface {
	SubmitClearanceApproval(ctx context.Context, clearanceID string) error
}

type Service struct {
	repo     Repository
	events   EventPublisher
	trail    AuditTrailLogger
	timeline TimelineLogger
	workflow WorkflowEngine
}

func NewService(repo Repository, events EventPublisher, trail AuditTrailLogger, timeline TimelineLogger, workflow WorkflowEngine) *Service {
	return &Service{
		repo:     repo,
		events:   events,
		trail:    trail,
		timeline: timeline,
		workflow: workflow,
	}
}

func (s *Service) CreateProfile(ctx context.Context, p *healthprofile.HealthProfile) error {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	p.MedicalStatus = string(status.CodeScheduled)
	p.ClearanceStatus = "UNFIT"

	if err := p.Validate(); err != nil {
		return err
	}

	if err := s.repo.SaveProfile(ctx, p); err != nil {
		return err
	}

	_ = s.events.Publish(ctx, "health.profile.created", p)

	trailLog := audittrail.NewEntry("CREATE", "health_profile", p.ID, "SYSTEM", nil, map[string]string{"worker_id": p.WorkerID})
	_ = s.trail.LogAuditTrail(ctx, trailLog)

	return nil
}

type TransitionStatusCommand struct {
	ProfileID  string `json:"profile_id"`
	TargetCode string `json:"target_code"`
	ActorID    string `json:"actor_id"`
	Reason     string `json:"reason,omitempty"`
}

func (s *Service) TransitionStatus(ctx context.Context, cmd TransitionStatusCommand) error {
	p, err := s.repo.GetProfile(ctx, cmd.ProfileID)
	if err != nil {
		return err
	}
	if p == nil {
		return errors.New("health profile not found")
	}

	fromStatus := status.Code(p.MedicalStatus)
	toStatus := status.Code(cmd.TargetCode)

	if err := status.ValidateTransition(fromStatus, toStatus); err != nil {
		return err
	}

	p.MedicalStatus = string(toStatus)
	p.UpdatedAt = time.Now()

	// Business rule mapping: if transitioning to MEDICAL_CLEARANCE, state clearance status
	if toStatus == status.CodeMedicalClearance {
		p.ClearanceStatus = "CLEARED"
	} else if toStatus == status.CodeTemporarilyUnfit || toStatus == status.CodePermanentlyUnfit {
		p.ClearanceStatus = "UNFIT"
	} else if toStatus == status.CodeRestricted {
		p.ClearanceStatus = "RESTRICTED"
	}

	if err := s.repo.SaveProfile(ctx, p); err != nil {
		return err
	}

	// Trigger specific events
	if toStatus == status.CodeMedicalClearance {
		_ = s.events.Publish(ctx, "medical.clearance.granted", p)
	} else if toStatus == status.CodeRestricted {
		_ = s.events.Publish(ctx, "medical.restriction.applied", p)
	} else if toStatus == status.CodeFitnessAssessment {
		_ = s.events.Publish(ctx, "fitness.assessed", p)
	}

	eventLog := timeline.NewEvent(p.ID, "STATUS_TRANSITION", cmd.ActorID, "Status transitioned successfully", map[string]string{
		"from": string(fromStatus),
		"to":   string(toStatus),
	})
	_ = s.timeline.LogTimeline(ctx, eventLog)

	return nil
}

func (s *Service) AddRestriction(ctx context.Context, r *restriction.MedicalRestriction) error {
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
	if err := r.Validate(); err != nil {
		return err
	}

	if err := s.repo.SaveRestriction(ctx, r); err != nil {
		return err
	}

	// Automatically transition health profile to RESTRICTED state
	_ = s.TransitionStatus(ctx, TransitionStatusCommand{
		ProfileID:  r.HealthProfileID,
		TargetCode: string(status.CodeRestricted),
		ActorID:    "SYSTEM",
		Reason:     "Active medical restriction applied",
	})

	return nil
}
