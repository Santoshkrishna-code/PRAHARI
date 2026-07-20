package approval

import (
	"context"

	approvalDomain "prahari/services/risk/internal/domain/approval"
)

// Repository persistence approvals signatures.
type Repository interface {
	Create(ctx context.Context, a *approvalDomain.Approval) error
	FindByRiskID(ctx context.Context, riskID string) ([]*approvalDomain.Approval, error)
}

// Service manages signatures logic.
type Service struct {
	repo Repository
}

// NewService instantiates Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// LogApproval persists workflow approvals signatures.
func (s *Service) LogApproval(ctx context.Context, a *approvalDomain.Approval) (*approvalDomain.Approval, error) {
	if err := a.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, a); err != nil {
		return nil, err
	}
	return a, nil
}
