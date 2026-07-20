package exposure

import (
	"context"
	"time"

	"prahari/services/occupational-health/internal/domain/exposure"
	"prahari/services/occupational-health/internal/domain/policy"
)

// Repository manages monitoring logs store.
type Repository interface {
	SaveExposure(ctx context.Context, e *exposure.ExposureRecord) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) RecordExposure(ctx context.Context, er *exposure.ExposureRecord) error {
	er.MonitoringDate = time.Now()
	er.CreatedAt = time.Now()
	er.UpdatedAt = time.Now()
	if err := er.Validate(); err != nil {
		return err
	}
	er.IsOverLimit = policy.EvaluateExposureLimit(er)
	return s.repo.SaveExposure(ctx, er)
}
