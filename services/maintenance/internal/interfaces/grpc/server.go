package grpc

import (
	"context"
	"errors"

	maintenanceApp "prahari/services/maintenance/internal/application/maintenance"
	searchApp "prahari/services/maintenance/internal/application/search"
	searchDomain "prahari/services/maintenance/internal/domain/search"
	maintenanceDomain "prahari/services/maintenance/internal/domain/maintenance"
)

// Server exposes gRPC endpoints.
type Server struct {
	maintenance *maintenanceApp.Service
	search      *searchApp.Service
}

// NewServer instantiates Server.
func NewServer(
	maintenance *maintenanceApp.Service,
	search *searchApp.Service,
) *Server {
	return &Server{
		maintenance: maintenance,
		search:      search,
	}
}

// CreateMaintenance inserts profile.
func (s *Server) CreateMaintenance(ctx context.Context, cmd maintenanceApp.CreateMaintenanceCommand) (*maintenanceDomain.Maintenance, error) {
	return s.maintenance.CreateMaintenance(ctx, cmd, "grpc-actor")
}

// GetMaintenance returns maintenance profile details.
func (s *Server) GetMaintenance(ctx context.Context, id string) (*maintenanceDomain.Maintenance, error) {
	if id == "" {
		return nil, errors.New("maintenance ID is required")
	}
	return s.maintenance.GetMaintenance(ctx, id)
}

// AssignMaintenance transitions status.
func (s *Server) AssignMaintenance(ctx context.Context, id, actor string) error {
	cmd := maintenanceApp.TransitionStatusCommand{
		MaintenanceID: id,
		TargetCode:    "ASSIGNED",
		ActorID:       actor,
	}
	return s.maintenance.TransitionStatus(ctx, cmd)
}

// CompleteMaintenance transitions status to completed.
func (s *Server) CompleteMaintenance(ctx context.Context, id, actor string) error {
	cmd := maintenanceApp.TransitionStatusCommand{
		MaintenanceID: id,
		TargetCode:    "COMPLETED",
		ActorID:       actor,
	}
	return s.maintenance.TransitionStatus(ctx, cmd)
}

// VerifyMaintenance transitions status to verified.
func (s *Server) VerifyMaintenance(ctx context.Context, id, actor string) error {
	cmd := maintenanceApp.TransitionStatusCommand{
		MaintenanceID: id,
		TargetCode:    "VERIFIED",
		ActorID:       actor,
	}
	return s.maintenance.TransitionStatus(ctx, cmd)
}

// SearchMaintenance query matches.
func (s *Server) SearchMaintenance(ctx context.Context, criteria *searchDomain.Criteria) (*searchDomain.Result, error) {
	return s.search.Search(ctx, criteria)
}
