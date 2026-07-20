package reporting

import (
	"context"

	"prahari/services/emergency/internal/domain/emergency"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetEmergencyByID(ctx context.Context, id string) (*emergency.Emergency, error)
	ListEmergencies(ctx context.Context, plantID string) ([]*emergency.Emergency, error)
	GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetEmergency(ctx context.Context, id string) (*emergency.Emergency, error) {
	return s.repo.GetEmergencyByID(ctx, id)
}

func (s *Service) ListEmergencies(ctx context.Context, plantID string) ([]*emergency.Emergency, error) {
	return s.repo.ListEmergencies(ctx, plantID)
}

func (s *Service) GetExecutiveMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving executive emergency dashboard metrics", prahariLogger.String("plant_id", plantID))
	return s.repo.GetDashboardMetrics(ctx, plantID)
}
