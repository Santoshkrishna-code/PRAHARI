package execution

import (
	"context"
	"time"

	"prahari/services/action/internal/domain/action"
	"prahari/services/action/internal/domain/events"
	"prahari/services/action/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetActionByID(ctx context.Context, id string) (*action.Action, error)
	SaveAction(ctx context.Context, act *action.Action) error
}

type EventPublisher interface {
	Publish(ctx context.Context, eventType string, payload any) error
}

type Service struct {
	repo      Repository
	publisher EventPublisher
}

func NewService(repo Repository, pub EventPublisher) *Service {
	return &Service{
		repo:      repo,
		publisher: pub,
	}
}

func (s *Service) AssignActionItem(ctx context.Context, actionID, userID string) error {
	act, err := s.repo.GetActionByID(ctx, actionID)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(act.Status), status.CodeAssigned); err != nil {
		return err
	}

	act.Status = string(status.CodeAssigned)
	act.AssignedTo = userID
	act.UpdatedAt = time.Now()

	if err := s.repo.SaveAction(ctx, act); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventActionAssigned, act)
	prahariLogger.Info(ctx, "Action item assigned", prahariLogger.String("action_id", actionID), prahariLogger.String("assigned_to", userID))
	return nil
}

func (s *Service) StartActionInProgress(ctx context.Context, actionID string) error {
	act, err := s.repo.GetActionByID(ctx, actionID)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(act.Status), status.CodeInProgress); err != nil {
		return err
	}

	act.Status = string(status.CodeInProgress)
	act.UpdatedAt = time.Now()

	return s.repo.SaveAction(ctx, act)
}

func (s *Service) SubmitEvidence(ctx context.Context, actionID string) error {
	act, err := s.repo.GetActionByID(ctx, actionID)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(act.Status), status.CodeEvidenceSubmitted); err != nil {
		return err
	}

	act.Status = string(status.CodeEvidenceSubmitted)
	act.UpdatedAt = time.Now()

	return s.repo.SaveAction(ctx, act)
}
