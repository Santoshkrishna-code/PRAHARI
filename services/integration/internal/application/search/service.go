package search

import (
	"context"

	"prahari/services/integration/internal/domain/connector"
	"prahari/services/integration/internal/domain/search"
)

type Repository interface {
	SearchConnectors(ctx context.Context, criteria *search.Criteria) ([]*connector.Connector, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*connector.Connector, int64, error) {
	return s.repo.SearchConnectors(ctx, criteria)
}
