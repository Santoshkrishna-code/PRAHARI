package grpc

import (
	"context"

	attendanceApp "prahari/services/meetings/internal/application/attendance"
	closureApp "prahari/services/meetings/internal/application/closure"
	conductApp "prahari/services/meetings/internal/application/conduct"
	reportingApp "prahari/services/meetings/internal/application/reporting"
	schedulingApp "prahari/services/meetings/internal/application/scheduling"
	attendanceDomain "prahari/services/meetings/internal/domain/attendance"
	"prahari/services/meetings/internal/domain/meeting"
	"prahari/services/meetings/internal/domain/minutes"
)

type Server struct {
	schedulingSvc *schedulingApp.Service
	conductSvc    *conductApp.Service
	attendanceSvc *attendanceApp.Service
	closureSvc    *closureApp.Service
	reportingSvc  *reportingApp.Service
}

func NewServer(
	schedulingSvc *schedulingApp.Service,
	conductSvc *conductApp.Service,
	attendanceSvc *attendanceApp.Service,
	closureSvc *closureApp.Service,
	reportingSvc *reportingApp.Service,
) *Server {
	return &Server{
		schedulingSvc: schedulingSvc,
		conductSvc:    conductSvc,
		attendanceSvc: attendanceSvc,
		closureSvc:    closureSvc,
		reportingSvc:  reportingSvc,
	}
}

func (s *Server) CreateMeeting(ctx context.Context, mtg *meeting.Meeting) error {
	return s.schedulingSvc.ScheduleMeeting(ctx, mtg)
}

func (s *Server) StartMeeting(ctx context.Context, meetingID string) error {
	return s.conductSvc.StartMeeting(ctx, meetingID)
}

func (s *Server) RecordAttendance(ctx context.Context, meetingID string, rec *attendanceDomain.Record) error {
	return s.attendanceSvc.RecordAttendance(ctx, meetingID, rec)
}

func (s *Server) ApproveMinutes(ctx context.Context, meetingID string, m *minutes.Minutes) error {
	return s.closureSvc.ApproveMinutes(ctx, meetingID, m)
}

func (s *Server) CloseMeeting(ctx context.Context, meetingID string) error {
	return s.closureSvc.CloseMeeting(ctx, meetingID)
}

func (s *Server) GetMeeting(ctx context.Context, id string) (*meeting.Meeting, error) {
	return s.reportingSvc.GetMeeting(ctx, id)
}
