package certification

import (
	"context"

	certDomain "prahari/services/training/internal/domain/certification"
)

// Repository persistence credentials.
type Repository interface {
	Create(ctx context.Context, c *certDomain.Certification) error
}

// Service manages validity logs.
type Service struct {
	repo Repository
}

// NewService instantiates Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// IssueCertification issues credentials.
func (s *Service) IssueCertification(ctx context.Context, c *certDomain.Certification) (*certDomain.Certification, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}
