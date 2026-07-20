package licensing

import (
	"context"
	"fmt"
	"time"

	"prahari/services/administration/internal/domain/events"
	"prahari/services/administration/internal/domain/license"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveLicense(ctx context.Context, lic *license.License) error
	GetLicenseByTenantID(ctx context.Context, tenantID string) (*license.License, error)
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

func (s *Service) AssignLicense(ctx context.Context, lic *license.License) error {
	lic.ID = fmt.Sprintf("lic-%d", time.Now().UnixNano())
	lic.LicensedAt = time.Now()

	if err := s.repo.SaveLicense(ctx, lic); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventLicenseUpdated, lic)
	prahariLogger.Info(ctx, "Tenant license assigned",
		prahariLogger.String("tenant_id", lic.TenantID),
		prahariLogger.String("tier", lic.Tier))
	return nil
}

func (s *Service) GetLicense(ctx context.Context, tenantID string) (*license.License, error) {
	return s.repo.GetLicenseByTenantID(ctx, tenantID)
}
