package search

import (
	"context"

	"prahari/services/emergency/internal/domain/emergency"
	"prahari/services/emergency/internal/domain/search"
)

type Repository interface {
	SearchEmergencies(ctx context.Context, criteria *search.Criteria) ([]*emergency.Emergency, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*emergency.Emergency, int64, error) {
	return s.repo.SearchEmergencies(ctx, criteria)
}
