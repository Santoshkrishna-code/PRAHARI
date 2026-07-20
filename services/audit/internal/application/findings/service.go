package findings

import (
	"context"

	findingDomain "prahari/services/audit/internal/domain/finding"
)

// Repository persistence findings gap.
type Repository interface {
	Create(ctx context.Context, f *findingDomain.Finding) error
}

// Service tracks observations and NCR corrective CAPA tasks.
type Service struct {
	repo Repository
}

// NewService instantiates Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// LogFinding registers gap findings.
func (s *Service) LogFinding(ctx context.Context, f *findingDomain.Finding) (*findingDomain.Finding, error) {
	if err := f.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, f); err != nil {
		return nil, err
	}
	return f, nil
}
