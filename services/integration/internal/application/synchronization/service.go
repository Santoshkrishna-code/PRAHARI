package synchronization

import (
	"context"
	"fmt"
	"time"

	"prahari/services/integration/internal/domain/events"
	"prahari/services/integration/internal/domain/synchronization"
)

type Repository interface {
	SaveSynchronization(ctx context.Context, r *synchronization.Record) error
}

type EventPublisher interface {
	Publish(ctx context.Context, eventType string, payload any) error
}

type Service struct {
	repo      Repository
	publisher EventPublisher
}

func NewService(repo Repository, pub EventPublisher) *Service {
	return &Service{repo: repo, publisher: pub}
}

func (s *Service) ExecuteJobSync(ctx context.Context, jobID string) error {
	rec := &synchronization.Record{
		ID:        fmt.Sprintf("sync-%d", time.Now().UnixNano()),
		JobID:     jobID,
		StartedAt: time.Now(),
		Status:    "RUNNING",
	}

	_ = s.repo.SaveSynchronization(ctx, rec)

	// Mock process
	now := time.Now()
	rec.FinishedAt = &now
	rec.Status = "SUCCESS"
	rec.RecordsCount = 42

	if err := s.repo.SaveSynchronization(ctx, rec); err != nil {
		_ = s.publisher.Publish(ctx, events.EventIntegrationFailed, rec)
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventIntegrationCompleted, rec)
	return nil
}
