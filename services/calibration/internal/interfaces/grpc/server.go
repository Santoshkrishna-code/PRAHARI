package grpc

import (
	"context"

	certificateApp "prahari/services/calibration/internal/application/certificate"
	executionApp "prahari/services/calibration/internal/application/execution"
	reportingApp "prahari/services/calibration/internal/application/reporting"
	schedulingApp "prahari/services/calibration/internal/application/scheduling"
	"prahari/services/calibration/internal/domain/calibration"
	"prahari/services/calibration/internal/domain/calibrationschedule"
)

type Server struct {
	schedulingSvc  *schedulingApp.Service
	executionSvc   *executionApp.Service
	certificateSvc *certificateApp.Service
	reportingSvc   *reportingApp.Service
}

func NewServer(
	schedulingSvc *schedulingApp.Service,
	executionSvc *executionApp.Service,
	certificateSvc *certificateApp.Service,
	reportingSvc *reportingApp.Service,
) *Server {
	return &Server{
		schedulingSvc:  schedulingSvc,
		executionSvc:   executionSvc,
		certificateSvc: certificateSvc,
		reportingSvc:   reportingSvc,
	}
}

func (s *Server) CreateCalibration(ctx context.Context, rec *calibration.Record) error {
	return s.executionSvc.StartCalibration(ctx, rec)
}

func (s *Server) ScheduleCalibration(ctx context.Context, sched *calibrationschedule.Schedule) error {
	return s.schedulingSvc.ScheduleCalibrationTask(ctx, sched)
}

func (s *Server) ExecuteCalibration(ctx context.Context, rec *calibration.Record) error {
	return s.executionSvc.StartCalibration(ctx, rec)
}

func (s *Server) ApproveCalibration(ctx context.Context, calID, supervisorID string) error {
	return s.executionSvc.ApproveCalibration(ctx, calID, supervisorID)
}

func (s *Server) GetCalibration(ctx context.Context, id string) (*calibration.Record, error) {
	return s.reportingSvc.GetCalibration(ctx, id)
}
