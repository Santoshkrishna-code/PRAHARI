package barrier

import (
	"context"
	"fmt"
	"time"

	"prahari/services/barrier/internal/domain/barrier"
	"prahari/services/barrier/internal/domain/events"
	"prahari/services/barrier/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveBarrier(ctx context.Context, b *barrier.Barrier) error
	GetBarrierByID(ctx context.Context, id string) (*barrier.Barrier, error)
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

func (s *Service) CreateBarrier(ctx context.Context, b *barrier.Barrier) error {
	b.ID = fmt.Sprintf("bar-%d", time.Now().UnixNano())
	b.BarrierCode = fmt.Sprintf("BAR-%s-%d", b.PlantID, time.Now().Unix()%100000)
	b.Status = string(status.CodeRegistered)
	b.HealthScore = 100.0
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()

	if err := s.repo.SaveBarrier(ctx, b); err != nil {
		return fmt.Errorf("failed to save barrier: %w", err)
	}

	_ = s.publisher.Publish(ctx, events.EventBarrierCreated, b)
	prahariLogger.Info(ctx, "Safety barrier registered", prahariLogger.String("barrier_code", b.BarrierCode))
	return nil
}

func (s *Service) GetBarrier(ctx context.Context, id string) (*barrier.Barrier, error) {
	return s.repo.GetBarrierByID(ctx, id)
}
