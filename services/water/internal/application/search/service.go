package search

import (
	"context"

	"prahari/services/water/internal/domain/search"
	"prahari/services/water/internal/domain/waterprofile"
)

type Repository interface {
	SearchProfiles(ctx context.Context, criteria *search.Criteria) ([]*waterprofile.Profile, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*waterprofile.Profile, int64, error) {
	return s.repo.SearchProfiles(ctx, criteria)
}
