package grpc

import (
	"context"
	"errors"

	inspectionApp "prahari/services/inspection/internal/application/inspection"
	searchApp "prahari/services/inspection/internal/application/search"
	searchDomain "prahari/services/inspection/internal/domain/search"
	inspectionDomain "prahari/services/inspection/internal/domain/inspection"
)

// Server exposes walkthrough inspection APIs.
type Server struct {
	inspection *inspectionApp.Service
	search     *searchApp.Service
}

// NewServer instantiates Server.
func NewServer(
	inspection *inspectionApp.Service,
	search *searchApp.Service,
) *Server {
	return &Server{
		inspection: inspection,
		search:     search,
	}
}

// CreateInspection inserts inspections.
func (s *Server) CreateInspection(ctx context.Context, cmd inspectionApp.CreateInspectionCommand) (*inspectionDomain.Inspection, error) {
	return s.inspection.CreateInspection(ctx, cmd)
}

// GetInspection returns inspection aggregate.
func (s *Server) GetInspection(ctx context.Context, id string) (*inspectionDomain.Inspection, error) {
	if id == "" {
		return nil, errors.New("inspection ID is required")
	}
	return s.inspection.GetInspection(ctx, id)
}

// UpdateInspection saves edits.
func (s *Server) UpdateInspection(ctx context.Context, id string, cmd inspectionApp.UpdateInspectionCommand, actor string) (*inspectionDomain.Inspection, error) {
	return s.inspection.UpdateInspection(ctx, id, cmd, actor)
}

// AssignInspection transitions status code.
func (s *Server) AssignInspection(ctx context.Context, id, actor string) error {
	cmd := inspectionApp.TransitionStatusCommand{
		InspectionID: id,
		TargetCode:   "ASSIGNED",
		ActorID:      actor,
	}
	return s.inspection.TransitionStatus(ctx, cmd)
}

// CompleteInspection transitions status to completed.
func (s *Server) CompleteInspection(ctx context.Context, id, actor string) error {
	cmd := inspectionApp.TransitionStatusCommand{
		InspectionID: id,
		TargetCode:   "COMPLETED",
		ActorID:      actor,
	}
	return s.inspection.TransitionStatus(ctx, cmd)
}

// ApproveInspection transitions status to approved.
func (s *Server) ApproveInspection(ctx context.Context, id, actor string) error {
	cmd := inspectionApp.TransitionStatusCommand{
		InspectionID: id,
		TargetCode:   "APPROVED",
		ActorID:      actor,
	}
	return s.inspection.TransitionStatus(ctx, cmd)
}

// SearchInspections lists matches.
func (s *Server) SearchInspections(ctx context.Context, criteria *searchDomain.Criteria) (*searchDomain.Result, error) {
	return s.search.Search(ctx, criteria)
}
