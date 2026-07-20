package hierarchy

import (
	"context"
	"fmt"
	"time"

	"prahari/services/administration/internal/domain/events"
	"prahari/services/administration/internal/domain/organization"
	"prahari/services/administration/internal/domain/plant"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveOrganization(ctx context.Context, org *organization.Organization) error
	SavePlant(ctx context.Context, plt *plant.Plant) error
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

func (s *Service) CreateOrganization(ctx context.Context, org *organization.Organization) error {
	org.ID = fmt.Sprintf("org-%d", time.Now().UnixNano())
	org.Status = "CONFIGURED"
	org.CreatedAt = time.Now()
	org.UpdatedAt = time.Now()

	if err := s.repo.SaveOrganization(ctx, org); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventOrganizationCreated, org)
	prahariLogger.Info(ctx, "Organization created", prahariLogger.String("org_id", org.ID))
	return nil
}

func (s *Service) CreatePlant(ctx context.Context, plt *plant.Plant) error {
	plt.ID = fmt.Sprintf("plt-%d", time.Now().UnixNano())
	plt.CreatedAt = time.Now()
	plt.UpdatedAt = time.Now()

	if err := s.repo.SavePlant(ctx, plt); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventPlantCreated, plt)
	return nil
}
