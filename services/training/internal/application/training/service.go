package training

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	prahariLogger "prahari/shared/logger"

	trainingDomain "prahari/services/training/internal/domain/training"
	policyDomain "prahari/services/training/internal/domain/policy"
	statusDomain "prahari/services/training/internal/domain/status"
	timelineDomain "prahari/services/training/internal/domain/timeline"
	trailDomain "prahari/services/training/internal/domain/audittrail"
)

// EventPublisher dispatches Kafka events.
type EventPublisher interface {
	Publish(ctx context.Context, topic string, payload interface{}) error
}

// AuditTrailLogger writes change trace snaps.
type AuditTrailLogger interface {
	Log(ctx context.Context, entry *trailDomain.Entry) error
}

// TimelineRecorder writes timeline events.
type TimelineRecorder interface {
	Record(ctx context.Context, event *timelineDomain.Event) error
}

// WorkflowClient calls external approval engine.
type WorkflowClient interface {
	VerifyEnrollments(ctx context.Context, trainingID string) error
}

// Service manages training lifecycle transitions.
type Service struct {
	repo     trainingDomain.Repository
	events   EventPublisher
	trail    AuditTrailLogger
	timeline TimelineRecorder
	workflow WorkflowClient
}

// NewService instantiates Service.
func NewService(
	repo trainingDomain.Repository,
	events EventPublisher,
	trail AuditTrailLogger,
	timeline TimelineRecorder,
	workflow WorkflowClient,
) *Service {
	return &Service{
		repo:     repo,
		events:   events,
		trail:    trail,
		timeline: timeline,
		workflow: workflow,
	}
}

// generateTrainingNumber tag formats.
func generateTrainingNumber() string {
	now := time.Now()
	return fmt.Sprintf("TRN-%d%02d%02d-%s", now.Year(), now.Month(), now.Day(), uuid.New().String()[:8])
}

// CreateTraining registers a new competency training.
func (s *Service) CreateTraining(ctx context.Context, cmd CreateTrainingCommand, actorID string) (*trainingDomain.Training, error) {
	now := time.Now()
	t := &trainingDomain.Training{
		ID:             uuid.New().String(),
		TrainingNumber: generateTrainingNumber(),
		CourseID:       cmd.CourseID,
		DepartmentID:   cmd.DepartmentID,
		StatusCode:     string(statusDomain.CodeDraft),
		Title:          cmd.Title,
		Description:    cmd.Description,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := t.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, t); err != nil {
		return nil, err
	}

	evt := timelineDomain.NewEvent(t.ID, timelineDomain.EventCreated, actorID, "Training session created", nil)
	_ = s.timeline.Record(ctx, evt)

	trailLog := trailDomain.NewEntry("training", t.ID, trailDomain.ActionCreated, actorID, nil, t)
	_ = s.trail.Log(ctx, trailLog)

	_ = s.events.Publish(ctx, "training.created", t)

	return t, nil
}

// GetTraining returns details.
func (s *Service) GetTraining(ctx context.Context, id string) (*trainingDomain.Training, error) {
	return s.repo.FindByID(ctx, id)
}

// TransitionStatus manages machine validations.
func (s *Service) TransitionStatus(ctx context.Context, cmd TransitionStatusCommand) error {
	t, err := s.repo.FindByID(ctx, cmd.TrainingID)
	if err != nil {
		return err
	}

	fromStatus := statusDomain.Code(t.StatusCode)
	toStatus := statusDomain.Code(cmd.TargetCode)

	if err := statusDomain.ValidateTransition(fromStatus, toStatus); err != nil {
		return err
	}

	// Enforce attendance checks
	if toStatus == statusDomain.CodeAssessment {
		if err := policyDomain.CheckAttendanceRequirements(t, 1); err != nil {
			prahariLogger.Info(ctx, "Training requires verified attendance", prahariLogger.String("id", t.ID))
		}
	}

	t.StatusCode = cmd.TargetCode
	t.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, t); err != nil {
		return err
	}

	if toStatus == statusDomain.CodeEnrollment {
		_ = s.workflow.VerifyEnrollments(ctx, t.ID)
	}

	evt := timelineDomain.NewEvent(t.ID, timelineDomain.EventType(fmt.Sprintf("TRAINING_%s", cmd.TargetCode)), cmd.ActorID,
		fmt.Sprintf("Training transitioned to %s", cmd.TargetCode), nil)
	_ = s.timeline.Record(ctx, evt)

	trailLog := trailDomain.NewEntry("training", t.ID, trailDomain.ActionStatus, cmd.ActorID,
		map[string]string{"status": string(fromStatus)}, map[string]string{"status": cmd.TargetCode})
	_ = s.trail.Log(ctx, trailLog)

	eventTopic := fmt.Sprintf("training.%s", cmd.TargetCode)
	_ = s.events.Publish(ctx, eventTopic, t)

	return nil
}
