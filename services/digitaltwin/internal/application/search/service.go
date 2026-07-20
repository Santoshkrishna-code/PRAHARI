package search

import (
	"context"

	"prahari/services/digitaltwin/internal/domain/search"
	"prahari/services/digitaltwin/internal/domain/twin"
)

type Repository interface {
	SearchTwins(ctx context.Context, criteria *search.Criteria) ([]*twin.DigitalTwin, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*twin.DigitalTwin, int64, error) {
	return s.repo.SearchTwins(ctx, criteria)
}
