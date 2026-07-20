package summarization

import (
	"context"
	"fmt"
	"time"

	"prahari/services/ai/internal/domain/events"
	"prahari/services/ai/internal/domain/summarization"
)

type Repository interface {
	SaveSummary(ctx context.Context, s *summarization.Summary) error
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

func (s *Service) SummarizeDocument(ctx context.Context, sourceID, original string) (*summarization.Summary, error) {
	sum := &summarization.Summary{
		ID:        fmt.Sprintf("sum-%d", time.Now().UnixNano()),
		SourceID:  sourceID,
		Original:  original,
		Condensed: "Condensed summary containing highlighted findings.",
		CreatedAt: time.Now(),
	}

	if err := s.repo.SaveSummary(ctx, sum); err != nil {
		return nil, err
	}

	_ = s.publisher.Publish(ctx, events.EventSummaryGenerated, sum)
	return sum, nil
}
