package dashboards

import (
	"context"
	"fmt"
	"time"

	"prahari/services/analytics/internal/domain/dashboard"
	"prahari/services/analytics/internal/domain/events"
)

type Repository interface {
	SaveDashboard(ctx context.Context, d *dashboard.Dashboard) error
	GetDashboardByID(ctx context.Context, id string) (*dashboard.Dashboard, error)
}

type EventPublisher interface {
	Publish(ctx context.Context, eventType string, payload any) error
}

type Service struct {
	repo      Repository
	publisher EventPublisher
}

func NewService(repo Repository, pub EventPublisher) *Service {
	return &Service{repo: repo, publisher: pub}
}

func (s *Service) CreateDashboard(ctx context.Context, d *dashboard.Dashboard) error {
	d.ID = fmt.Sprintf("db-%d", time.Now().UnixNano())
	d.UpdatedAt = time.Now()

	if err := s.repo.SaveDashboard(ctx, d); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventDashboardUpdated, d)
	return nil
}

func (s *Service) GetDashboard(ctx context.Context, id string) (*dashboard.Dashboard, error) {
	return s.repo.GetDashboardByID(ctx, id)
}
