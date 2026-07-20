package search

import (
	"context"

	"prahari/services/chemical/internal/domain/chemical"
	"prahari/services/chemical/internal/domain/search"
)

type Repository interface {
	SearchChemicals(ctx context.Context, criteria *search.Criteria) ([]*chemical.Chemical, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*chemical.Chemical, int64, error) {
	return s.repo.SearchChemicals(ctx, criteria)
}
