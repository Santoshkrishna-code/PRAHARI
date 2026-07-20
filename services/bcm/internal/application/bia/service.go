package bia

import (
	"context"
	"fmt"
	"time"

	"prahari/services/bcm/internal/domain/businessimpactanalysis"
	"prahari/services/bcm/internal/domain/events"
	"prahari/services/bcm/internal/domain/policy"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveBIA(ctx context.Context, bia *businessimpactanalysis.Analysis) error
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

func (s *Service) ExecuteBIA(ctx context.Context, bia *businessimpactanalysis.Analysis) (*businessimpactanalysis.Analysis, error) {
	bia.ID = fmt.Sprintf("bia-%d", time.Now().UnixNano())
	bia.EvaluatedAt = time.Now()
	bia.CreatedAt = time.Now()

	tier := policy.DeterminePriorityTier(bia.MaximumTolerableDowntimeHrs)
	bia.OperationalImpact = tier

	if err := s.repo.SaveBIA(ctx, bia); err != nil {
		return nil, fmt.Errorf("failed to save Business Impact Analysis: %w", err)
	}

	_ = s.publisher.Publish(ctx, events.EventBIACompleted, bia)
	prahariLogger.Info(ctx, "Business Impact Analysis (BIA) completed",
		prahariLogger.String("process_id", bia.ProcessID),
		prahariLogger.String("priority_tier", tier))
	return bia, nil
}
