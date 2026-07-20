package export

import (
	"context"
	"fmt"

	"prahari/services/action/internal/domain/action"
	"prahari/services/action/internal/domain/search"
)

type Repository interface {
	SearchActions(ctx context.Context, criteria *search.Criteria) ([]*action.Action, int64, error)
	GetActionByID(ctx context.Context, id string) (*action.Action, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportCSV(ctx context.Context, criteria *search.Criteria) ([]byte, error) {
	actions, _, err := s.repo.SearchActions(ctx, criteria)
	if err != nil {
		return nil, err
	}

	csvData := "ID,SourceModule,SourceRefID,Title,ActionType,Status,DueDate\n"
	for _, c := range actions {
		csvData += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s\n",
			c.ID, c.SourceModule, c.SourceRefID, c.Title, c.ActionType, c.Status, c.DueDate.Format("2006-01-02 15:04:05"))
	}

	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	c, err := s.repo.GetActionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock Corrective and Preventive Action (CAPA) Closeout Report\nID: %s\nSource: %s\nStatus: %s\nTitle: %s\n",
		c.ID, c.SourceModule, c.Status, c.Title)
	return []byte(pdfDoc), nil
}
