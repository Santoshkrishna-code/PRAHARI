package search

import (
	"context"

	"prahari/services/visitor/internal/domain/search"
	"prahari/services/visitor/internal/domain/visit"
)

type Repository interface {
	SearchVisits(ctx context.Context, criteria *search.Criteria) ([]*visit.Visit, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*visit.Visit, int64, error) {
	return s.repo.SearchVisits(ctx, criteria)
}
