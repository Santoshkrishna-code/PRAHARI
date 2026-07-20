package search

import (
	"context"
	"fmt"

	assetDomain "prahari/services/asset/internal/domain/asset"
	searchDomain "prahari/services/asset/internal/domain/search"
)

// SearchRepository queries dynamic parameters.
type SearchRepository interface {
	Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*assetDomain.Asset, int, error)
}

// Service executes search filters.
type Service struct {
	repo SearchRepository
}

// NewService instantiates Search Service.
func NewService(repo SearchRepository) *Service {
	return &Service{repo: repo}
}

// Search retrieves page metrics.
func (s *Service) Search(ctx context.Context, criteria *searchDomain.Criteria) (*searchDomain.Result, error) {
	criteria.Normalize()

	assets, totalCount, err := s.repo.Search(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("search query failure: %w", err)
	}

	totalPages := totalCount / criteria.PageSize
	if totalCount%criteria.PageSize > 0 {
		totalPages++
	}

	return &searchDomain.Result{
		Items:      assets,
		TotalCount: totalCount,
		Page:       criteria.Page,
		PageSize:   criteria.PageSize,
		TotalPages: totalPages,
	}, nil
}
