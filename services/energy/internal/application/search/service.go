package search

import (
	"context"

	"prahari/services/energy/internal/domain/energyprofile"
	"prahari/services/energy/internal/domain/search"
)

type Repository interface {
	SearchProfiles(ctx context.Context, criteria search.Criteria) ([]energyprofile.Profile, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Search(ctx context.Context, criteria search.Criteria) ([]energyprofile.Profile, error) {
	if criteria.Limit == 0 {
		criteria.Limit = 100
	}
	return s.repo.SearchProfiles(ctx, criteria)
}
