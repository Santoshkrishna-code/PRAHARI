package conduct

import (
	"context"
	"time"

	"prahari/services/meetings/internal/domain/events"
	"prahari/services/meetings/internal/domain/meeting"
	"prahari/services/meetings/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetMeetingByID(ctx context.Context, id string) (*meeting.Meeting, error)
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

func (s *Service) StartMeeting(ctx context.Context, meetingID string) error {
	mtg, err := s.repo.GetMeetingByID(ctx, meetingID)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(mtg.Status), status.CodeInProgress); err != nil {
		return err
	}

	now := time.Now()
	mtg.Status = string(status.CodeInProgress)
	mtg.StartedAt = &now
	mtg.UpdatedAt = now

	if err := s.repo.SaveMeeting(ctx, mtg); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventMeetingStarted, mtg)
	prahariLogger.Info(ctx, "Meeting started", prahariLogger.String("meeting_id", meetingID))
	return nil
}
