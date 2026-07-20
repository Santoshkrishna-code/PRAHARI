package export

import (
	"context"
	"fmt"

	"prahari/services/loto/internal/domain/isolationcertificate"
	"prahari/services/loto/internal/domain/search"
)

type Repository interface {
	SearchCertificates(ctx context.Context, criteria *search.Criteria) ([]*isolationcertificate.Certificate, int64, error)
	GetCertificateByID(ctx context.Context, id string) (*isolationcertificate.Certificate, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportCSV(ctx context.Context, criteria *search.Criteria) ([]byte, error) {
	certificates, _, err := s.repo.SearchCertificates(ctx, criteria)
	if err != nil {
		return nil, err
	}

	csvData := "ID,PlanID,PermitID,IssuerID,ReceiverID,Status\n"
	for _, c := range certificates {
		csvData += fmt.Sprintf("%s,%s,%s,%s,%s,%s\n",
			c.ID, c.PlanID, c.PermitID, c.IssuerID, c.ReceiverID, c.Status)
	}

	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	c, err := s.repo.GetCertificateByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock Hazardous Energy Control LOTO Certificate Audit Report\nID: %s\nPlan: %s\nPermit ID: %s\nStatus: %s\n",
		c.ID, c.PlanID, c.PermitID, c.Status)
	return []byte(pdfDoc), nil
}
