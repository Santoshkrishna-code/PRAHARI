package catalog

import (
	"context"
	"fmt"
	"time"

	"prahari/services/ppe/internal/domain/events"
	"prahari/services/ppe/internal/domain/ppe"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SavePPE(ctx context.Context, p *ppe.PPE) error
}

type EventPublisher interface {
	Publish(ctx context.Context, eventType string, payload any) error
}

type Service struct {
	repo      Repository
	publisher EventPublisher
}

func NewService(repo Repository, pub EventPublisher) *Service {
	return &Service{
		repo:      repo,
		publisher: pub,
	}
}

func (s *Service) CreateCatalogPPE(ctx context.Context, p *ppe.PPE) error {
	p.ID = fmt.Sprintf("ppe-%d", time.Now().UnixNano())
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	if err := s.repo.SavePPE(ctx, p); err != nil {
		return fmt.Errorf("failed to save ppe: %w", err)
	}

	_ = s.publisher.Publish(ctx, events.EventPPECreated, p)
	prahariLogger.Info(ctx, "New PPE model registered in catalog",
		prahariLogger.String("model_name", p.ModelName))
	return nil
}
