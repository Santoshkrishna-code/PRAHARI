package planning

import (
	"context"
	"fmt"
	"time"

	"prahari/services/emergency/internal/domain/responseplan"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SavePlan(ctx context.Context, plan *responseplan.Plan) error
	GetPlanByID(ctx context.Context, id string) (*responseplan.Plan, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateResponsePlan(ctx context.Context, plan *responseplan.Plan) error {
	plan.ID = fmt.Sprintf("plan-%d", time.Now().UnixNano())
	plan.PlanNumber = fmt.Sprintf("ERP-%s-%d", plan.PlantID, time.Now().Unix()%100000)
	plan.CreatedAt = time.Now()

	if err := s.repo.SavePlan(ctx, plan); err != nil {
		return fmt.Errorf("failed to save response plan: %w", err)
	}

	prahariLogger.Info(ctx, "Emergency Response Plan (ERP) created", prahariLogger.String("plan_number", plan.PlanNumber))
	return nil
}
