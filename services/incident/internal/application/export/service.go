package export

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"

	incidentDomain "prahari/services/incident/internal/domain/incident"
	searchDomain "prahari/services/incident/internal/domain/search"
)

// SearchRepository defines the persistence port for retrieving incidents for export.
type SearchRepository interface {
	Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*incidentDomain.Incident, int, error)
}

// Service orchestrates data export operations.
type Service struct {
	repo SearchRepository
}

// NewService constructs a Service with the search repository injected.
func NewService(repo SearchRepository) *Service {
	return &Service{repo: repo}
}

// ExportCSV generates a CSV byte stream of incidents matching the provided criteria.
func (s *Service) ExportCSV(ctx context.Context, criteria *searchDomain.Criteria) ([]byte, error) {
	criteria.Normalize()
	criteria.PageSize = 10000 // Export up to 10k records

	incidents, _, err := s.repo.Search(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve incidents for export: %w", err)
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// Write header row
	header := []string{
		"Incident Number", "Title", "Type", "Severity", "Priority",
		"Status", "Department", "Location", "Occurred At", "Reported At",
	}
	if err := writer.Write(header); err != nil {
		return nil, fmt.Errorf("failed to write CSV header: %w", err)
	}

	// Write data rows
	for _, inc := range incidents {
		row := []string{
			inc.IncidentNumber,
			inc.Title,
			string(inc.Type),
			inc.SeverityLevel,
			inc.PriorityLevel,
			inc.StatusCode,
			inc.DepartmentID,
			inc.LocationID,
			inc.OccurredAt.Format("2006-01-02 15:04:05"),
			inc.ReportedAt.Format("2006-01-02 15:04:05"),
		}
		if err := writer.Write(row); err != nil {
			return nil, fmt.Errorf("failed to write CSV row: %w", err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, fmt.Errorf("CSV flush error: %w", err)
	}

	return buf.Bytes(), nil
}

// ExportPDF generates a PDF byte stream for a single incident report.
// In production, integrate with a PDF rendering library (e.g., gofpdf).
func (s *Service) ExportPDF(ctx context.Context, incidentID string) ([]byte, error) {
	criteria := &searchDomain.Criteria{
		IncidentNumber: incidentID,
		PageSize:       1,
	}
	criteria.Normalize()

	incidents, _, err := s.repo.Search(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve incident for PDF: %w", err)
	}

	if len(incidents) == 0 {
		return nil, fmt.Errorf("incident not found: %s", incidentID)
	}

	inc := incidents[0]

	// Generate a structured text-based report (PDF rendering library integration point)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("INCIDENT REPORT\n"))
	buf.WriteString(fmt.Sprintf("===============\n\n"))
	buf.WriteString(fmt.Sprintf("Incident Number: %s\n", inc.IncidentNumber))
	buf.WriteString(fmt.Sprintf("Title: %s\n", inc.Title))
	buf.WriteString(fmt.Sprintf("Type: %s\n", inc.Type))
	buf.WriteString(fmt.Sprintf("Severity: %s\n", inc.SeverityLevel))
	buf.WriteString(fmt.Sprintf("Priority: %s\n", inc.PriorityLevel))
	buf.WriteString(fmt.Sprintf("Status: %s\n", inc.StatusCode))
	buf.WriteString(fmt.Sprintf("Description:\n%s\n\n", inc.Description))
	buf.WriteString(fmt.Sprintf("Occurred At: %s\n", inc.OccurredAt.Format("2006-01-02 15:04:05")))
	buf.WriteString(fmt.Sprintf("Reported At: %s\n", inc.ReportedAt.Format("2006-01-02 15:04:05")))

	return buf.Bytes(), nil
}
