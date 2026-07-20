package grpc

import (
	"context"
	"net"

	"google.golang.org/grpc"

	consumptionApp "prahari/services/energy/internal/application/consumption"
	forecastingApp "prahari/services/energy/internal/application/forecasting"
	reportingApp "prahari/services/energy/internal/application/reporting"
	"prahari/services/energy/internal/domain/energyprofile"
	"prahari/services/energy/internal/domain/meterreading"
	prahariLogger "prahari/shared/logger"
)

type Server struct {
	sustSvc   *reportingApp.Service
	monSvc    *consumptionApp.Service
	carbonSvc *forecastingApp.Service
}

func NewServer(sustSvc *reportingApp.Service, monSvc *consumptionApp.Service, carbonSvc *forecastingApp.Service) *Server {
	return &Server{
		sustSvc:   sustSvc,
		monSvc:    monSvc,
		carbonSvc: carbonSvc,
	}
}

func (s *Server) Start(ctx context.Context, port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	prahariLogger.Info(ctx, "Starting Energy gRPC Server", prahariLogger.String("port", port))

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			prahariLogger.Error(ctx, "gRPC server shutdown with error", prahariLogger.Err(err))
		}
	}()

	return nil
}

// Mock gRPC contracts implementation

func (s *Server) CreateEnergyProfile(ctx context.Context, req *energyprofile.Profile) (*energyprofile.Profile, error) {
	err := s.sustSvc.CreateProfile(ctx, req)
	return req, err
}

func (s *Server) GetEnergyProfile(ctx context.Context, id string) (*energyprofile.Profile, error) {
	return &energyprofile.Profile{
		ID:           id,
		PlantID:      "plant-3001",
		FacilityName: "Main Boiler Facility",
	}, nil
}

func (s *Server) RecordMeterReading(ctx context.Context, req *meterreading.Reading) (*meterreading.Reading, error) {
	err := s.monSvc.RecordMeterReading(ctx, req)
	return req, err
}
