package grpc

import (
	"context"

	barrierApp "prahari/services/barrier/internal/application/barrier"
	prooftestApp "prahari/services/barrier/internal/application/prooftest"
	reportingApp "prahari/services/barrier/internal/application/reporting"
	"prahari/services/barrier/internal/domain/barrier"
	"prahari/services/barrier/internal/domain/prooftest"
)

type Server struct {
	barrierSvc   *barrierApp.Service
	prooftestSvc *prooftestApp.Service
	reportingSvc *reportingApp.Service
}

func NewServer(barrierSvc *barrierApp.Service, prooftestSvc *prooftestApp.Service, reportingSvc *reportingApp.Service) *Server {
	return &Server{
		barrierSvc:   barrierSvc,
		prooftestSvc: prooftestSvc,
		reportingSvc: reportingSvc,
	}
}

func (s *Server) CreateBarrier(ctx context.Context, b *barrier.Barrier) error {
	return s.barrierSvc.CreateBarrier(ctx, b)
}

func (s *Server) GetBarrier(ctx context.Context, id string) (*barrier.Barrier, error) {
	return s.reportingSvc.GetBarrier(ctx, id)
}

func (s *Server) RecordProofTest(ctx context.Context, pt *prooftest.Test) error {
	return s.prooftestSvc.RecordProofTest(ctx, pt)
}
