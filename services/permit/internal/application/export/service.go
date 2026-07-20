package export

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"

	permitDomain "prahari/services/permit/internal/domain/permit"
	searchDomain "prahari/services/permit/internal/domain/search"
)

// SearchRepository retrieves permits for export.
type SearchRepository interface {
	Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*permitDomain.Permit, int, error)
}

// Service manages CSV/PDF binary formatting.
type Service struct {
	repo SearchRepository
}

// NewService instantiates an Export Service.
func NewService(repo SearchRepository) *Service {
	return &Service{repo: repo}
}

// ExportCSV streams a CSV array matching criteria.
func (s *Service) ExportCSV(ctx context.Context, criteria *searchDomain.Criteria) ([]byte, error) {
	criteria.Normalize()
	criteria.PageSize = 10000 // Export upper cap

	permits, _, err := s.repo.Search(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve permits for export: %w", err)
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	header := []string{
		"Permit Number", "Title", "Status", "Risk Level",
		"Applicant ID", "Department ID", "Work Area ID", "Planned Start", "Planned End",
	}
	_ = writer.Write(header)

	for _, p := range permits {
		row := []string{
			p.PermitNumber,
			p.Title,
			p.StatusCode,
			string(p.RiskLevel),
			p.ApplicantID,
			p.DepartmentID,
			p.WorkAreaID,
			p.PlannedStartAt.Format("2006-01-02 15:04:05"),
			p.PlannedEndAt.Format("2006-01-02 15:04:05"),
		}
		_ = writer.Write(row)
	}

	writer.Flush()
	return buf.Bytes(), nil
}

// ExportPDF returns permit summary layout string bytes.
func (s *Service) ExportPDF(ctx context.Context, permitID string) ([]byte, error) {
	criteria := &searchDomain.Criteria{
		PermitNumber: permitID,
		PageSize:     1,
	}
	criteria.Normalize()

	permits, _, err := s.repo.Search(ctx, criteria)
	if err != nil || len(permits) == 0 {
		return nil, fmt.Errorf("permit not found or failed query: %w", err)
	}

	p := permits[0]

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("PERMIT-TO-WORK AUTHORIZATION SHEET\n"))
	buf.WriteString(fmt.Sprintf("==================================\n\n"))
	buf.WriteString(fmt.Sprintf("Permit Number: %s\n", p.PermitNumber))
	buf.WriteString(fmt.Sprintf("Title: %s\n", p.Title))
	buf.WriteString(fmt.Sprintf("Status: %s\n", p.StatusCode))
	buf.WriteString(fmt.Sprintf("Risk Level: %s\n", p.RiskLevel))
	buf.WriteString(fmt.Sprintf("Applicant: %s\n", p.ApplicantID))
	buf.WriteString(fmt.Sprintf("Work Area: %s\n\n", p.WorkAreaID))
	buf.WriteString(fmt.Sprintf("Description of Work:\n%s\n", p.WorkDescription))

	return buf.Bytes(), nil
}
