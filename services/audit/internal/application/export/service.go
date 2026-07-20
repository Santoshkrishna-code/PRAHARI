package export

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"

	auditDomain "prahari/services/audit/internal/domain/audit"
	searchDomain "prahari/services/audit/internal/domain/search"
)

// SearchRepository lists matches.
type SearchRepository interface {
	Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*auditDomain.Audit, int, error)
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

	list, _, err := s.repo.Search(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve audit records: %w", err)
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	header := []string{"Audit Number", "Rating", "Status", "Title"}
	_ = writer.Write(header)

	for _, a := range list {
		row := []string{
			a.AuditNumber,
			fmt.Sprintf("%.2f", a.ComplianceRating),
			a.StatusCode,
			a.Title,
		}
		_ = writer.Write(row)
	}

	writer.Flush()
	return buf.Bytes(), nil
}

// ExportPDF prints layout summary text bytes.
func (s *Service) ExportPDF(ctx context.Context, auditID string) ([]byte, error) {
	criteria := &searchDomain.Criteria{
		AuditNumber: auditID,
		PageSize:    1,
	}
	criteria.Normalize()

	list, _, err := s.repo.Search(ctx, criteria)
	if err != nil || len(list) == 0 {
		return nil, fmt.Errorf("failed to retrieve audit record: %w", err)
	}

	a := list[0]

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("ENTERPRISE GOVERNANCE ASSURANCE AUDIT REPORT\n"))
	buf.WriteString(fmt.Sprintf("==================================================\n\n"))
	buf.WriteString(fmt.Sprintf("Audit Number: %s\n", a.AuditNumber))
	buf.WriteString(fmt.Sprintf("Compliance Rating: %.2f%%\n", a.ComplianceRating))
	buf.WriteString(fmt.Sprintf("Status: %s\n\n", a.StatusCode))
	buf.WriteString(fmt.Sprintf("Title: %s\n", a.Title))
	buf.WriteString(fmt.Sprintf("Description:\n%s\n", a.Description))

	return buf.Bytes(), nil
}
