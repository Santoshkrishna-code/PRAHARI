package search

import (
	"context"

	"prahari/services/environmental/internal/domain/environment"
	"prahari/services/environmental/internal/domain/search"
)

type Repository interface {
	SearchAspects(ctx context.Context, criteria search.Criteria) ([]environment.EnvironmentalAspect, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Search(ctx context.Context, criteria search.Criteria) ([]environment.EnvironmentalAspect, error) {
	if criteria.Limit == 0 {
		criteria.Limit = 100
	}
	return s.repo.SearchAspects(ctx, criteria)
}
