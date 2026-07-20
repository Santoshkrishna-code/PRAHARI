package grpc

import (
	"context"
	"errors"

	assessmentApp "prahari/services/risk/internal/application/assessment"
	searchApp "prahari/services/risk/internal/application/search"
	riskDomain "prahari/services/risk/internal/domain/risk"
	searchDomain "prahari/services/risk/internal/domain/search"
)

// Server exposes gRPC endpoints.
type Server struct {
	assessment *assessmentApp.Service
	search     *searchApp.Service
}

// NewServer instantiates Server.
func NewServer(
	assessment *assessmentApp.Service,
	search *searchApp.Service,
) *Server {
	return &Server{
		assessment: assessment,
		search:     search,
	}
}

// CreateRiskAssessment registers new entry.
func (s *Server) CreateRiskAssessment(ctx context.Context, cmd assessmentApp.CreateRiskCommand) (*riskDomain.Risk, error) {
	return s.assessment.CreateRiskAssessment(ctx, cmd, "grpc-actor")
}

// GetRiskAssessment returns registers details.
func (s *Server) GetRiskAssessment(ctx context.Context, id string) (*riskDomain.Risk, error) {
	if id == "" {
		return nil, errors.New("risk ID is required")
	}
	return s.assessment.GetRiskAssessment(ctx, id)
}

// SearchRiskAssessments queries matches.
func (s *Server) SearchRiskAssessments(ctx context.Context, criteria *searchDomain.Criteria) (*searchDomain.Result, error) {
	return s.search.Search(ctx, criteria)
}
