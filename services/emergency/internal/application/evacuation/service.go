package evacuation

import (
	"context"
	"fmt"
	"time"

	"prahari/services/emergency/internal/domain/evacuation"
	"prahari/services/emergency/internal/domain/events"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveEvacuation(ctx context.Context, rec *evacuation.Record) error
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

func (s *Service) InitiateEvacuation(ctx context.Context, rec *evacuation.Record) error {
	rec.ID = fmt.Sprintf("evac-%d", time.Now().UnixNano())
	rec.Status = "IN_PROGRESS"
	rec.InitiatedAt = time.Now()

	if err := s.repo.SaveEvacuation(ctx, rec); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventEvacuationStarted, rec)
	prahariLogger.Warn(ctx, "Evacuation initiated for plant zone", prahariLogger.String("zone_id", rec.ZoneID))
	return nil
}
