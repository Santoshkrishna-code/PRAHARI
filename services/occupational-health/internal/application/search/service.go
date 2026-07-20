package search

import (
	"context"

	"prahari/services/occupational-health/internal/domain/healthprofile"
	"prahari/services/occupational-health/internal/domain/search"
)

// Repository manages paginated criteria queries.
type Repository interface {
	Search(ctx context.Context, criteria search.Criteria) ([]healthprofile.HealthProfile, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) SearchProfiles(ctx context.Context, query search.Criteria) ([]healthprofile.HealthProfile, error) {
	if query.Limit <= 0 {
		query.Limit = 10
	}
	return s.repo.Search(ctx, query)
}
