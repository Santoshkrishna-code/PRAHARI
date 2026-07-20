package grpc

import (
	"context"
	"errors"

	complianceApp "prahari/services/compliance/internal/application/compliance"
	searchApp "prahari/services/compliance/internal/application/search"
	complianceDomain "prahari/services/compliance/internal/domain/compliance"
	searchDomain "prahari/services/compliance/internal/domain/search"
)

// Server exposes gRPC endpoints.
type Server struct {
	compliance *complianceApp.Service
	search     *searchApp.Service
}

// NewServer instantiates Server.
func NewServer(
	compliance *complianceApp.Service,
	search *searchApp.Service,
) *Server {
	return &Server{
		compliance: compliance,
		search:     search,
	}
}

// CreateCompliance registers new entry.
func (s *Server) CreateCompliance(ctx context.Context, cmd complianceApp.CreateComplianceCommand) (*complianceDomain.Compliance, error) {
	return s.compliance.CreateCompliance(ctx, cmd, "grpc-actor")
}

// GetCompliance returns registers details.
func (s *Server) GetCompliance(ctx context.Context, id string) (*complianceDomain.Compliance, error) {
	if id == "" {
		return nil, errors.New("compliance ID is required")
	}
	return s.compliance.GetCompliance(ctx, id)
}

// SearchCompliance queries matches.
func (s *Server) SearchCompliance(ctx context.Context, criteria *searchDomain.Criteria) (*searchDomain.Result, error) {
	return s.search.Search(ctx, criteria)
}
