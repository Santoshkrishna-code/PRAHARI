package grpc

import (
	"context"

	executionApp "prahari/services/action/internal/application/execution"
	planningApp "prahari/services/action/internal/application/planning"
	reportingApp "prahari/services/action/internal/application/reporting"
	verificationApp "prahari/services/action/internal/application/verification"
	"prahari/services/action/internal/domain/action"
	"prahari/services/action/internal/domain/effectivenessreview"
)

type Server struct {
	planningSvc     *planningApp.Service
	executionSvc    *executionApp.Service
	verificationSvc *verificationApp.Service
	reportingSvc    *reportingApp.Service
}

func NewServer(
	planningSvc *planningApp.Service,
	executionSvc *executionApp.Service,
	verificationSvc *verificationApp.Service,
	reportingSvc *reportingApp.Service,
) *Server {
	return &Server{
		planningSvc:     planningSvc,
		executionSvc:    executionSvc,
		verificationSvc: verificationSvc,
		reportingSvc:    reportingSvc,
	}
}

func (s *Server) CreateAction(ctx context.Context, act *action.Action) error {
	return s.planningSvc.CreateActionItem(ctx, act)
}

func (s *Server) AssignAction(ctx context.Context, actionID, userID string) error {
	return s.executionSvc.AssignActionItem(ctx, actionID, userID)
}

func (s *Server) SubmitEvidence(ctx context.Context, actionID string) error {
	return s.executionSvc.SubmitEvidence(ctx, actionID)
}

func (s *Server) ReviewAction(ctx context.Context, actionID string, r *effectivenessreview.Review) error {
	return s.verificationSvc.ReviewActionItem(ctx, actionID, r)
}

func (s *Server) CloseAction(ctx context.Context, actionID string) error {
	return s.verificationSvc.CloseActionItem(ctx, actionID)
}

func (s *Server) GetAction(ctx context.Context, id string) (*action.Action, error) {
	return s.reportingSvc.GetAction(ctx, id)
}
