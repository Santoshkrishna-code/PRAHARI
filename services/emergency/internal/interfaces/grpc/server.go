package grpc

import (
	"context"

	reportingApp "prahari/services/emergency/internal/application/reporting"
	responseApp "prahari/services/emergency/internal/application/response"
	"prahari/services/emergency/internal/domain/emergency"
)

type Server struct {
	responseSvc  *responseApp.Service
	reportingSvc *reportingApp.Service
}

func NewServer(responseSvc *responseApp.Service, reportingSvc *reportingApp.Service) *Server {
	return &Server{
		responseSvc:  responseSvc,
		reportingSvc: reportingSvc,
	}
}

func (s *Server) CreateEmergency(ctx context.Context, em *emergency.Emergency) error {
	return s.responseSvc.DeclareEmergency(ctx, em)
}

func (s *Server) GetEmergency(ctx context.Context, id string) (*emergency.Emergency, error) {
	return s.reportingSvc.GetEmergency(ctx, id)
}
