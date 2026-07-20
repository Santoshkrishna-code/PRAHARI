package recovery

import (
	"context"
	"time"

	"prahari/services/bcm/internal/domain/continuityplan"
	"prahari/services/bcm/internal/domain/events"
	"prahari/services/bcm/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)


type Repository interface {
	GetPlanByID(ctx context.Context, id string) (*continuityplan.Plan, error)
	SavePlan(ctx context.Context, plan *continuityplan.Plan) error
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

func (s *Service) CompleteRecovery(ctx context.Context, planID string) error {
	plan, err := s.repo.GetPlanByID(ctx, planID)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(plan.Status), status.CodeRecovery); err != nil {
		return err
	}

	plan.Status = string(status.CodeRecovery)
	plan.UpdatedAt = time.Now()

	if err := s.repo.SavePlan(ctx, plan); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventRecoveryStarted, plan)
	prahariLogger.Info(ctx, "Business continuity recovery phase initiated", prahariLogger.String("plan_id", planID))
	return nil
}
