package recommendation

import (
	"context"
	"fmt"
	"time"

	"prahari/services/ai/internal/domain/events"
	"prahari/services/ai/internal/domain/recommendation"
)

type Repository interface {
	SaveRecommendation(ctx context.Context, r *recommendation.Recommendation) error
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

func (s *Service) GenerateSuggestions(ctx context.Context, plantID, sourceID, recType string) (*recommendation.Recommendation, error) {
	rec := &recommendation.Recommendation{
		ID:        fmt.Sprintf("rec-%d", time.Now().UnixNano()),
		PlantID:   plantID,
		Type:      recType,
		SourceID:  sourceID,
		Content:   "Suggested preventive safety training based on previous historical incident logs.",
		CreatedAt: time.Now(),
	}

	if err := s.repo.SaveRecommendation(ctx, rec); err != nil {
		return nil, err
	}

	_ = s.publisher.Publish(ctx, events.EventRecommendationCreated, rec)
	return rec, nil
}
