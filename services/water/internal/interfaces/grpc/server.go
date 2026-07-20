package grpc

import (
	"context"

	reportingApp "prahari/services/water/internal/application/reporting"
	"prahari/services/water/internal/domain/waterprofile"
)

type Server struct {
	reportingSvc *reportingApp.Service
}

func NewServer(reportingSvc *reportingApp.Service) *Server {
	return &Server{reportingSvc: reportingSvc}
}

func (s *Server) CreateWaterProfile(ctx context.Context, profile *waterprofile.Profile) error {
	return s.reportingSvc.CreateProfile(ctx, profile)
}

func (s *Server) GetWaterProfile(ctx context.Context, id string) (*waterprofile.Profile, error) {
	return s.reportingSvc.GetProfile(ctx, id)
}
