package export

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"

	contractorDomain "prahari/services/contractor/internal/domain/contractor"
	searchDomain "prahari/services/contractor/internal/domain/search"
)

// SearchRepository lists matches.
type SearchRepository interface {
	Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*contractorDomain.Contractor, int, error)
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
		return nil, fmt.Errorf("failed to retrieve contractor records for export: %w", err)
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	header := []string{"Contractor Number", "Company Name", "Tax ID", "Status", "Insurance Expiry"}
	_ = writer.Write(header)

	for _, c := range list {
		row := []string{
			c.ContractorNumber,
			c.CompanyName,
			c.TaxID,
			c.StatusCode,
			c.InsuranceExpiry.Format("2006-01-02"),
		}
		_ = writer.Write(row)
	}

	writer.Flush()
	return buf.Bytes(), nil
}

// ExportPDF prints layout summary text bytes.
func (s *Service) ExportPDF(ctx context.Context, contractorID string) ([]byte, error) {
	criteria := &searchDomain.Criteria{
		ContractorNumber: contractorID,
		PageSize:         1,
	}
	criteria.Normalize()

	list, _, err := s.repo.Search(ctx, criteria)
	if err != nil || len(list) == 0 {
		return nil, fmt.Errorf("failed to retrieve contractor record: %w", err)
	}

	c := list[0]

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("CONTRACTOR PROFILE DATA SHEET\n"))
	buf.WriteString(fmt.Sprintf("===============================\n\n"))
	buf.WriteString(fmt.Sprintf("Contractor Number: %s\n", c.ContractorNumber))
	buf.WriteString(fmt.Sprintf("Company Name: %s\n", c.CompanyName))
	buf.WriteString(fmt.Sprintf("Tax ID: %s\n", c.TaxID))
	buf.WriteString(fmt.Sprintf("Status Code: %s\n", c.StatusCode))
	buf.WriteString(fmt.Sprintf("Insurance Expiry: %s\n", c.InsuranceExpiry.Format("2006-01-02")))

	return buf.Bytes(), nil
}
