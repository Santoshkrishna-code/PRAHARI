package search

import (
	"context"
	"fmt"

	searchDomain "prahari/services/inspection/internal/domain/search"
	inspectionDomain "prahari/services/inspection/internal/domain/inspection"
)

// SearchRepository queries dynamic parameters.
type SearchRepository interface {
	Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*inspectionDomain.Inspection, int, error)
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

	inspections, totalCount, err := s.repo.Search(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("search queries failed: %w", err)
	}

	totalPages := totalCount / criteria.PageSize
	if totalCount%criteria.PageSize > 0 {
		totalPages++
	}

	return &searchDomain.Result{
		Items:      inspections,
		TotalCount: totalCount,
		Page:       criteria.Page,
		PageSize:   criteria.PageSize,
		TotalPages: totalPages,
	}, nil
}
