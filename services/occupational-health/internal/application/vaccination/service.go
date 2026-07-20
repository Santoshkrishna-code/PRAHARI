package vaccination

import (
	"context"
	"time"

	"prahari/services/occupational-health/internal/domain/vaccination"
)

// Repository manages immunization profiles store.
type Repository interface {
	SaveVaccination(ctx context.Context, v *vaccination.Vaccination) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) RegisterVaccination(ctx context.Context, v *vaccination.Vaccination) error {
	v.CreatedAt = time.Now()
	v.UpdatedAt = time.Now()
	if err := v.Validate(); err != nil {
		return err
	}
	return s.repo.SaveVaccination(ctx, v)
}
