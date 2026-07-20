package grpc

import (
	"context"
	"net"

	"google.golang.org/grpc"

	reportingApp "prahari/services/esg/internal/application/reporting"
	sustainabilityApp "prahari/services/esg/internal/application/sustainability"
	"prahari/services/esg/internal/domain/esgobjective"
	prahariLogger "prahari/shared/logger"
)

type Server struct {
	sustSvc *sustainabilityApp.Service
	monSvc  *reportingApp.Service
}

func NewServer(sustSvc *sustainabilityApp.Service, monSvc *reportingApp.Service) *Server {
	return &Server{
		sustSvc: sustSvc,
		monSvc:  monSvc,
	}
}

func (s *Server) Start(ctx context.Context, port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	prahariLogger.Info(ctx, "Starting ESG gRPC Server", prahariLogger.String("port", port))

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			prahariLogger.Error(ctx, "gRPC server shutdown with error", prahariLogger.Err(err))
		}
	}()

	return nil
}

// Mock gRPC contracts implementation

func (s *Server) CreateSustainabilityObjective(ctx context.Context, req *esgobjective.Objective) (*esgobjective.Objective, error) {
	err := s.sustSvc.CreateObjective(ctx, req)
	return req, err
}

func (s *Server) GetSustainabilityObjective(ctx context.Context, id string) (*esgobjective.Objective, error) {
	return &esgobjective.Objective{
		ID:             id,
		BusinessUnitID: "bu-4001",
		Title:          "Reduce carbon intensity rating",
		Category:       "ENVIRONMENTAL",
	}, nil
}
