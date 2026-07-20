package grpc

import (
	"context"

	configurationApp "prahari/services/administration/internal/application/configuration"
	hierarchyApp "prahari/services/administration/internal/application/hierarchy"
	reportingApp "prahari/services/administration/internal/application/reporting"
	tenantApp "prahari/services/administration/internal/application/tenant"
	"prahari/services/administration/internal/domain/configuration"
	"prahari/services/administration/internal/domain/organization"
	"prahari/services/administration/internal/domain/plant"
	"prahari/services/administration/internal/domain/tenant"
)

type Server struct {
	tenantSvc        *tenantApp.Service
	hierarchySvc     *hierarchyApp.Service
	configurationSvc *configurationApp.Service
	reportingSvc     *reportingApp.Service
}

func NewServer(
	tenantSvc *tenantApp.Service,
	hierarchySvc *hierarchyApp.Service,
	configurationSvc *configurationApp.Service,
	reportingSvc *reportingApp.Service,
) *Server {
	return &Server{
		tenantSvc:        tenantSvc,
		hierarchySvc:     hierarchySvc,
		configurationSvc: configurationSvc,
		reportingSvc:     reportingSvc,
	}
}

func (s *Server) CreateTenant(ctx context.Context, t *tenant.Tenant) error {
	return s.tenantSvc.CreateTenant(ctx, t)
}

func (s *Server) CreateOrganization(ctx context.Context, org *organization.Organization) error {
	return s.hierarchySvc.CreateOrganization(ctx, org)
}

func (s *Server) CreatePlant(ctx context.Context, plt *plant.Plant) error {
	return s.hierarchySvc.CreatePlant(ctx, plt)
}

func (s *Server) ResolveOrganization(ctx context.Context, id string) (*tenant.Tenant, error) {
	return s.reportingSvc.GetTenant(ctx, id)
}

func (s *Server) ResolvePlant(ctx context.Context, id string) (*tenant.Tenant, error) {
	return s.reportingSvc.GetTenant(ctx, id)
}

func (s *Server) GetConfiguration(ctx context.Context, tenantID, key string) (*configuration.Param, error) {
	return s.configurationSvc.GetConfiguration(ctx, tenantID, key)
}
