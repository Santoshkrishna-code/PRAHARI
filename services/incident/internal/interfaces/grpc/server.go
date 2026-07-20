package grpc

import (
	"context"
	"errors"

	incidentApp "prahari/services/incident/internal/application/incident"
	searchApp "prahari/services/incident/internal/application/search"
	assignmentApp "prahari/services/incident/internal/application/assignment"
	assignmentDomain "prahari/services/incident/internal/domain/assignment"
	searchDomain "prahari/services/incident/internal/domain/search"
	incidentDomain "prahari/services/incident/internal/domain/incident"
)

// Server implements the gRPC service interface for incident operations.
// Used by other platform services to interact with incidents programmatically.
type Server struct {
	incidentSvc   *incidentApp.Service
	assignmentSvc *assignmentApp.Service
	searchSvc     *searchApp.Service
}

// NewServer constructs a gRPC Server.
func NewServer(
	incidentSvc *incidentApp.Service,
	assignmentSvc *assignmentApp.Service,
	searchSvc *searchApp.Service,
) *Server {
	return &Server{
		incidentSvc:   incidentSvc,
		assignmentSvc: assignmentSvc,
		searchSvc:     searchSvc,
	}
}

// CreateIncident handles gRPC incident creation requests.
func (s *Server) CreateIncident(ctx context.Context, cmd incidentApp.CreateIncidentCommand) (*incidentDomain.Incident, error) {
	return s.incidentSvc.CreateIncident(ctx, cmd)
}

// GetIncident handles gRPC incident retrieval requests.
func (s *Server) GetIncident(ctx context.Context, id string) (*incidentDomain.Incident, error) {
	if id == "" {
		return nil, errors.New("incident ID is required")
	}
	return s.incidentSvc.GetIncident(ctx, id)
}

// UpdateIncident handles gRPC incident update requests.
func (s *Server) UpdateIncident(ctx context.Context, id string, cmd incidentApp.UpdateIncidentCommand, actorID string) (*incidentDomain.Incident, error) {
	return s.incidentSvc.UpdateIncident(ctx, id, cmd, actorID)
}

// AssignIncident handles gRPC incident assignment requests.
func (s *Server) AssignIncident(ctx context.Context, incidentID, assigneeID, assignerID string, role string) (*assignmentDomain.Assignment, error) {
	return s.assignmentSvc.AssignIncident(ctx, incidentID, assigneeID, assignerID, assignmentDomain.Role(role))
}

// ResolveIncident handles gRPC incident resolution requests.
func (s *Server) ResolveIncident(ctx context.Context, incidentID, actorID string) error {
	cmd := incidentApp.TransitionStatusCommand{
		IncidentID: incidentID,
		TargetCode: "RESOLVED",
		ActorID:    actorID,
	}
	return s.incidentSvc.TransitionStatus(ctx, cmd)
}

// SearchIncidents handles gRPC incident search requests.
func (s *Server) SearchIncidents(ctx context.Context, criteria *searchDomain.Criteria) (*searchDomain.Result, error) {
	return s.searchSvc.Search(ctx, criteria)
}
