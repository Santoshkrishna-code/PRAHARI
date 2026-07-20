package export

import (
	"context"
	"fmt"

	"prahari/services/barrier/internal/domain/barrier"
	"prahari/services/barrier/internal/domain/search"
)

type Repository interface {
	SearchBarriers(ctx context.Context, criteria *search.Criteria) ([]*barrier.Barrier, int64, error)
	GetBarrierByID(ctx context.Context, id string) (*barrier.Barrier, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportCSV(ctx context.Context, criteria *search.Criteria) ([]byte, error) {
	barriers, _, err := s.repo.SearchBarriers(ctx, criteria)
	if err != nil {
		return nil, err
	}

	csvData := "ID,BarrierCode,PlantID,UnitID,Title,Type,SILLevel,IsIPL,HealthScore,Status\n"
	for _, b := range barriers {
		csvData += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%t,%.2f,%s\n",
			b.ID, b.BarrierCode, b.PlantID, b.UnitID, b.Title, b.Type, b.SILLevel, b.IsIPL, b.HealthScore, b.Status)
	}

	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	b, err := s.repo.GetBarrierByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock Barrier Management Executive Report\nID: %s\nCode: %s\nTitle: %s\nType: %s\nSIL: %s\nHealth Score: %.2f%%\nStatus: %s\n",
		b.ID, b.BarrierCode, b.Title, b.Type, b.SILLevel, b.HealthScore, b.Status)
	return []byte(pdfDoc), nil
}
