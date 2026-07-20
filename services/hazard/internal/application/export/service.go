package export

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"

	hazardDomain "prahari/services/hazard/internal/domain/hazard"
	searchDomain "prahari/services/hazard/internal/domain/search"
)

// SearchRepository lists matches.
type SearchRepository interface {
	Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*hazardDomain.Hazard, int, error)
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
		return nil, fmt.Errorf("failed to retrieve hazard records for export: %w", err)
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	header := []string{"Hazard Number", "Type", "Initial Risk", "Residual Risk", "Status", "Title"}
	_ = writer.Write(header)

	for _, h := range list {
		row := []string{
			h.HazardNumber,
			h.HazardType,
			fmt.Sprintf("%d", h.InitialRiskScore),
			fmt.Sprintf("%d", h.ResidualRiskScore),
			h.StatusCode,
			h.Title,
		}
		_ = writer.Write(row)
	}

	writer.Flush()
	return buf.Bytes(), nil
}

// ExportPDF prints layout summary text bytes.
func (s *Service) ExportPDF(ctx context.Context, hazardID string) ([]byte, error) {
	criteria := &searchDomain.Criteria{
		HazardNumber: hazardID,
		PageSize:     1,
	}
	criteria.Normalize()

	list, _, err := s.repo.Search(ctx, criteria)
	if err != nil || len(list) == 0 {
		return nil, fmt.Errorf("failed to retrieve hazard record: %w", err)
	}

	h := list[0]

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("HAZARD IDENTIFICATION & RISK ASSESSMENT DATA SHEET\n"))
	buf.WriteString(fmt.Sprintf("==================================================\n\n"))
	buf.WriteString(fmt.Sprintf("Hazard Number: %s\n", h.HazardNumber))
	buf.WriteString(fmt.Sprintf("Type: %s\n", h.HazardType))
	buf.WriteString(fmt.Sprintf("Initial Risk Score: %d\n", h.InitialRiskScore))
	buf.WriteString(fmt.Sprintf("Residual Risk Score: %d\n", h.ResidualRiskScore))
	buf.WriteString(fmt.Sprintf("Status: %s\n\n", h.StatusCode))
	buf.WriteString(fmt.Sprintf("Title: %s\n", h.Title))
	buf.WriteString(fmt.Sprintf("Description:\n%s\n", h.Description))

	return buf.Bytes(), nil
}
