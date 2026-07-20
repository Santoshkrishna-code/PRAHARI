package integrity

import (
	"context"
	"fmt"
	"time"

	"prahari/services/barrier/internal/domain/barrier"
	"prahari/services/barrier/internal/domain/events"
	"prahari/services/barrier/internal/domain/integrityassessment"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetBarrierByID(ctx context.Context, id string) (*barrier.Barrier, error)
	SaveBarrier(ctx context.Context, b *barrier.Barrier) error
	SaveIntegrityAssessment(ctx context.Context, ia *integrityassessment.Assessment) error
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

func (s *Service) AssessIntegrity(ctx context.Context, ia *integrityassessment.Assessment) error {
	b, err := s.repo.GetBarrierByID(ctx, ia.BarrierID)
	if err != nil {
		return err
	}

	ia.ID = fmt.Sprintf("ia-%d", time.Now().UnixNano())
	ia.AssessedAt = time.Now()

	b.HealthScore = ia.HealthScore
	b.UpdatedAt = time.Now()

	if ia.Status == "CRITICAL" || ia.HealthScore < 50.0 {
		_ = s.publisher.Publish(ctx, events.EventBarrierIntegrityFailed, ia)
		prahariLogger.Warn(ctx, "Barrier integrity assessment failed", prahariLogger.String("barrier_id", ia.BarrierID))
	}

	if err := s.repo.SaveIntegrityAssessment(ctx, ia); err != nil {
		return err
	}
	return s.repo.SaveBarrier(ctx, b)
}
