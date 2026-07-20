package export

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"

	"prahari/services/esg/internal/domain/esgobjective"
	"prahari/services/esg/internal/domain/search"
)

type Repository interface {
	SearchObjectives(ctx context.Context, criteria search.Criteria) ([]esgobjective.Objective, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) WriteCSVReport(ctx context.Context, writer io.Writer) error {
	objectives, err := s.repo.SearchObjectives(ctx, search.Criteria{Limit: 1000, Offset: 0})
	if err != nil {
		return err
	}

	csvWriter := csv.NewWriter(writer)
	defer csvWriter.Flush()

	// Headers
	_ = csvWriter.Write([]string{"ID", "BusinessUnitID", "Title", "Category", "Status"})

	for _, o := range objectives {
		row := []string{
			o.ID,
			o.BusinessUnitID,
			o.Title,
			o.Category,
			o.Status,
		}
		if err := csvWriter.Write(row); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) GetPDFReport(ctx context.Context, recordID string) ([]byte, error) {
	// Returns a mock PDF file representing corporate ESG report details
	reportContent := fmt.Sprintf("PRAHARI ESG & SUSTAINABILITY DISCLOSURE REPORT\nRecordID: %s\nStatus: VALIDATED\n", recordID)
	return []byte(reportContent), nil
}
