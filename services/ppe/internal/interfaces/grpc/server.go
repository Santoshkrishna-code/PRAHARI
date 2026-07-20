package grpc

import (
	"context"

	catalogApp "prahari/services/ppe/internal/application/catalog"
	inspectionApp "prahari/services/ppe/internal/application/inspection"
	issuanceApp "prahari/services/ppe/internal/application/issuance"
	reportingApp "prahari/services/ppe/internal/application/reporting"
	"prahari/services/ppe/internal/domain/ppe"
	"prahari/services/ppe/internal/domain/ppeinspection"
	"prahari/services/ppe/internal/domain/ppeissue"
)

type Server struct {
	catalogSvc    *catalogApp.Service
	issuanceSvc   *issuanceApp.Service
	inspectionSvc *inspectionApp.Service
	reportingSvc  *reportingApp.Service
}

func NewServer(
	catalogSvc *catalogApp.Service,
	issuanceSvc *issuanceApp.Service,
	inspectionSvc *inspectionApp.Service,
	reportingSvc *reportingApp.Service,
) *Server {
	return &Server{
		catalogSvc:    catalogSvc,
		issuanceSvc:   issuanceSvc,
		inspectionSvc: inspectionSvc,
		reportingSvc:  reportingSvc,
	}
}

func (s *Server) CreatePPE(ctx context.Context, p *ppe.PPE) error {
	return s.catalogSvc.CreateCatalogPPE(ctx, p)
}

func (s *Server) IssuePPE(ctx context.Context, rec *ppeissue.Record) error {
	return s.issuanceSvc.IssuePPEItem(ctx, rec)
}

func (s *Server) InspectPPE(ctx context.Context, rec *ppeinspection.Record) error {
	return s.inspectionSvc.InspectPPEItem(ctx, rec)
}

func (s *Server) GetPPE(ctx context.Context, id string) (*ppe.PPE, error) {
	return s.reportingSvc.GetPPE(ctx, id)
}
