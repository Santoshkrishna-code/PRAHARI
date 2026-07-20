package closure

import (
	"context"
	"time"

	"prahari/services/meetings/internal/domain/events"
	"prahari/services/meetings/internal/domain/meeting"
	"prahari/services/meetings/internal/domain/minutes"
	"prahari/services/meetings/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetMeetingByID(ctx context.Context, id string) (*meeting.Meeting, error)
	SaveMeeting(ctx context.Context, mtg *meeting.Meeting) error
	SaveMinutes(ctx context.Context, m *minutes.Minutes) error
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

func (s *Service) ApproveMinutes(ctx context.Context, meetingID string, m *minutes.Minutes) error {
	mtg, err := s.repo.GetMeetingByID(ctx, meetingID)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(mtg.Status), status.CodeMinutesApproved); err != nil {
		return err
	}

	now := time.Now()
	m.MeetingID = meetingID
	m.Status = "APPROVED"
	m.ApprovedAt = &now
	m.UpdatedAt = now

	if err := s.repo.SaveMinutes(ctx, m); err != nil {
		return err
	}

	mtg.Status = string(status.CodeMinutesApproved)
	mtg.UpdatedAt = now
	return s.repo.SaveMeeting(ctx, mtg)
}

func (s *Service) CloseMeeting(ctx context.Context, meetingID string) error {
	mtg, err := s.repo.GetMeetingByID(ctx, meetingID)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(mtg.Status), status.CodeClosed); err != nil {
		return err
	}

	now := time.Now()
	mtg.Status = string(status.CodeClosed)
	mtg.EndedAt = &now
	mtg.UpdatedAt = now

	if err := s.repo.SaveMeeting(ctx, mtg); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventMeetingClosed, mtg)
	prahariLogger.Info(ctx, "Meeting closed", prahariLogger.String("meeting_id", meetingID))
	return nil
}
