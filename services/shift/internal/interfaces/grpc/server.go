package grpc

import (
	"context"

	handoverApp "prahari/services/shift/internal/application/handover"
	reportingApp "prahari/services/shift/internal/application/reporting"
	schedulingApp "prahari/services/shift/internal/application/scheduling"
	"prahari/services/shift/internal/domain/shift"
)

type Server struct {
	schedulingSvc *schedulingApp.Service
	handoverSvc   *handoverApp.Service
	reportingSvc  *reportingApp.Service
}

func NewServer(
	schedulingSvc *schedulingApp.Service,
	handoverSvc *handoverApp.Service,
	reportingSvc *reportingApp.Service,
) *Server {
	return &Server{
		schedulingSvc: schedulingSvc,
		handoverSvc:   handoverSvc,
		reportingSvc:  reportingSvc,
	}
}

func (s *Server) CreateShift(ctx context.Context, sh *shift.Shift) error {
	return s.schedulingSvc.CreateShift(ctx, sh)
}

func (s *Server) StartShift(ctx context.Context, id string) error {
	return s.schedulingSvc.StartShift(ctx, id)
}

func (s *Server) AcceptHandover(ctx context.Context, handoverID string) error {
	return s.handoverSvc.AcceptHandover(ctx, handoverID)
}

func (s *Server) GetShift(ctx context.Context, id string) (*shift.Shift, error) {
	return s.reportingSvc.GetShift(ctx, id)
}
