package search

import (
	"context"
	"fmt"

	searchDomain "prahari/services/permit/internal/domain/search"
	permitDomain "prahari/services/permit/internal/domain/permit"
)

// SearchRepository defines search persistence contracts.
type SearchRepository interface {
	Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*permitDomain.Permit, int, error)
}

// Service manages search queries.
type Service struct {
	repo SearchRepository
}

// NewService instantiates a Search Service.
func NewService(repo SearchRepository) *Service {
	return &Service{repo: repo}
}

// Search executes criteria lookups.
func (s *Service) Search(ctx context.Context, criteria *searchDomain.Criteria) (*searchDomain.Result, error) {
	criteria.Normalize()

	permits, totalCount, err := s.repo.Search(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}

	totalPages := totalCount / criteria.PageSize
	if totalCount%criteria.PageSize > 0 {
		totalPages++
	}

	return &searchDomain.Result{
		Items:      permits,
		TotalCount: totalCount,
		Page:       criteria.Page,
		PageSize:   criteria.PageSize,
		TotalPages: totalPages,
	}, nil
}
