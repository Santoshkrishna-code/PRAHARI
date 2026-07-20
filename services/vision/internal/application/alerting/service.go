package alerting

import (
	"context"
	"fmt"
	"time"

	"prahari/services/vision/internal/domain/alert"
	"prahari/services/vision/internal/domain/events"
	"prahari/services/vision/internal/domain/policy"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveAlert(ctx context.Context, a *alert.EventTrigger) error
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

func (s *Service) ProcessDetectionAlert(ctx context.Context, cameraID, label string, bx, by, bw, bh float64) error {
	// Restrict zone coordinates check
	if policy.IsWithinZone(bx, by, bw, bh, 100.0, 100.0, 500.0, 500.0) {
		a := &alert.EventTrigger{
			ID:          fmt.Sprintf("al-%d", time.Now().UnixNano()),
			CameraID:    cameraID,
			Label:       label,
			TriggeredAt: time.Now(),
			SnapshotURL: fmt.Sprintf("http://s3.aws/snapshots/%s.jpg", label),
		}

		if err := s.repo.SaveAlert(ctx, a); err != nil {
			return err
		}

		_ = s.publisher.Publish(ctx, events.EventVisionAlertTriggered, a)
		prahariLogger.Error(ctx, "Restricted zone safety intrusion alert triggered by perception layer",
			prahariLogger.String("camera", cameraID),
			prahariLogger.String("label", label))
	}
	return nil
}
