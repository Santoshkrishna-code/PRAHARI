package reporting

import (
	"context"

	"prahari/services/document/internal/domain/document"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetDocumentByID(ctx context.Context, id string) (*document.Document, error)
	ListDocuments(ctx context.Context, plantID string) ([]*document.Document, error)
	GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetDocument(ctx context.Context, id string) (*document.Document, error) {
	return s.repo.GetDocumentByID(ctx, id)
}

func (s *Service) ListDocuments(ctx context.Context, plantID string) ([]*document.Document, error) {
	return s.repo.ListDocuments(ctx, plantID)
}

func (s *Service) GetExecutiveMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving executive Document Management dashboard metrics", prahariLogger.String("plant_id", plantID))
	return s.repo.GetDashboardMetrics(ctx, plantID)
}
