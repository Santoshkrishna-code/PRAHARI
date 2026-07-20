package export

import (
	"context"
	"fmt"

	"prahari/services/integration/internal/domain/connector"
	"prahari/services/integration/internal/domain/search"
)

type Repository interface {
	SearchConnectors(ctx context.Context, criteria *search.Criteria) ([]*connector.Connector, int64, error)
	GetConnectorByID(ctx context.Context, id string) (*connector.Connector, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportCSV(ctx context.Context, criteria *search.Criteria) ([]byte, error) {
	connectors, _, err := s.repo.SearchConnectors(ctx, criteria)
	if err != nil {
		return nil, err
	}

	csvData := "ID,PlantID,Name,Type,Status\n"
	for _, c := range connectors {
		csvData += fmt.Sprintf("%s,%s,%s,%s,%s\n", c.ID, c.PlantID, c.Name, c.Type, c.Status)
	}

	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	c, err := s.repo.GetConnectorByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock Connector Summary\nID: %s\nName: %s\nType: %s\nStatus: %s\n",
		c.ID, c.Name, c.Type, c.Status)
	return []byte(pdfDoc), nil
}
