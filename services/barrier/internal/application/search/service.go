package search

import (
	"context"

	"prahari/services/barrier/internal/domain/barrier"
	"prahari/services/barrier/internal/domain/search"
)

type Repository interface {
	SearchBarriers(ctx context.Context, criteria *search.Criteria) ([]*barrier.Barrier, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*barrier.Barrier, int64, error) {
	return s.repo.SearchBarriers(ctx, criteria)
}
