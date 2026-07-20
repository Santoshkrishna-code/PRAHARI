package export

import (
	"context"
	"fmt"

	"prahari/services/bcm/internal/domain/continuityplan"
	"prahari/services/bcm/internal/domain/search"
)

type Repository interface {
	SearchPlans(ctx context.Context, criteria *search.Criteria) ([]*continuityplan.Plan, int64, error)
	GetPlanByID(ctx context.Context, id string) (*continuityplan.Plan, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportCSV(ctx context.Context, criteria *search.Criteria) ([]byte, error) {
	plans, _, err := s.repo.SearchPlans(ctx, criteria)
	if err != nil {
		return nil, err
	}

	csvData := "ID,PlanNumber,PlantID,BusinessUnit,Title,Version,Status,ApprovedBy\n"
	for _, p := range plans {
		csvData += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s\n",
			p.ID, p.PlanNumber, p.PlantID, p.BusinessUnit, p.Title, p.Version, p.Status, p.ApprovedBy)
	}

	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	p, err := s.repo.GetPlanByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock Business Continuity Management ISO 22301 Report\nID: %s\nNumber: %s\nTitle: %s\nUnit: %s\nVersion: %s\nStatus: %s\n",
		p.ID, p.PlanNumber, p.Title, p.BusinessUnit, p.Version, p.Status)
	return []byte(pdfDoc), nil
}
