package distribution

import (
	"context"
	"fmt"

	"prahari/services/water/internal/domain/distributionnetwork"
	"prahari/services/water/internal/domain/leakdetection"
	"prahari/services/water/internal/domain/pipeline"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveNetwork(ctx context.Context, net *distributionnetwork.Network) error
	SavePipeline(ctx context.Context, pipe *pipeline.Pipeline) error
	SaveLeak(ctx context.Context, leak *leakdetection.Leak) error
}

type MaintenanceClient interface {
	CreateWorkOrder(ctx context.Context, plantID, title, description string) (string, error)
}

type Service struct {
	repo        Repository
	maintClient MaintenanceClient
}

func NewService(repo Repository, maintClient MaintenanceClient) *Service {
	return &Service{
		repo:        repo,
		maintClient: maintClient,
	}
}

func (s *Service) ReportLeak(ctx context.Context, leak *leakdetection.Leak) error {
	if s.maintClient != nil {
		woID, err := s.maintClient.CreateWorkOrder(ctx, leak.PlantID, fmt.Sprintf("Water Leak Repair: Zone %s", leak.ZoneCode), leak.LocationDesc)
		if err == nil {
			leak.WorkOrderID = woID
		}
	}

	if err := s.repo.SaveLeak(ctx, leak); err != nil {
		return fmt.Errorf("failed to save leak record: %w", err)
	}

	prahariLogger.Warn(ctx, "Water leak reported", prahariLogger.String("zone", leak.ZoneCode), prahariLogger.String("severity", leak.Severity))
	return nil
}
