package consumption

import (
	"context"
	"fmt"
	"time"

	"prahari/services/water/internal/domain/consumption"
	"prahari/services/water/internal/domain/events"
	"prahari/services/water/internal/domain/meterreading"
	"prahari/services/water/internal/domain/policy"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveReading(ctx context.Context, reading *meterreading.Reading) error
	SaveConsumption(ctx context.Context, c *consumption.Consumption) error
	GetLatestReading(ctx context.Context, meterID string) (*meterreading.Reading, error)
}

type EventPublisher interface {
	Publish(ctx context.Context, eventType string, payload any) error
}

type EnvironmentalClient interface {
	SendWaterUsageMetrics(ctx context.Context, plantID string, usageKL float64) error
}

type ESGClient interface {
	LogWaterStewardshipIndex(ctx context.Context, plantID string, recyclePct float64) error
}

type Service struct {
	repo          Repository
	publisher     EventPublisher
	envClient     EnvironmentalClient
	esgClient     ESGClient
}

func NewService(repo Repository, pub EventPublisher, envClient EnvironmentalClient, esgClient ESGClient) *Service {
	return &Service{
		repo:      repo,
		publisher: pub,
		envClient: envClient,
		esgClient: esgClient,
	}
}

func (s *Service) RecordReading(ctx context.Context, reading *meterreading.Reading) error {
	reading.ID = fmt.Sprintf("rdg-%d", time.Now().UnixNano())
	reading.CreatedAt = time.Now()

	if err := s.repo.SaveReading(ctx, reading); err != nil {
		return fmt.Errorf("failed to save meter reading: %w", err)
	}

	_ = s.publisher.Publish(ctx, events.EventFlowmeterReadingRec, reading)
	prahariLogger.Info(ctx, "Flow meter reading recorded", prahariLogger.String("meter_id", reading.MeterID))
	return nil
}

func (s *Service) CalculateConsumption(ctx context.Context, plantID, meterID string, periodStart, periodEnd time.Time, consumptionKL, budgetKL float64) (*consumption.Consumption, error) {
	c := &consumption.Consumption{
		ID:            fmt.Sprintf("con-%d", time.Now().UnixNano()),
		PlantID:       plantID,
		MeterID:       meterID,
		PeriodStart:   periodStart,
		PeriodEnd:     periodEnd,
		ConsumptionKL: consumptionKL,
		CreatedAt:     time.Now(),
	}

	if err := s.repo.SaveConsumption(ctx, c); err != nil {
		return nil, fmt.Errorf("failed to save consumption: %w", err)
	}

	if !policy.EvaluateWaterTarget(c, budgetKL) {
		_ = s.publisher.Publish(ctx, events.EventWaterTargetExceeded, map[string]any{
			"plant_id":       plantID,
			"consumption_kl": consumptionKL,
			"budget_kl":      budgetKL,
		})
	}

	if s.envClient != nil {
		_ = s.envClient.SendWaterUsageMetrics(ctx, plantID, consumptionKL)
	}

	return c, nil
}
