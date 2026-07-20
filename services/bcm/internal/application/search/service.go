package search

import (
	"context"

	"prahari/services/bcm/internal/domain/continuityplan"
	"prahari/services/bcm/internal/domain/search"
)

type Repository interface {
	SearchPlans(ctx context.Context, criteria *search.Criteria) ([]*continuityplan.Plan, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*continuityplan.Plan, int64, error) {
	return s.repo.SearchPlans(ctx, criteria)
}
