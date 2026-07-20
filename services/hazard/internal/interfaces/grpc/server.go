package grpc

import (
	"context"
	"errors"

	hazardApp "prahari/services/hazard/internal/application/hazard"
	searchApp "prahari/services/hazard/internal/application/search"
	searchDomain "prahari/services/hazard/internal/domain/search"
	hazardDomain "prahari/services/hazard/internal/domain/hazard"
)

// Server exposes gRPC endpoints.
type Server struct {
	hazard *hazardApp.Service
	search *searchApp.Service
}

// NewServer instantiates Server.
func NewServer(
	hazard *hazardApp.Service,
	search *searchApp.Service,
) *Server {
	return &Server{
		hazard: hazard,
		search: search,
	}
}

// CreateHazard inserts profile.
func (s *Server) CreateHazard(ctx context.Context, cmd hazardApp.CreateHazardCommand) (*hazardDomain.Hazard, error) {
	return s.hazard.CreateHazard(ctx, cmd, "grpc-actor")
}

// GetHazard returns hazard profile details.
func (s *Server) GetHazard(ctx context.Context, id string) (*hazardDomain.Hazard, error) {
	if id == "" {
		return nil, errors.New("hazard ID is required")
	}
	return s.hazard.GetHazard(ctx, id)
}

// SearchHazards query matches.
func (s *Server) SearchHazards(ctx context.Context, criteria *searchDomain.Criteria) (*searchDomain.Result, error) {
	return s.search.Search(ctx, criteria)
}
