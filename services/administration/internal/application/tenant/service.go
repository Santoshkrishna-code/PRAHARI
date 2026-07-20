package tenant

import (
	"context"
	"fmt"
	"time"

	"prahari/services/administration/internal/domain/events"
	"prahari/services/administration/internal/domain/tenant"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveTenant(ctx context.Context, t *tenant.Tenant) error
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

func (s *Service) CreateTenant(ctx context.Context, t *tenant.Tenant) error {
	t.ID = fmt.Sprintf("ten-%d", time.Now().UnixNano())
	t.Status = "ACTIVE"
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()

	if err := s.repo.SaveTenant(ctx, t); err != nil {
		return fmt.Errorf("failed to save tenant: %w", err)
	}

	_ = s.publisher.Publish(ctx, events.EventTenantCreated, t)
	prahariLogger.Info(ctx, "SaaS tenant created",
		prahariLogger.String("tenant_id", t.ID),
		prahariLogger.String("domain", t.Domain))
	return nil
}
