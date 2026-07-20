package export

import (
	"context"
	"fmt"

	"prahari/services/digitaltwin/internal/domain/search"
	"prahari/services/digitaltwin/internal/domain/twin"
)

type Repository interface {
	SearchTwins(ctx context.Context, criteria *search.Criteria) ([]*twin.DigitalTwin, int64, error)
	GetTwinByID(ctx context.Context, id string) (*twin.DigitalTwin, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportCSV(ctx context.Context, criteria *search.Criteria) ([]byte, error) {
	twins, _, err := s.repo.SearchTwins(ctx, criteria)
	if err != nil {
		return nil, err
	}

	csvData := "ID,PlantID,Name,Status,Version,CreatedAt\n"
	for _, t := range twins {
		csvData += fmt.Sprintf("%s,%s,%s,%s,%d,%s\n", t.ID, t.PlantID, t.Name, t.Status, t.Version, t.CreatedAt)
	}

	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	t, err := s.repo.GetTwinByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock Digital Twin Export\nID: %s\nPlantID: %s\nName: %s\nStatus: %s\nVersion: %d\n",
		t.ID, t.PlantID, t.Name, t.Status, t.Version)
	return []byte(pdfDoc), nil
}
