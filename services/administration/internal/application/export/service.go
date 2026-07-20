package export

import (
	"context"
	"fmt"

	"prahari/services/administration/internal/domain/search"
	"prahari/services/administration/internal/domain/tenant"
)

type Repository interface {
	SearchTenants(ctx context.Context, criteria *search.Criteria) ([]*tenant.Tenant, int64, error)
	GetTenantByID(ctx context.Context, id string) (*tenant.Tenant, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportCSV(ctx context.Context, criteria *search.Criteria) ([]byte, error) {
	tenants, _, err := s.repo.SearchTenants(ctx, criteria)
	if err != nil {
		return nil, err
	}

	csvData := "ID,Name,Domain,Status\n"
	for _, t := range tenants {
		csvData += fmt.Sprintf("%s,%s,%s,%s\n", t.ID, t.Name, t.Domain, t.Status)
	}

	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	t, err := s.repo.GetTenantByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock Tenant Provisioning Summary\nID: %s\nName: %s\nDomain: %s\nStatus: %s\n",
		t.ID, t.Name, t.Domain, t.Status)
	return []byte(pdfDoc), nil
}
