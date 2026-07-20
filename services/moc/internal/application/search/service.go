package search

import (
	"context"

	"prahari/services/moc/internal/domain/changerequest"
	"prahari/services/moc/internal/domain/search"
)

type Repository interface {
	SearchRequests(ctx context.Context, criteria *search.Criteria) ([]*changerequest.Request, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*changerequest.Request, int64, error) {
	return s.repo.SearchRequests(ctx, criteria)
}
