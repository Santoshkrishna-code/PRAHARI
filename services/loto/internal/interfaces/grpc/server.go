package grpc

import (
	"context"

	executionApp "prahari/services/loto/internal/application/execution"
	planningApp "prahari/services/loto/internal/application/planning"
	reportingApp "prahari/services/loto/internal/application/reporting"
	restorationApp "prahari/services/loto/internal/application/restoration"
	verificationApp "prahari/services/loto/internal/application/verification"
	"prahari/services/loto/internal/domain/isolationcertificate"
	"prahari/services/loto/internal/domain/isolationplan"
	"prahari/services/loto/internal/domain/restoration"
	"prahari/services/loto/internal/domain/verification"
)

type Server struct {
	planningSvc     *planningApp.Service
	executionSvc    *executionApp.Service
	verificationSvc *verificationApp.Service
	restorationSvc  *restorationApp.Service
	reportingSvc    *reportingApp.Service
}

func NewServer(
	planningSvc *planningApp.Service,
	executionSvc *executionApp.Service,
	verificationSvc *verificationApp.Service,
	restorationSvc *restorationApp.Service,
	reportingSvc *reportingApp.Service,
) *Server {
	return &Server{
		planningSvc:     planningSvc,
		executionSvc:    executionSvc,
		verificationSvc: verificationSvc,
		restorationSvc:  restorationSvc,
		reportingSvc:    reportingSvc,
	}
}

func (s *Server) CreateIsolationPlan(ctx context.Context, plan *isolationplan.Plan) error {
	return s.planningSvc.CreateIsolationPlan(ctx, plan)
}

func (s *Server) ApproveIsolation(ctx context.Context, cert *isolationcertificate.Certificate) error {
	return s.executionSvc.ApproveIsolation(ctx, cert)
}

func (s *Server) ApplyLocks(ctx context.Context, certID string) error {
	return s.executionSvc.ApplyLocksAndTags(ctx, certID)
}

func (s *Server) VerifyZeroEnergy(ctx context.Context, certID string, v *verification.ZeroEnergy) error {
	return s.verificationSvc.VerifyZeroEnergy(ctx, certID, v)
}

func (s *Server) RestoreSystem(ctx context.Context, certID string, r *restoration.Record) error {
	return s.restorationSvc.RestoreSystem(ctx, certID, r)
}

func (s *Server) GetLOTO(ctx context.Context, id string) (*isolationcertificate.Certificate, error) {
	return s.reportingSvc.GetCertificate(ctx, id)
}
