package grpc

import (
	"context"
	"errors"

	nearmissApp "prahari/services/nearmiss/internal/application/nearmiss"
	searchApp "prahari/services/nearmiss/internal/application/search"
	nearmissDomain "prahari/services/nearmiss/internal/domain/nearmiss"
	searchDomain "prahari/services/nearmiss/internal/domain/search"
)

// Server exposes gRPC endpoints.
type Server struct {
	nearmiss *nearmissApp.Service
	search   *searchApp.Service
}

// NewServer instantiates Server.
func NewServer(
	nearmiss *nearmissApp.Service,
	search *searchApp.Service,
) *Server {
	return &Server{
		nearmiss: nearmiss,
		search:   search,
	}
}

// CreateNearMiss inserts profile.
func (s *Server) CreateNearMiss(ctx context.Context, cmd nearmissApp.CreateNearMissCommand) (*nearmissDomain.NearMiss, error) {
	return s.nearmiss.CreateNearMiss(ctx, cmd, "grpc-actor")
}

// GetNearMiss returns profile details.
func (s *Server) GetNearMiss(ctx context.Context, id string) (*nearmissDomain.NearMiss, error) {
	if id == "" {
		return nil, errors.New("near miss ID is required")
	}
	return s.nearmiss.GetNearMiss(ctx, id)
}

// SearchNearMisses query matches.
func (s *Server) SearchNearMisses(ctx context.Context, criteria *searchDomain.Criteria) (*searchDomain.Result, error) {
	return s.search.Search(ctx, criteria)
}
