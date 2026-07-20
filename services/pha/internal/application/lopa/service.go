package lopa

import (
	"context"
	"fmt"
	"time"

	"prahari/services/pha/internal/domain/events"
	"prahari/services/pha/internal/domain/lopa"
	"prahari/services/pha/internal/domain/policy"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveLOPA(ctx context.Context, l *lopa.Analysis) error
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

func (s *Service) ExecuteLOPA(ctx context.Context, l *lopa.Analysis) (*lopa.Analysis, error) {
	l.ID = fmt.Sprintf("lopa-%d", time.Now().UnixNano())
	l.CreatedAt = time.Now()

	mitFreq, rrf, sil := policy.CalculateLOPARequiredSIL(l.InitiatingEventFreq, l.TolerableTargetFreq, l.TotalIPLmitigation)
	l.MitigatedEventFreq = mitFreq
	l.RequiredRRF = rrf
	l.TargetSIL = sil

	if err := s.repo.SaveLOPA(ctx, l); err != nil {
		return nil, fmt.Errorf("failed to save LOPA analysis: %w", err)
	}

	_ = s.publisher.Publish(ctx, events.EventLOPACompleted, l)
	prahariLogger.Info(ctx, "LOPA evaluation executed",
		prahariLogger.String("target_sil", l.TargetSIL),
		prahariLogger.String("scenario_id", l.ScenarioID))
	return l, nil
}
