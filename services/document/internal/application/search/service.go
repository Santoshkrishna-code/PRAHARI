package search

import (
	"context"

	"prahari/services/document/internal/domain/document"
	"prahari/services/document/internal/domain/search"
)

type Repository interface {
	SearchDocuments(ctx context.Context, criteria *search.Criteria) ([]*document.Document, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*document.Document, int64, error) {
	return s.repo.SearchDocuments(ctx, criteria)
}
