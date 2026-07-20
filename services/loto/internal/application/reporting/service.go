package reporting

import (
	"context"

	"prahari/services/loto/internal/domain/isolationcertificate"
	"prahari/services/loto/internal/domain/isolationplan"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetPlanByID(ctx context.Context, id string) (*isolationplan.Plan, error)
	GetCertificateByID(ctx context.Context, id string) (*isolationcertificate.Certificate, error)
	ListCertificates(ctx context.Context, plantID string) ([]*isolationcertificate.Certificate, error)
	GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetPlan(ctx context.Context, id string) (*isolationplan.Plan, error) {
	return s.repo.GetPlanByID(ctx, id)
}

func (s *Service) GetCertificate(ctx context.Context, id string) (*isolationcertificate.Certificate, error) {
	return s.repo.GetCertificateByID(ctx, id)
}

func (s *Service) ListCertificates(ctx context.Context, plantID string) ([]*isolationcertificate.Certificate, error) {
	return s.repo.ListCertificates(ctx, plantID)
}

func (s *Service) GetExecutiveMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving LOTO execution compliance metrics and metrics", prahariLogger.String("plant_id", plantID))
	return s.repo.GetDashboardMetrics(ctx, plantID)
}
