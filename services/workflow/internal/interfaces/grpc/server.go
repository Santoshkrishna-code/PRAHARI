package grpc

import (
	"context"
	"errors"

	prahariApproval "prahari/services/workflow/internal/application/approval"
)

// Server implements gRPC workflow engines control endpoints contracts.
type Server struct {
	approvalSvc *prahariApproval.Service
}

// NewServer constructs a Server.
func NewServer(svc *prahariApproval.Service) *Server {
	return &Server{approvalSvc: svc}
}

// CompleteTask complete execution human tasks steps.
func (s *Server) CompleteTask(ctx context.Context, taskID, approverID, decision string) error {
	if taskID == "" {
		return errors.New("task ID is required")
	}

	return s.approvalSvc.SubmitDecision(ctx, taskID, approverID, decision)
}
