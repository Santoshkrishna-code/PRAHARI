package grpc

import (
	"context"

	reportingApp "prahari/services/pha/internal/application/reporting"
	studyApp "prahari/services/pha/internal/application/study"
	"prahari/services/pha/internal/domain/phastudy"
)

type Server struct {
	studySvc     *studyApp.Service
	reportingSvc *reportingApp.Service
}

func NewServer(studySvc *studyApp.Service, reportingSvc *reportingApp.Service) *Server {
	return &Server{
		studySvc:     studySvc,
		reportingSvc: reportingSvc,
	}
}

func (s *Server) CreatePHAStudy(ctx context.Context, st *phastudy.Study) error {
	return s.studySvc.CreateStudy(ctx, st)
}

func (s *Server) GetPHAStudy(ctx context.Context, id string) (*phastudy.Study, error) {
	return s.reportingSvc.GetStudy(ctx, id)
}
