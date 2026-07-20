package reporting

import (
	"context"

	"prahari/services/pha/internal/domain/phastudy"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetStudyByID(ctx context.Context, id string) (*phastudy.Study, error)
	ListStudies(ctx context.Context, plantID string) ([]*phastudy.Study, error)
	GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetStudy(ctx context.Context, id string) (*phastudy.Study, error) {
	return s.repo.GetStudyByID(ctx, id)
}

func (s *Service) ListStudies(ctx context.Context, plantID string) ([]*phastudy.Study, error) {
	return s.repo.ListStudies(ctx, plantID)
}

func (s *Service) GetExecutiveMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving executive PHA dashboard metrics", prahariLogger.String("plant_id", plantID))
	return s.repo.GetDashboardMetrics(ctx, plantID)
}
