package search

import (
	"context"
	"fmt"

	searchDomain "prahari/services/incident/internal/domain/search"
	incidentDomain "prahari/services/incident/internal/domain/incident"
)

// SearchRepository defines the persistence port for search operations.
type SearchRepository interface {
	Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*incidentDomain.Incident, int, error)
}

// Service orchestrates search operations with filter normalization and pagination.
type Service struct {
	repo SearchRepository
}

// NewService constructs a Service with the search repository injected.
func NewService(repo SearchRepository) *Service {
	return &Service{repo: repo}
}

// Search executes a parameterized search across incidents using the provided criteria.
// It normalizes pagination defaults before delegating to the repository.
func (s *Service) Search(ctx context.Context, criteria *searchDomain.Criteria) (*searchDomain.Result, error) {
	criteria.Normalize()

	incidents, totalCount, err := s.repo.Search(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("search execution failed: %w", err)
	}

	totalPages := totalCount / criteria.PageSize
	if totalCount%criteria.PageSize > 0 {
		totalPages++
	}

	return &searchDomain.Result{
		Items:      incidents,
		TotalCount: totalCount,
		Page:       criteria.Page,
		PageSize:   criteria.PageSize,
		TotalPages: totalPages,
	}, nil
}
