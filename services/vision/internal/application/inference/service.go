package inference

import (
	"context"
	"fmt"
	"time"

	"prahari/services/vision/internal/domain/events"
	"prahari/services/vision/internal/domain/inference"
)

type Repository interface {
	SaveInferenceJob(ctx context.Context, job *inference.Job) error
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

func (s *Service) StartJob(ctx context.Context, cameraID, modelID string) (*inference.Job, error) {
	job := &inference.Job{
		ID:        fmt.Sprintf("job-%d", time.Now().UnixNano()),
		CameraID:  cameraID,
		ModelID:   modelID,
		Status:    "RUNNING",
		FPSRate:   30.0,
		StartedAt: time.Now(),
	}

	if err := s.repo.SaveInferenceJob(ctx, job); err != nil {
		return nil, err
	}

	return job, nil
}

func (s *Service) StopJob(ctx context.Context, job *inference.Job) error {
	job.Status = "STOPPED"
	if err := s.repo.SaveInferenceJob(ctx, job); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventVisionInferenceCompleted, job)
	return nil
}
