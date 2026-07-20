package planning

import (
	"context"
	"fmt"
	"time"

	"prahari/services/loto/internal/domain/events"
	"prahari/services/loto/internal/domain/isolationplan"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SavePlan(ctx context.Context, plan *isolationplan.Plan) error
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

func (s *Service) CreateIsolationPlan(ctx context.Context, plan *isolationplan.Plan) error {
	plan.ID = fmt.Sprintf("pln-%d", time.Now().UnixNano())
	plan.CreatedAt = time.Now()
	plan.UpdatedAt = time.Now()

	if err := s.repo.SavePlan(ctx, plan); err != nil {
		return fmt.Errorf("failed to save isolation plan: %w", err)
	}

	_ = s.publisher.Publish(ctx, events.EventLOTOPlanned, plan)
	prahariLogger.Info(ctx, "LOTO isolation plan registered",
		prahariLogger.String("equipment_id", plan.EquipmentID))
	return nil
}
