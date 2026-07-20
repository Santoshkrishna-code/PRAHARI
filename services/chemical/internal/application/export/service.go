package export

import (
	"context"
	"fmt"

	"prahari/services/chemical/internal/domain/chemical"
	"prahari/services/chemical/internal/domain/search"
)

type Repository interface {
	SearchChemicals(ctx context.Context, criteria *search.Criteria) ([]*chemical.Chemical, int64, error)
	GetChemicalByID(ctx context.Context, id string) (*chemical.Chemical, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportCSV(ctx context.Context, criteria *search.Criteria) ([]byte, error) {
	chems, _, err := s.repo.SearchChemicals(ctx, criteria)
	if err != nil {
		return nil, err
	}

	csvData := "ID,PlantID,Name,CASNumber,Formula,IsRestricted,Status\n"
	for _, c := range chems {
		csvData += fmt.Sprintf("%s,%s,%s,%s,%s,%t,%s\n",
			c.ID, c.PlantID, c.Name, c.CASNumber, c.Formula, c.IsRestricted, c.Status)
	}

	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	c, err := s.repo.GetChemicalByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock Safety Data Sheet Summary\nID: %s\nName: %s\nCAS: %s\nRestricted: %t\n",
		c.ID, c.Name, c.CASNumber, c.IsRestricted)
	return []byte(pdfDoc), nil
}
