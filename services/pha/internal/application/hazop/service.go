package hazop

import (
	"context"
	"fmt"
	"time"

	"prahari/services/pha/internal/domain/events"
	"prahari/services/pha/internal/domain/hazardscenario"
	"prahari/services/pha/internal/domain/hazop"
	"prahari/services/pha/internal/domain/policy"
	"prahari/services/pha/internal/domain/processnode"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveNode(ctx context.Context, node *processnode.Node) error
	SaveScenario(ctx context.Context, sc *hazardscenario.Scenario) error
	SaveSession(ctx context.Context, sess *hazop.Session) error
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

func (s *Service) CreateNode(ctx context.Context, node *processnode.Node) error {
	node.ID = fmt.Sprintf("node-%d", time.Now().UnixNano())
	node.CreatedAt = time.Now()
	return s.repo.SaveNode(ctx, node)
}

func (s *Service) RecordScenario(ctx context.Context, sc *hazardscenario.Scenario) error {
	sc.ID = fmt.Sprintf("scn-%d", time.Now().UnixNano())
	sc.CreatedAt = time.Now()

	rank, cat := policy.CalculateRiskRank(sc.Severity, sc.Likelihood)
	sc.RiskRank = rank
	sc.RiskCategory = cat

	if sc.RiskRank >= 10 {
		sc.IsSILRelevant = true
	}

	if err := s.repo.SaveScenario(ctx, sc); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventHAZOPCompleted, sc)
	prahariLogger.Info(ctx, "HAZOP hazard scenario recorded", prahariLogger.String("risk_category", sc.RiskCategory))
	return nil
}
