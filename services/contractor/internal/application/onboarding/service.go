package onboarding

import (
	"context"

	workerDomain "prahari/services/contractor/internal/domain/worker"
)

// Repository manages worker persistence stores.
type Repository interface {
	Create(ctx context.Context, w *workerDomain.Worker) error
	FindByID(ctx context.Context, id string) (*workerDomain.Worker, error)
	FindByContractorID(ctx context.Context, contractorID string) ([]*workerDomain.Worker, error)
}

// Service manages worker onboarding lifecycle.
type Service struct {
	repo Repository
}

// NewService instantiates Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// OnboardWorker inserts worker profile.
func (s *Service) OnboardWorker(ctx context.Context, w *workerDomain.Worker) (*workerDomain.Worker, error) {
	w.OnboardingStatus = "Pending"
	if err := w.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, w); err != nil {
		return nil, err
	}
	return w, nil
}
