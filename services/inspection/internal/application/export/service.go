package export

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"

	inspectionDomain "prahari/services/inspection/internal/domain/inspection"
	searchDomain "prahari/services/inspection/internal/domain/search"
)

// SearchRepository lists matches.
type SearchRepository interface {
	Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*inspectionDomain.Inspection, int, error)
}

// Service writes raw binary streams.
type Service struct {
	repo SearchRepository
}

// NewService instantiates Export Service.
func NewService(repo SearchRepository) *Service {
	return &Service{repo: repo}
}

// ExportCSV streams a CSV array.
func (s *Service) ExportCSV(ctx context.Context, criteria *searchDomain.Criteria) ([]byte, error) {
	criteria.Normalize()
	criteria.PageSize = 10000

	inspections, _, err := s.repo.Search(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve inspections for export: %w", err)
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	header := []string{"Inspection Number", "Title", "Type", "Status", "Inspector", "Department", "Score"}
	_ = writer.Write(header)

	for _, i := range inspections {
		row := []string{
			i.InspectionNumber,
			i.Title,
			string(i.InspectionType),
			i.StatusCode,
			i.InspectorID,
			i.DepartmentID,
			fmt.Sprintf("%.2f", i.ComplianceScore),
		}
		_ = writer.Write(row)
	}

	writer.Flush()
	return buf.Bytes(), nil
}

// ExportPDF prints layout summary text bytes.
func (s *Service) ExportPDF(ctx context.Context, inspectionID string) ([]byte, error) {
	criteria := &searchDomain.Criteria{
		InspectionNumber: inspectionID,
		PageSize:         1,
	}
	criteria.Normalize()

	inspections, _, err := s.repo.Search(ctx, criteria)
	if err != nil || len(inspections) == 0 {
		return nil, fmt.Errorf("failed to retrieve inspection: %w", err)
	}

	i := inspections[0]

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("SAFETY AUDIT INSPECTION SHEET\n"))
	buf.WriteString(fmt.Sprintf("=============================\n\n"))
	buf.WriteString(fmt.Sprintf("Inspection Number: %s\n", i.InspectionNumber))
	buf.WriteString(fmt.Sprintf("Title: %s\n", i.Title))
	buf.WriteString(fmt.Sprintf("Type: %s\n", i.InspectionType))
	buf.WriteString(fmt.Sprintf("Status: %s\n", i.StatusCode))
	buf.WriteString(fmt.Sprintf("Compliance Score: %.2f%%\n\n", i.ComplianceScore))
	buf.WriteString(fmt.Sprintf("Description:\n%s\n", i.Description))

	return buf.Bytes(), nil
}
