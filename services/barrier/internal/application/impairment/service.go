package impairment

import (
	"context"
	"fmt"
	"time"

	"prahari/services/barrier/internal/domain/barrier"
	"prahari/services/barrier/internal/domain/bypass"
	"prahari/services/barrier/internal/domain/events"
	"prahari/services/barrier/internal/domain/impairment"
	"prahari/services/barrier/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetBarrierByID(ctx context.Context, id string) (*barrier.Barrier, error)
	SaveBarrier(ctx context.Context, b *barrier.Barrier) error
	SaveImpairment(ctx context.Context, imp *impairment.Record) error
	SaveBypass(ctx context.Context, bp *bypass.Record) error
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

func (s *Service) RegisterImpairment(ctx context.Context, imp *impairment.Record) error {
	b, err := s.repo.GetBarrierByID(ctx, imp.BarrierID)
	if err != nil {
		return err
	}

	imp.ID = fmt.Sprintf("imp-%d", time.Now().UnixNano())
	imp.ImpairedAt = time.Now()
	imp.IsActive = true

	b.Status = string(status.CodeImpaired)
	b.HealthScore = 50.0
	b.UpdatedAt = time.Now()

	if err := s.repo.SaveImpairment(ctx, imp); err != nil {
		return err
	}

	_ = s.repo.SaveBarrier(ctx, b)
	_ = s.publisher.Publish(ctx, events.EventBarrierIntegrityFailed, imp)
	prahariLogger.Warn(ctx, "Barrier impairment registered", prahariLogger.String("barrier_id", imp.BarrierID))
	return nil
}

func (s *Service) RegisterBypass(ctx context.Context, bp *bypass.Record) error {
	b, err := s.repo.GetBarrierByID(ctx, bp.BarrierID)
	if err != nil {
		return err
	}

	bp.ID = fmt.Sprintf("bp-%d", time.Now().UnixNano())
	bp.BypassedAt = time.Now()
	bp.IsActive = true

	b.Status = string(status.CodeBypassed)
	b.UpdatedAt = time.Now()

	if err := s.repo.SaveBypass(ctx, bp); err != nil {
		return err
	}

	_ = s.repo.SaveBarrier(ctx, b)
	_ = s.publisher.Publish(ctx, events.EventBarrierBypassed, bp)
	prahariLogger.Warn(ctx, "Barrier bypass activated", prahariLogger.String("barrier_id", bp.BarrierID))
	return nil
}
