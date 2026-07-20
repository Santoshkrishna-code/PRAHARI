package grpc

import (
	"context"
	"errors"

	auditApp "prahari/services/audit/internal/application/audit"
	searchApp "prahari/services/audit/internal/application/search"
	auditDomain "prahari/services/audit/internal/domain/audit"
	searchDomain "prahari/services/audit/internal/domain/search"
)

// Server exposes gRPC endpoints.
type Server struct {
	audit  *auditApp.Service
	search *searchApp.Service
}

// NewServer instantiates Server.
func NewServer(
	audit *auditApp.Service,
	search *searchApp.Service,
) *Server {
	return &Server{
		audit:  audit,
		search: search,
	}
}

// CreateAudit registers new entry.
func (s *Server) CreateAudit(ctx context.Context, cmd auditApp.CreateAuditCommand) (*auditDomain.Audit, error) {
	return s.audit.CreateAudit(ctx, cmd, "grpc-actor")
}

// GetAudit returns registers details.
func (s *Server) GetAudit(ctx context.Context, id string) (*auditDomain.Audit, error) {
	if id == "" {
		return nil, errors.New("audit ID is required")
	}
	return s.audit.GetAudit(ctx, id)
}

// SearchAudits queries matches.
func (s *Server) SearchAudits(ctx context.Context, criteria *searchDomain.Criteria) (*searchDomain.Result, error) {
	return s.search.Search(ctx, criteria)
}
