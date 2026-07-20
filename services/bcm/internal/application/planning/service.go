package planning

import (
	"context"
	"fmt"
	"time"

	"prahari/services/bcm/internal/domain/continuityplan"
	"prahari/services/bcm/internal/domain/events"
	"prahari/services/bcm/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SavePlan(ctx context.Context, plan *continuityplan.Plan) error
	GetPlanByID(ctx context.Context, id string) (*continuityplan.Plan, error)
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

func (s *Service) CreateContinuityPlan(ctx context.Context, plan *continuityplan.Plan) error {
	plan.ID = fmt.Sprintf("bcp-%d", time.Now().UnixNano())
	plan.PlanNumber = fmt.Sprintf("BCP-%s-%d", plan.PlantID, time.Now().Unix()%100000)
	plan.Status = string(status.CodePlanned)
	plan.CreatedAt = time.Now()
	plan.UpdatedAt = time.Now()

	if err := s.repo.SavePlan(ctx, plan); err != nil {
		return fmt.Errorf("failed to save continuity plan: %w", err)
	}

	_ = s.publisher.Publish(ctx, events.EventBCMPlanCreated, plan)
	prahariLogger.Info(ctx, "Business Continuity Plan (BCP) created", prahariLogger.String("plan_number", plan.PlanNumber))
	return nil
}

func (s *Service) ActivatePlan(ctx context.Context, id string) error {
	plan, err := s.repo.GetPlanByID(ctx, id)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(plan.Status), status.CodeActivation); err != nil {
		return err
	}

	plan.Status = string(status.CodeActivation)
	plan.UpdatedAt = time.Now()

	if err := s.repo.SavePlan(ctx, plan); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventContinuityActivated, plan)
	prahariLogger.Warn(ctx, "Business Continuity Plan (BCP) activated!", prahariLogger.String("plan_id", id))
	return nil
}
