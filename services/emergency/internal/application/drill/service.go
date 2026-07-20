package drill

import (
	"context"
	"fmt"
	"time"

	"prahari/services/emergency/internal/domain/drill"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveDrill(ctx context.Context, d *drill.Drill) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ScheduleDrill(ctx context.Context, d *drill.Drill) error {
	d.ID = fmt.Sprintf("drl-%d", time.Now().UnixNano())
	d.Status = "SCHEDULED"
	d.CreatedAt = time.Now()

	if err := s.repo.SaveDrill(ctx, d); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Emergency preparedness drill scheduled", prahariLogger.String("title", d.Title))
	return nil
}
