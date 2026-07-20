package prediction

import (
	"context"
	"fmt"
	"time"

	"prahari/services/ai/internal/domain/events"
	"prahari/services/ai/internal/domain/prediction"
)

type Repository interface {
	SaveForecast(ctx context.Context, f *prediction.Forecast) error
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

func (s *Service) PredictRisk(ctx context.Context, plantID, topic string) (*prediction.Forecast, error) {
	f := &prediction.Forecast{
		ID:          fmt.Sprintf("pred-%d", time.Now().UnixNano()),
		PlantID:     plantID,
		TargetTopic: topic,
		Probability: 0.12,
		CreatedAt:   time.Now(),
	}

	if err := s.repo.SaveForecast(ctx, f); err != nil {
		return nil, err
	}

	_ = s.publisher.Publish(ctx, events.EventPredictionCompleted, f)
	return f, nil
}
