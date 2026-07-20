package search

import (
	"context"

	"prahari/services/pha/internal/domain/phastudy"
	"prahari/services/pha/internal/domain/search"
)

type Repository interface {
	SearchStudies(ctx context.Context, criteria *search.Criteria) ([]*phastudy.Study, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*phastudy.Study, int64, error) {
	return s.repo.SearchStudies(ctx, criteria)
}
