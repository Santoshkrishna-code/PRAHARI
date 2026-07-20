package surveillance

import (
	"context"
	"time"

	"prahari/services/occupational-health/internal/domain/surveillance"
)

// Repository manages monitoring programs store.
type Repository interface {
	SaveSurveillance(ctx context.Context, s *surveillance.HealthSurveillance) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) StartSurveillance(ctx context.Context, hs *surveillance.HealthSurveillance) error {
	hs.StartDate = time.Now()
	hs.CreatedAt = time.Now()
	hs.UpdatedAt = time.Now()
	if err := hs.Validate(); err != nil {
		return err
	}
	return s.repo.SaveSurveillance(ctx, hs)
}
