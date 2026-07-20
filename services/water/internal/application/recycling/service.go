package recycling

import (
	"context"
	"fmt"
	"time"

	"prahari/services/water/internal/domain/events"
	"prahari/services/water/internal/domain/recycling"
	"prahari/services/water/internal/domain/reuse"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveRecyclingProgram(ctx context.Context, prog *recycling.Program) error
	SaveReuseProgram(ctx context.Context, prog *reuse.Program) error
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

func (s *Service) RegisterRecyclingProgram(ctx context.Context, prog *recycling.Program) error {
	prog.ID = fmt.Sprintf("rec-%d", time.Now().UnixNano())
	prog.CreatedAt = time.Now()
	prog.UpdatedAt = time.Now()

	if err := s.repo.SaveRecyclingProgram(ctx, prog); err != nil {
		return fmt.Errorf("failed to save recycling program: %w", err)
	}

	_ = s.publisher.Publish(ctx, events.EventWaterRecycled, prog)
	prahariLogger.Info(ctx, "Water recycling program registered", prahariLogger.String("program_id", prog.ID))
	return nil
}
