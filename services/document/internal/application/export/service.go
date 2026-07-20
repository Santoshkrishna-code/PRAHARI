package export

import (
	"context"
	"fmt"

	"prahari/services/document/internal/domain/document"
	"prahari/services/document/internal/domain/search"
)

type Repository interface {
	SearchDocuments(ctx context.Context, criteria *search.Criteria) ([]*document.Document, int64, error)
	GetDocumentByID(ctx context.Context, id string) (*document.Document, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportCSV(ctx context.Context, criteria *search.Criteria) ([]byte, error) {
	docs, _, err := s.repo.SearchDocuments(ctx, criteria)
	if err != nil {
		return nil, err
	}

	csvData := "ID,DocumentNumber,PlantID,Title,DocumentType,CurrentVersion,Status,OwnerID\n"
	for _, d := range docs {
		csvData += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s\n",
			d.ID, d.DocumentNumber, d.PlantID, d.Title, d.DocumentType, d.CurrentVersion, d.Status, d.OwnerID)
	}

	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	d, err := s.repo.GetDocumentByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock Controlled Document Governance Report\nID: %s\nNumber: %s\nTitle: %s\nType: %s\nVersion: %s\nStatus: %s\n",
		d.ID, d.DocumentNumber, d.Title, d.DocumentType, d.CurrentVersion, d.Status)
	return []byte(pdfDoc), nil
}
