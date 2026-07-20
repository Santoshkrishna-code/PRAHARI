package grpc

import (
	"context"
	"net"

	"google.golang.org/grpc"

	environmentApp "prahari/services/environmental/internal/application/environment"
	searchApp "prahari/services/environmental/internal/application/search"
	"prahari/services/environmental/internal/domain/environment"
	prahariLogger "prahari/shared/logger"
)

type Server struct {
	envSvc    *environmentApp.Service
	searchSvc *searchApp.Service
}

func NewServer(envSvc *environmentApp.Service, searchSvc *searchApp.Service) *Server {
	return &Server{
		envSvc:    envSvc,
		searchSvc: searchSvc,
	}
}

func (s *Server) Start(ctx context.Context, port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	prahariLogger.Info(ctx, "Starting Environmental gRPC Server", prahariLogger.String("port", port))

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			prahariLogger.Error(ctx, "gRPC server shutdown with error", prahariLogger.Err(err))
		}
	}()

	return nil
}

// Mock gRPC contracts implementation

func (s *Server) CreateEnvironmentalRecord(ctx context.Context, req *environment.EnvironmentalAspect) (*environment.EnvironmentalAspect, error) {
	err := s.envSvc.RegisterAspect(ctx, req)
	return req, err
}

func (s *Server) GetEnvironmentalRecord(ctx context.Context, id string) (*environment.EnvironmentalAspect, error) {
	return &environment.EnvironmentalAspect{
		ID:             id,
		PlantID:        "plant-101",
		DepartmentID:   "dept-202",
		Name:           "Air Stack Emission Outlet 1",
		AspectCategory: "AIR_EMISSION",
	}, nil
}
