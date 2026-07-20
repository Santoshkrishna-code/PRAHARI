package attendance

import (
	"context"
	"fmt"
	"time"

	attendanceDomain "prahari/services/meetings/internal/domain/attendance"
	"prahari/services/meetings/internal/domain/events"
	"prahari/services/meetings/internal/domain/meeting"
	"prahari/services/meetings/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetMeetingByID(ctx context.Context, id string) (*meeting.Meeting, error)
	SaveMeeting(ctx context.Context, mtg *meeting.Meeting) error
	SaveAttendance(ctx context.Context, rec *attendanceDomain.Record) error
	ListAttendance(ctx context.Context, meetingID string) ([]*attendanceDomain.Record, error)
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

func (s *Service) RecordAttendance(ctx context.Context, meetingID string, rec *attendanceDomain.Record) error {
	mtg, err := s.repo.GetMeetingByID(ctx, meetingID)
	if err != nil {
		return err
	}

	if mtg.Status != string(status.CodeInProgress) && mtg.Status != string(status.CodeAttendanceRecorded) {
		return fmt.Errorf("cannot record attendance when meeting is in state %s", mtg.Status)
	}

	rec.ID = fmt.Sprintf("att-%d", time.Now().UnixNano())
	rec.MeetingID = meetingID
	rec.CheckInAt = time.Now()
	rec.Verified = true

	if err := s.repo.SaveAttendance(ctx, rec); err != nil {
		return err
	}

	// Transition to ATTENDANCE_RECORDED on first attendance record
	if mtg.Status == string(status.CodeInProgress) {
		mtg.Status = string(status.CodeAttendanceRecorded)
		mtg.UpdatedAt = time.Now()
		if err := s.repo.SaveMeeting(ctx, mtg); err != nil {
			return err
		}
	}

	_ = s.publisher.Publish(ctx, events.EventAttendanceRecorded, rec)
	prahariLogger.Info(ctx, "Attendance recorded",
		prahariLogger.String("meeting_id", meetingID),
		prahariLogger.String("attendee_id", rec.AttendeeID))
	return nil
}

func (s *Service) ListAttendance(ctx context.Context, meetingID string) ([]*attendanceDomain.Record, error) {
	return s.repo.ListAttendance(ctx, meetingID)
}
