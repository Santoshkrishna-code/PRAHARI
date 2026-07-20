package export

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"

	complianceDomain "prahari/services/compliance/internal/domain/compliance"
	searchDomain "prahari/services/compliance/internal/domain/search"
)

// SearchRepository lists matches.
type SearchRepository interface {
	Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*complianceDomain.Compliance, int, error)
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
		return nil, fmt.Errorf("failed to retrieve compliance register records: %w", err)
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	header := []string{"Compliance Number", "Score", "Status", "Title"}
	_ = writer.Write(header)

	for _, c := range list {
		row := []string{
			c.ComplianceNumber,
			fmt.Sprintf("%.2f", c.ComplianceScore),
			c.StatusCode,
			c.Title,
		}
		_ = writer.Write(row)
	}

	writer.Flush()
	return buf.Bytes(), nil
}

// ExportPDF prints layout summary text bytes.
func (s *Service) ExportPDF(ctx context.Context, complianceID string) ([]byte, error) {
	criteria := &searchDomain.Criteria{
		ComplianceNumber: complianceID,
		PageSize:         1,
	}
	criteria.Normalize()

	list, _, err := s.repo.Search(ctx, criteria)
	if err != nil || len(list) == 0 {
		return nil, fmt.Errorf("failed to retrieve compliance record: %w", err)
	}

	c := list[0]

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("ENTERPRISE STATUTORY GOVERNANCE COMPLIANCE SHEET\n"))
	buf.WriteString(fmt.Sprintf("==================================================\n\n"))
	buf.WriteString(fmt.Sprintf("Compliance Number: %s\n", c.ComplianceNumber))
	buf.WriteString(fmt.Sprintf("Compliance Score: %.2f%%\n", c.ComplianceScore))
	buf.WriteString(fmt.Sprintf("Status: %s\n\n", c.StatusCode))
	buf.WriteString(fmt.Sprintf("Title: %s\n", c.Title))
	buf.WriteString(fmt.Sprintf("Description:\n%s\n", c.Description))

	return buf.Bytes(), nil
}
