package playback

import (
	"context"
	"time"

	"prahari/services/digitaltwin/internal/domain/events"
	"prahari/services/digitaltwin/internal/domain/playback"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SavePlaybackSession(ctx context.Context, ps *playback.Session) error
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

func (s *Service) StartPlayback(ctx context.Context, twinID string, start, end time.Time, speed float64) (*playback.Session, error) {
	ps := &playback.Session{
		ID:        "play-" + time.Now().Format("20060102150405"),
		TwinID:    twinID,
		StartTime: start,
		EndTime:   end,
		Speed:     speed,
		Status:    "PLAYING",
		CreatedAt: time.Now(),
	}

	if err := s.repo.SavePlaybackSession(ctx, ps); err != nil {
		return nil, err
	}

	prahariLogger.Info(ctx, "Started historical operational replay timeline session",
		prahariLogger.String("twin_id", twinID),
		prahariLogger.String("start", start.Format(time.RFC3339)))

	// Mock completion callback
	go func() {
		time.Sleep(150 * time.Millisecond)
		bgCtx := context.Background()
		ps.Status = "STOPPED"
		_ = s.repo.SavePlaybackSession(bgCtx, ps)
		_ = s.publisher.Publish(bgCtx, events.EventTwinPlaybackCompleted, ps)
	}()

	return ps, nil
}
