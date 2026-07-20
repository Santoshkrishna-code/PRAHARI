package search

import (
	"context"

	"prahari/services/action/internal/domain/action"
	"prahari/services/action/internal/domain/search"
)

type Repository interface {
	SearchActions(ctx context.Context, criteria *search.Criteria) ([]*action.Action, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*action.Action, int64, error) {
	return s.repo.SearchActions(ctx, criteria)
}
