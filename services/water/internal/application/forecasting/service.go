package forecasting

import (
	"context"
	"fmt"
	"time"

	"prahari/services/water/internal/domain/events"
	"prahari/services/water/internal/domain/forecasting"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveForecast(ctx context.Context, fc *forecasting.Forecast) error
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

func (s *Service) GenerateForecast(ctx context.Context, plantID, period string, baselineKL float64) (*forecasting.Forecast, error) {
	fc := &forecasting.Forecast{
		ID:             fmt.Sprintf("fct-%d", time.Now().UnixNano()),
		PlantID:        plantID,
		ForecastPeriod: period,
		PredictedKL:    baselineKL * 1.05, // 5% growth projection default
		ConfidenceRate: 92.5,
		SeasonalFactor: 1.10,
		GeneratedAt:    time.Now(),
	}

	if err := s.repo.SaveForecast(ctx, fc); err != nil {
		return nil, fmt.Errorf("failed to save water forecast: %w", err)
	}

	_ = s.publisher.Publish(ctx, events.EventWaterForecastGen, fc)
	prahariLogger.Info(ctx, "Water demand forecast generated", prahariLogger.String("plant_id", plantID))
	return fc, nil
}
