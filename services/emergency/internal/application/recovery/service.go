package recovery

import (
	"context"
	"fmt"
	"time"

	"prahari/services/emergency/internal/domain/events"
	"prahari/services/emergency/internal/domain/recovery"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveRecovery(ctx context.Context, rec *recovery.Plan) error
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

func (s *Service) InitiateRecovery(ctx context.Context, rec *recovery.Plan) error {
	rec.ID = fmt.Sprintf("rec-%d", time.Now().UnixNano())
	rec.Status = "IN_PROGRESS"
	rec.CreatedAt = time.Now()

	if err := s.repo.SaveRecovery(ctx, rec); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventRecoveryStarted, rec)
	prahariLogger.Info(ctx, "Post-emergency site recovery initiated", prahariLogger.String("emergency_id", rec.EmergencyID))
	return nil
}
