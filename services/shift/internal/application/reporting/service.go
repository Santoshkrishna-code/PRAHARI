package reporting

import (
	"context"

	"prahari/services/shift/internal/domain/shift"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetShiftByID(ctx context.Context, id string) (*shift.Shift, error)
	ListShifts(ctx context.Context, plantID string) ([]*shift.Shift, error)
	GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetShift(ctx context.Context, id string) (*shift.Shift, error) {
	return s.repo.GetShiftByID(ctx, id)
}

func (s *Service) ListShifts(ctx context.Context, plantID string) ([]*shift.Shift, error) {
	return s.repo.ListShifts(ctx, plantID)
}

func (s *Service) GetExecutiveMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving executive shift dashboard metrics", prahariLogger.String("plant_id", plantID))
	return s.repo.GetDashboardMetrics(ctx, plantID)
}
