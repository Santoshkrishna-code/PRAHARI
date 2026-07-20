package resilience

import (
	"context"
	"fmt"
	"time"

	"prahari/services/bcm/internal/domain/resilienceassessment"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveAssessment(ctx context.Context, ra *resilienceassessment.Assessment) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AssessResilience(ctx context.Context, ra *resilienceassessment.Assessment) error {
	ra.ID = fmt.Sprintf("ra-%d", time.Now().UnixNano())
	ra.AssessedAt = time.Now()

	if err := s.repo.SaveAssessment(ctx, ra); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Organizational ISO 22301 resilience assessment completed",
		prahariLogger.String("business_unit", ra.BusinessUnit))
	return nil
}
