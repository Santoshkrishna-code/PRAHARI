package export

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"

	"prahari/services/occupational-health/internal/domain/healthprofile"
)

type Repository interface {
	GetAllProfiles(ctx context.Context) ([]healthprofile.HealthProfile, error)
	GetProfile(ctx context.Context, id string) (*healthprofile.HealthProfile, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) WriteCSVReport(ctx context.Context, writer io.Writer) error {
	profiles, err := s.repo.GetAllProfiles(ctx)
	if err != nil {
		return err
	}

	csvWriter := csv.NewWriter(writer)
	defer csvWriter.Flush()

	// Write headers
	_ = csvWriter.Write([]string{"ID", "WorkerID", "WorkerType", "DepartmentID", "ClearanceStatus", "MedicalStatus"})

	for _, p := range profiles {
		row := []string{
			p.ID,
			p.WorkerID,
			p.WorkerType,
			p.DepartmentID,
			p.ClearanceStatus,
			p.MedicalStatus,
		}
		if err := csvWriter.Write(row); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) GetPDFReport(ctx context.Context, profileID string) ([]byte, error) {
	p, err := s.repo.GetProfile(ctx, profileID)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, fmt.Errorf("profile not found: %s", profileID)
	}

	// Returns mock PDF report bytes matching the worker's fitness card details
	reportContent := fmt.Sprintf("PRAHARI OCCUPATIONAL MEDICAL FIT CARD\nWorker: %s\nStatus: %s\nClearance: %s\n",
		p.WorkerID, p.MedicalStatus, p.ClearanceStatus)

	return []byte(reportContent), nil
}
