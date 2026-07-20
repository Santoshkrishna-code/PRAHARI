package grpc

import (
	"context"
	"errors"

	contractorApp "prahari/services/contractor/internal/application/contractor"
	searchApp "prahari/services/contractor/internal/application/search"
	searchDomain "prahari/services/contractor/internal/domain/search"
	contractorDomain "prahari/services/contractor/internal/domain/contractor"
)

// Server exposes gRPC endpoints.
type Server struct {
	contractor *contractorApp.Service
	search     *searchApp.Service
}

// NewServer instantiates Server.
func NewServer(
	contractor *contractorApp.Service,
	search *searchApp.Service,
) *Server {
	return &Server{
		contractor: contractor,
		search:     search,
	}
}

// CreateContractor inserts profile.
func (s *Server) CreateContractor(ctx context.Context, cmd contractorApp.RegisterContractorCommand) (*contractorDomain.Contractor, error) {
	return s.contractor.CreateContractor(ctx, cmd, "grpc-actor")
}

// GetContractor returns contractor profile details.
func (s *Server) GetContractor(ctx context.Context, id string) (*contractorDomain.Contractor, error) {
	if id == "" {
		return nil, errors.New("contractor ID is required")
	}
	return s.contractor.GetContractor(ctx, id)
}

// SearchContractors query matches.
func (s *Server) SearchContractors(ctx context.Context, criteria *searchDomain.Criteria) (*searchDomain.Result, error) {
	return s.search.Search(ctx, criteria)
}
