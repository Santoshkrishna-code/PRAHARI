package export

import (
	"context"
	"fmt"

	"prahari/services/moc/internal/domain/changerequest"
	"prahari/services/moc/internal/domain/search"
)

type Repository interface {
	SearchRequests(ctx context.Context, criteria *search.Criteria) ([]*changerequest.Request, int64, error)
	GetRequestByID(ctx context.Context, id string) (*changerequest.Request, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportCSV(ctx context.Context, criteria *search.Criteria) ([]byte, error) {
	requests, _, err := s.repo.SearchRequests(ctx, criteria)
	if err != nil {
		return nil, err
	}

	csvData := "ID,MOCNumber,PlantID,Title,Category,ChangeType,RiskLevel,Status,RequesterID\n"
	for _, req := range requests {
		csvData += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
			req.ID, req.MOCNumber, req.PlantID, req.Title, req.Category, req.ChangeType, req.RiskLevel, req.Status, req.RequesterID)
	}

	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	req, err := s.repo.GetRequestByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock MOC Executive Report\nID: %s\nMOC Number: %s\nTitle: %s\nCategory: %s\nRisk Level: %s\nStatus: %s\n",
		req.ID, req.MOCNumber, req.Title, req.Category, req.RiskLevel, req.Status)
	return []byte(pdfDoc), nil
}
