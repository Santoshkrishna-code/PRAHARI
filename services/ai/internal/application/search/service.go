package search

import (
	"context"

	"prahari/services/ai/internal/domain/document"
	"prahari/services/ai/internal/domain/search"
)

type Repository interface {
	SearchDocuments(ctx context.Context, criteria *search.Criteria) ([]*document.Doc, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*document.Doc, int64, error) {
	return s.repo.SearchDocuments(ctx, criteria)
}
