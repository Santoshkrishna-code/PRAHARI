package search

import (
	"context"

	"prahari/services/shift/internal/domain/search"
	"prahari/services/shift/internal/domain/shift"
)

type Repository interface {
	SearchShifts(ctx context.Context, criteria *search.Criteria) ([]*shift.Shift, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*shift.Shift, int64, error) {
	return s.repo.SearchShifts(ctx, criteria)
}
