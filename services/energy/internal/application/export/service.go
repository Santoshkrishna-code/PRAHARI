package export

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"

	"prahari/services/energy/internal/domain/energyprofile"
	"prahari/services/energy/internal/domain/search"
)

type Repository interface {
	SearchProfiles(ctx context.Context, criteria search.Criteria) ([]energyprofile.Profile, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) WriteCSVReport(ctx context.Context, writer io.Writer) error {
	profiles, err := s.repo.SearchProfiles(ctx, search.Criteria{Limit: 1000, Offset: 0})
	if err != nil {
		return err
	}

	csvWriter := csv.NewWriter(writer)
	defer csvWriter.Flush()

	// Headers
	_ = csvWriter.Write([]string{"ID", "PlantID", "DepartmentID", "FacilityName", "TargetScore"})

	for _, p := range profiles {
		row := []string{
			p.ID,
			p.PlantID,
			p.DepartmentID,
			p.FacilityName,
			fmt.Sprintf("%.2f", p.TargetScore),
		}
		if err := csvWriter.Write(row); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) GetPDFReport(ctx context.Context, recordID string) ([]byte, error) {
	// Returns a mock PDF file representing facility energy audit report details
	reportContent := fmt.Sprintf("PRAHARI INDUSTRIAL ENERGY AUDIT REPORT\nRecordID: %s\nStatus: VERIFIED\n", recordID)
	return []byte(reportContent), nil
}
