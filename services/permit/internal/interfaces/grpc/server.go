package grpc

import (
	"context"
	"errors"

	permitApp "prahari/services/permit/internal/application/permit"
	approvalApp "prahari/services/permit/internal/application/approval"
	searchApp "prahari/services/permit/internal/application/search"
	approvalDomain "prahari/services/permit/internal/domain/approval"
	searchDomain "prahari/services/permit/internal/domain/search"
	permitDomain "prahari/services/permit/internal/domain/permit"
)

// Server implements inter-service permit actions.
type Server struct {
	permit   *permitApp.Service
	approval *approvalApp.Service
	search   *searchApp.Service
}

// NewServer instantiates Server.
func NewServer(
	permit *permitApp.Service,
	approval *approvalApp.Service,
	search *searchApp.Service,
) *Server {
	return &Server{
		permit:   permit,
		approval: approval,
		search:   search,
	}
}

// CreatePermit triggers draft insertions.
func (s *Server) CreatePermit(ctx context.Context, cmd permitApp.CreatePermitCommand) (*permitDomain.Permit, error) {
	return s.permit.CreatePermit(ctx, cmd)
}

// GetPermit retrieves a permit.
func (s *Server) GetPermit(ctx context.Context, id string) (*permitDomain.Permit, error) {
	if id == "" {
		return nil, errors.New("permit ID is required")
	}
	return s.permit.GetPermit(ctx, id)
}

// UpdatePermit registers mutations.
func (s *Server) UpdatePermit(ctx context.Context, id string, cmd permitApp.UpdatePermitCommand, actor string) (*permitDomain.Permit, error) {
	return s.permit.UpdatePermit(ctx, id, cmd, actor)
}

// ApprovePermit signs stages.
func (s *Server) ApprovePermit(ctx context.Context, id, approver string, role string, signature string) (*approvalDomain.Approval, error) {
	return s.approval.SubmitApproval(ctx, id, approver, approvalDomain.Role(role), approvalDomain.DecisionApproved, "approved via gRPC", signature)
}

// IssuePermit transitions permit state.
func (s *Server) IssuePermit(ctx context.Context, id, actor string) error {
	cmd := permitApp.TransitionStatusCommand{
		PermitID:   id,
		TargetCode: "ISSUED",
		ActorID:    actor,
	}
	return s.permit.TransitionStatus(ctx, cmd)
}

// CompletePermit transitions status to completed.
func (s *Server) CompletePermit(ctx context.Context, id, actor string) error {
	cmd := permitApp.TransitionStatusCommand{
		PermitID:   id,
		TargetCode: "COMPLETED",
		ActorID:    actor,
	}
	return s.permit.TransitionStatus(ctx, cmd)
}

// SearchPermits executes dynamic list queries.
func (s *Server) SearchPermits(ctx context.Context, criteria *searchDomain.Criteria) (*searchDomain.Result, error) {
	return s.search.Search(ctx, criteria)
}
