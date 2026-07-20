package reporting

import (
	"context"

	"prahari/services/visitor/internal/domain/visit"
	"prahari/services/visitor/internal/domain/visitor"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetVisitorByID(ctx context.Context, id string) (*visitor.Visitor, error)
	GetVisitByID(ctx context.Context, id string) (*visit.Visit, error)
	ListVisits(ctx context.Context, plantID string) ([]*visit.Visit, error)
	GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetVisitor(ctx context.Context, id string) (*visitor.Visitor, error) {
	return s.repo.GetVisitorByID(ctx, id)
}

func (s *Service) GetVisit(ctx context.Context, id string) (*visit.Visit, error) {
	return s.repo.GetVisitByID(ctx, id)
}

func (s *Service) ListVisits(ctx context.Context, plantID string) ([]*visit.Visit, error) {
	return s.repo.ListVisits(ctx, plantID)
}

func (s *Service) GetExecutiveMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving visitor access dashboard metrics", prahariLogger.String("plant_id", plantID))
	return s.repo.GetDashboardMetrics(ctx, plantID)
}
