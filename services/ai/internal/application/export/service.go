package export

import (
	"context"
	"fmt"

	"prahari/services/ai/internal/domain/document"
	"prahari/services/ai/internal/domain/search"
)

type Repository interface {
	SearchDocuments(ctx context.Context, criteria *search.Criteria) ([]*document.Doc, int64, error)
	GetDocumentByID(ctx context.Context, id string) (*document.Doc, error)
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

	csvData := "ID,SourceID,Title,CreatedAt\n"
	for _, d := range docs {
		csvData += fmt.Sprintf("%s,%s,%s,%s\n", d.ID, d.SourceID, d.Title, d.CreatedAt)
	}

	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	d, err := s.repo.GetDocumentByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock Document Vector Metadata\nID: %s\nSourceID: %s\nTitle: %s\n",
		d.ID, d.SourceID, d.Title)
	return []byte(pdfDoc), nil
}
