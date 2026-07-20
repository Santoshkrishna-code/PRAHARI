package grpc

import (
	"context"
	"net"

	"google.golang.org/grpc"

	clearanceApp "prahari/services/occupational-health/internal/application/clearance"
	searchApp "prahari/services/occupational-health/internal/application/search"
	"prahari/services/occupational-health/internal/domain/healthprofile"
	prahariLogger "prahari/shared/logger"
)

type Server struct {
	clearance *clearanceApp.Service
	searchSvc *searchApp.Service
}

func NewServer(clearance *clearanceApp.Service, searchSvc *searchApp.Service) *Server {
	return &Server{
		clearance: clearance,
		searchSvc: searchSvc,
	}
}

func (s *Server) Start(ctx context.Context, port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	// Mock registration or custom services registration
	prahariLogger.Info(ctx, "Starting Occupational Health gRPC Server", prahariLogger.String("port", port))

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			prahariLogger.Error(ctx, "gRPC server shutdown with error", prahariLogger.Err(err))
		}
	}()

	return nil
}

// Mock gRPC endpoint definitions matching the required contracts.

func (s *Server) CreateHealthProfile(ctx context.Context, req *healthprofile.HealthProfile) (*healthprofile.HealthProfile, error) {
	err := s.clearance.CreateProfile(ctx, req)
	return req, err
}

func (s *Server) GetHealthProfile(ctx context.Context, id string) (*healthprofile.HealthProfile, error) {
	// Returns mock/real detail data from orchestrator
	return &healthprofile.HealthProfile{
		ID:              id,
		WorkerID:        "worker-123",
		WorkerType:      "EMPLOYEE",
		DepartmentID:    "dept-456",
		ClearanceStatus: "CLEARED",
		MedicalStatus:   "ACTIVE_MONITORING",
	}, nil
}
