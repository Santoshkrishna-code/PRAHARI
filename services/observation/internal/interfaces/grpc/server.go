package grpc

import (
	"context"
	"errors"

	observationApp "prahari/services/observation/internal/application/observation"
	searchApp "prahari/services/observation/internal/application/search"
	observationDomain "prahari/services/observation/internal/domain/observation"
	searchDomain "prahari/services/observation/internal/domain/search"
)

// Server exposes gRPC endpoints.
type Server struct {
	observation *observationApp.Service
	search      *searchApp.Service
}

// NewServer instantiates Server.
func NewServer(
	observation *observationApp.Service,
	search *searchApp.Service,
) *Server {
	return &Server{
		observation: observation,
		search:      search,
	}
}

// CreateObservation inserts profile.
func (s *Server) CreateObservation(ctx context.Context, cmd observationApp.CreateObservationCommand) (*observationDomain.Observation, error) {
	return s.observation.CreateObservation(ctx, cmd, "grpc-actor")
}

// GetObservation returns profile details.
func (s *Server) GetObservation(ctx context.Context, id string) (*observationDomain.Observation, error) {
	if id == "" {
		return nil, errors.New("observation ID is required")
	}
	return s.observation.GetObservation(ctx, id)
}

// SearchObservations query matches.
func (s *Server) SearchObservations(ctx context.Context, criteria *searchDomain.Criteria) (*searchDomain.Result, error) {
	return s.search.Search(ctx, criteria)
}
