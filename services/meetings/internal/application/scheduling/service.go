package scheduling

import (
	"context"
	"fmt"
	"time"

	"prahari/services/meetings/internal/domain/events"
	"prahari/services/meetings/internal/domain/meeting"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveMeeting(ctx context.Context, mtg *meeting.Meeting) error
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

func (s *Service) ScheduleMeeting(ctx context.Context, mtg *meeting.Meeting) error {
	mtg.ID = fmt.Sprintf("mtg-%d", time.Now().UnixNano())
	mtg.Status = "PLANNED"
	mtg.CreatedAt = time.Now()
	mtg.UpdatedAt = time.Now()

	if err := s.repo.SaveMeeting(ctx, mtg); err != nil {
		return fmt.Errorf("failed to save meeting: %w", err)
	}

	_ = s.publisher.Publish(ctx, events.EventMeetingCreated, mtg)
	prahariLogger.Info(ctx, "Meeting scheduled",
		prahariLogger.String("meeting_type", mtg.MeetingType),
		prahariLogger.String("title", mtg.Title))
	return nil
}
