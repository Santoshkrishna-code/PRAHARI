package investigation

import (
	"context"

	investigationDomain "prahari/services/nearmiss/internal/domain/investigation"
)

// Repository persistence investigations definitions.
type Repository interface {
	Create(ctx context.Context, i *investigationDomain.Investigation) error
	FindByNearMissID(ctx context.Context, nearmissID string) ([]*investigationDomain.Investigation, error)
}

// Service manages investigation workflows.
type Service struct {
	repo Repository
}

// NewService instantiates Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateInvestigation persists log data.
func (s *Service) CreateInvestigation(ctx context.Context, i *investigationDomain.Investigation) (*investigationDomain.Investigation, error) {
	if err := i.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, i); err != nil {
		return nil, err
	}
	return i, nil
}
