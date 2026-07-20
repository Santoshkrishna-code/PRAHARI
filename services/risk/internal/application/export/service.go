package export

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"

	riskDomain "prahari/services/risk/internal/domain/risk"
	searchDomain "prahari/services/risk/internal/domain/search"
)

// SearchRepository lists matches.
type SearchRepository interface {
	Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*riskDomain.Risk, int, error)
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
		return nil, fmt.Errorf("failed to retrieve risk register records: %w", err)
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	header := []string{"Risk Number", "Inherent Score", "Status", "Title"}
	_ = writer.Write(header)

	for _, r := range list {
		row := []string{
			r.RiskNumber,
			fmt.Sprintf("%d", r.InherentRiskScore),
			r.StatusCode,
			r.Title,
		}
		_ = writer.Write(row)
	}

	writer.Flush()
	return buf.Bytes(), nil
}

// ExportPDF prints layout summary text bytes.
func (s *Service) ExportPDF(ctx context.Context, riskID string) ([]byte, error) {
	criteria := &searchDomain.Criteria{
		RiskNumber: riskID,
		PageSize:   1,
	}
	criteria.Normalize()

	list, _, err := s.repo.Search(ctx, criteria)
	if err != nil || len(list) == 0 {
		return nil, fmt.Errorf("failed to retrieve risk record: %w", err)
	}

	r := list[0]

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("OPERATIONAL PROCESS SAFETY RISK ASSESSMENT SHEET\n"))
	buf.WriteString(fmt.Sprintf("==================================================\n\n"))
	buf.WriteString(fmt.Sprintf("Risk Number: %s\n", r.RiskNumber))
	buf.WriteString(fmt.Sprintf("Inherent Risk Score: %d\n", r.InherentRiskScore))
	buf.WriteString(fmt.Sprintf("Status: %s\n\n", r.StatusCode))
	buf.WriteString(fmt.Sprintf("Title: %s\n", r.Title))
	buf.WriteString(fmt.Sprintf("Description:\n%s\n", r.Description))

	return buf.Bytes(), nil
}
