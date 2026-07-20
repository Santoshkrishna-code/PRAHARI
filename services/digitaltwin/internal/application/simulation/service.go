package simulation

import (
	"context"
	"errors"
	"time"

	"prahari/services/digitaltwin/internal/domain/events"
	"prahari/services/digitaltwin/internal/domain/policy"
	"prahari/services/digitaltwin/internal/domain/simulation"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveScenario(ctx context.Context, sc *simulation.Scenario) error
}

type EventPublisher interface {
	Publish(ctx context.Context, eventType string, payload any) error
}

type Service struct {
	repo      Repository
	publisher EventPublisher
}

func NewService(repo Repository, publisher EventPublisher) *Service {
	return &Service{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *Service) RunScenario(ctx context.Context, twinID, name, params string) (*simulation.Scenario, error) {
	if !policy.ValidateSimulationParameters(params) {
		return nil, errors.New("invalid or oversized simulation parameter payloads")
	}

	sc := &simulation.Scenario{
		ID:         "sim-" + time.Now().Format("20060102150405"),
		TwinID:     twinID,
		Name:       name,
		Status:     "RUNNING",
		Parameters: params,
		StartedAt:  time.Now(),
	}

	if err := s.repo.SaveScenario(ctx, sc); err != nil {
		return nil, err
	}

	prahariLogger.Info(ctx, "Triggered physical scenario predictive simulation model",
		prahariLogger.String("twin_id", twinID),
		prahariLogger.String("scenario", name))

	// Mock async execution
	go func() {
		time.Sleep(100 * time.Millisecond)
		bgCtx := context.Background()
		sc.Status = "COMPLETED"
		sc.ResultData = `{"impact_factor": "HIGH", "cascade_failure_risk": 0.12, "recommended_action": "isolate_valve_v102"}`
		sc.CompletedAt = time.Now()
		_ = s.repo.SaveScenario(bgCtx, sc)
		_ = s.publisher.Publish(bgCtx, events.EventTwinSimulationCompleted, sc)
	}()

	return sc, nil
}
