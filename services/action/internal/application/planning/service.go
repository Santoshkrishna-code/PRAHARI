package planning

import (
	"context"
	"fmt"
	"time"

	"prahari/services/action/internal/domain/action"
	"prahari/services/action/internal/domain/events"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
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

func (s *Service) CreateActionItem(ctx context.Context, act *action.Action) error {
	act.ID = fmt.Sprintf("act-%d", time.Now().UnixNano())
	act.Status = "CREATED"
	act.CreatedAt = time.Now()
	act.UpdatedAt = time.Now()

	if err := s.repo.SaveAction(ctx, act); err != nil {
		return fmt.Errorf("failed to save action item: %w", err)
	}

	_ = s.publisher.Publish(ctx, events.EventActionCreated, act)
	prahariLogger.Info(ctx, "Continuous improvement action item created",
		prahariLogger.String("source_module", act.SourceModule),
		prahariLogger.String("source_ref_id", act.SourceRefID))
	return nil
}
