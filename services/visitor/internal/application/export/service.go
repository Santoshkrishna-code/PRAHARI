package export

import (
	"context"
	"fmt"

	"prahari/services/visitor/internal/domain/search"
	"prahari/services/visitor/internal/domain/visit"
)

type Repository interface {
	SearchVisits(ctx context.Context, criteria *search.Criteria) ([]*visit.Visit, int64, error)
	GetVisitByID(ctx context.Context, id string) (*visit.Visit, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportCSV(ctx context.Context, criteria *search.Criteria) ([]byte, error) {
	visits, _, err := s.repo.SearchVisits(ctx, criteria)
	if err != nil {
		return nil, err
	}

	csvData := "ID,VisitorID,HostID,PlantID,Purpose,ScheduledIn,ScheduledOut,Status\n"
	for _, v := range visits {
		csvData += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s\n",
			v.ID, v.VisitorID, v.HostID, v.PlantID, v.Purpose, v.ScheduledIn.Format("2006-01-02 15:04:05"), v.ScheduledOut.Format("2006-01-02 15:04:05"), v.Status)
	}

	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	v, err := s.repo.GetVisitByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock Plant Visitor Access Log & Badge Pass Report\nID: %s\nVisitor: %s\nHost: %s\nPlant: %s\nStatus: %s\n",
		v.ID, v.VisitorID, v.HostID, v.PlantID, v.Status)
	return []byte(pdfDoc), nil
}
