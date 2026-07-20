package search

import (
	"context"

	"prahari/services/ppe/internal/domain/ppe"
	"prahari/services/ppe/internal/domain/search"
)

type Repository interface {
	SearchPPEs(ctx context.Context, criteria *search.Criteria) ([]*ppe.PPE, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*ppe.PPE, int64, error) {
	return s.repo.SearchPPEs(ctx, criteria)
}
