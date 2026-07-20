package lifecycle

import (
	"context"

	timelineDomain "prahari/services/asset/internal/domain/timeline"
)

// Repository query milestones.
type Repository interface {
	FindByAssetID(ctx context.Context, assetID string) ([]*timelineDomain.Event, error)
}

// Service lists history trace.
type Service struct {
	repo Repository
}

// NewService instantiates Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// GetHistory returns milestone lists.
func (s *Service) GetHistory(ctx context.Context, assetID string) ([]*timelineDomain.Event, error) {
	return s.repo.FindByAssetID(ctx, assetID)
}
