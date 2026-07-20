package grpc

import (
	"context"

	biaApp "prahari/services/bcm/internal/application/bia"
	planningApp "prahari/services/bcm/internal/application/planning"
	recoveryApp "prahari/services/bcm/internal/application/recovery"
	reportingApp "prahari/services/bcm/internal/application/reporting"
	"prahari/services/bcm/internal/domain/businessimpactanalysis"
	"prahari/services/bcm/internal/domain/continuityplan"
)

type Server struct {
	biaSvc       *biaApp.Service
	planningSvc  *planningApp.Service
	recoverySvc  *recoveryApp.Service
	reportingSvc *reportingApp.Service
}

func NewServer(biaSvc *biaApp.Service, planningSvc *planningApp.Service, recoverySvc *recoveryApp.Service, reportingSvc *reportingApp.Service) *Server {
	return &Server{
		biaSvc:       biaSvc,
		planningSvc:  planningSvc,
		recoverySvc:  recoverySvc,
		reportingSvc: reportingSvc,
	}
}

func (s *Server) CreateContinuityPlan(ctx context.Context, plan *continuityplan.Plan) error {
	return s.planningSvc.CreateContinuityPlan(ctx, plan)
}

func (s *Server) GetContinuityPlan(ctx context.Context, id string) (*continuityplan.Plan, error) {
	return s.reportingSvc.GetPlan(ctx, id)
}

func (s *Server) ExecuteBusinessImpactAnalysis(ctx context.Context, bia *businessimpactanalysis.Analysis) (*businessimpactanalysis.Analysis, error) {
	return s.biaSvc.ExecuteBIA(ctx, bia)
}

func (s *Server) ActivateContinuityPlan(ctx context.Context, id string) error {
	return s.planningSvc.ActivatePlan(ctx, id)
}

func (s *Server) CompleteRecovery(ctx context.Context, id string) error {
	return s.recoverySvc.CompleteRecovery(ctx, id)
}
