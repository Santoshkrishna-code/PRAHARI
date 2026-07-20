package optimization

import (
	"context"
	"fmt"
	"time"

	"prahari/services/water/internal/domain/optimization"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveRecommendation(ctx context.Context, rec *optimization.Recommendation) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateRecommendation(ctx context.Context, rec *optimization.Recommendation) error {
	rec.ID = fmt.Sprintf("opt-%d", time.Now().UnixNano())
	rec.CreatedAt = time.Now()
	if rec.Status == "" {
		rec.Status = "RECOMMENDED"
	}

	if err := s.repo.SaveRecommendation(ctx, rec); err != nil {
		return fmt.Errorf("failed to save optimization recommendation: %w", err)
	}

	prahariLogger.Info(ctx, "Water optimization recommendation created", prahariLogger.String("id", rec.ID))
	return nil
}
