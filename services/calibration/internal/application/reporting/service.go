package reporting

import (
	"context"

	"prahari/services/calibration/internal/domain/calibration"
	"prahari/services/calibration/internal/domain/instrument"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetInstrumentByID(ctx context.Context, id string) (*instrument.Instrument, error)
	GetCalibrationByID(ctx context.Context, id string) (*calibration.Record, error)
	ListCalibrations(ctx context.Context, plantID string) ([]*calibration.Record, error)
	GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetInstrument(ctx context.Context, id string) (*instrument.Instrument, error) {
	return s.repo.GetInstrumentByID(ctx, id)
}

func (s *Service) GetCalibration(ctx context.Context, id string) (*calibration.Record, error) {
	return s.repo.GetCalibrationByID(ctx, id)
}

func (s *Service) ListCalibrations(ctx context.Context, plantID string) ([]*calibration.Record, error) {
	return s.repo.ListCalibrations(ctx, plantID)
}

func (s *Service) GetExecutiveMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving instrument calibration status and metrology accuracy index dashboard metrics", prahariLogger.String("plant_id", plantID))
	return s.repo.GetDashboardMetrics(ctx, plantID)
}
