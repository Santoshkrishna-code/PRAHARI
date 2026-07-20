package export

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"

	"prahari/services/environmental/internal/domain/environment"
	"prahari/services/environmental/internal/domain/search"
)

type Repository interface {
	SearchAspects(ctx context.Context, criteria search.Criteria) ([]environment.EnvironmentalAspect, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) WriteCSVReport(ctx context.Context, writer io.Writer) error {
	aspects, err := s.repo.SearchAspects(ctx, search.Criteria{Limit: 1000, Offset: 0})
	if err != nil {
		return err
	}

	csvWriter := csv.NewWriter(writer)
	defer csvWriter.Flush()

	// Headers
	_ = csvWriter.Write([]string{"ID", "PlantID", "DepartmentID", "Name", "AspectCategory"})

	for _, a := range aspects {
		row := []string{
			a.ID,
			a.PlantID,
			a.DepartmentID,
			a.Name,
			a.AspectCategory,
		}
		if err := csvWriter.Write(row); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) GetPDFReport(ctx context.Context, recordID string) ([]byte, error) {
	// Returns a mock PDF content byte slice representing environmental permit/record details
	reportContent := fmt.Sprintf("PRAHARI ENVIRONMENTAL RECORD REPORT\nRecordID: %s\nStatus: EVALUATED\n", recordID)
	return []byte(reportContent), nil
}
