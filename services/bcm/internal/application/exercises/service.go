package exercises

import (
	"context"
	"fmt"
	"time"

	"prahari/services/bcm/internal/domain/continuityexercise"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveExercise(ctx context.Context, ex *continuityexercise.Exercise) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ScheduleExercise(ctx context.Context, ex *continuityexercise.Exercise) error {
	ex.ID = fmt.Sprintf("ex-%d", time.Now().UnixNano())
	ex.Status = "SCHEDULED"
	ex.CreatedAt = time.Now()

	if err := s.repo.SaveExercise(ctx, ex); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Business continuity testing exercise scheduled", prahariLogger.String("title", ex.Title))
	return nil
}
