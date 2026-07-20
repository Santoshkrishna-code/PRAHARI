package verification

import (
	"context"
	"time"

	"prahari/services/action/internal/domain/action"
	"prahari/services/action/internal/domain/effectivenessreview"
	"prahari/services/action/internal/domain/events"
	"prahari/services/action/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetActionByID(ctx context.Context, id string) (*action.Action, error)
	SaveAction(ctx context.Context, act *action.Action) error
	SaveEffectivenessReview(ctx context.Context, r *effectivenessreview.Review) error
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

func (s *Service) ReviewActionItem(ctx context.Context, actionID string, r *effectivenessreview.Review) error {
	act, err := s.repo.GetActionByID(ctx, actionID)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(act.Status), status.CodeEffectivenessReview); err != nil {
		return err
	}

	act.Status = string(status.CodeEffectivenessReview)
	act.UpdatedAt = time.Now()

	r.ReviewedAt = time.Now()

	if err := s.repo.SaveEffectivenessReview(ctx, r); err != nil {
		return err
	}
	return s.repo.SaveAction(ctx, act)
}

func (s *Service) CloseActionItem(ctx context.Context, actionID string) error {
	act, err := s.repo.GetActionByID(ctx, actionID)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(act.Status), status.CodeClosed); err != nil {
		return err
	}

	now := time.Now()
	act.Status = string(status.CodeClosed)
	act.ClosedAt = &now
	act.UpdatedAt = now

	if err := s.repo.SaveAction(ctx, act); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventActionClosed, act)
	prahariLogger.Info(ctx, "Action item verified and closed", prahariLogger.String("action_id", actionID))
	return nil
}
