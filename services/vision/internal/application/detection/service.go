package detection

import (
	"context"
	"fmt"
	"time"

	"prahari/services/vision/internal/domain/detection"
	"prahari/services/vision/internal/domain/events"
)

type Repository interface {
	SaveDetection(ctx context.Context, d *detection.Detection) error
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

func (s *Service) RecordDetection(ctx context.Context, jobID, label string, confidence float64, bbox string) (*detection.Detection, error) {
	d := &detection.Detection{
		ID:         fmt.Sprintf("det-%d", time.Now().UnixNano()),
		JobID:      jobID,
		Label:      label,
		Confidence: confidence,
		BBox:       bbox,
		Timestamp:  time.Now(),
	}

	if err := s.repo.SaveDetection(ctx, d); err != nil {
		return nil, err
	}

	_ = s.publisher.Publish(ctx, events.EventVisionDetectionCreated, d)
	return d, nil
}
