package export

import (
	"context"
	"fmt"

	"prahari/services/ppe/internal/domain/ppe"
	"prahari/services/ppe/internal/domain/search"
)

type Repository interface {
	SearchPPEs(ctx context.Context, criteria *search.Criteria) ([]*ppe.PPE, int64, error)
	GetPPEByID(ctx context.Context, id string) (*ppe.PPE, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportCSV(ctx context.Context, criteria *search.Criteria) ([]byte, error) {
	ppeList, _, err := s.repo.SearchPPEs(ctx, criteria)
	if err != nil {
		return nil, err
	}

	csvData := "ID,ModelName,PlantID,CategoryID,Manufacturer,PartNumber,StandardRef\n"
	for _, p := range ppeList {
		csvData += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s\n",
			p.ID, p.ModelName, p.PlantID, p.CategoryID, p.Manufacturer, p.PartNumber, p.StandardRef)
	}

	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	p, err := s.repo.GetPPEByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock PPE Model Catalog Details & Standard Compliance Certification Report\nID: %s\nModel: %s\nPlant: %s\nStandard: %s\n",
		p.ID, p.ModelName, p.PlantID, p.StandardRef)
	return []byte(pdfDoc), nil
}
