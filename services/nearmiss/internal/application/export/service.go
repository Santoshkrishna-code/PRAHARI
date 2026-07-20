package export

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"

	nearmissDomain "prahari/services/nearmiss/internal/domain/nearmiss"
	searchDomain "prahari/services/nearmiss/internal/domain/search"
)

// SearchRepository lists matches.
type SearchRepository interface {
	Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*nearmissDomain.NearMiss, int, error)
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
		return nil, fmt.Errorf("failed to retrieve near miss records: %w", err)
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	header := []string{"Near Miss Number", "Classification", "Severity", "Status", "Title"}
	_ = writer.Write(header)

	for _, nm := range list {
		row := []string{
			nm.NearMissNumber,
			nm.Classification,
			nm.SeverityLevel,
			nm.StatusCode,
			nm.Title,
		}
		_ = writer.Write(row)
	}

	writer.Flush()
	return buf.Bytes(), nil
}

// ExportPDF prints layout summary text bytes.
func (s *Service) ExportPDF(ctx context.Context, nearmissID string) ([]byte, error) {
	criteria := &searchDomain.Criteria{
		NearMissNumber: nearmissID,
		PageSize:       1,
	}
	criteria.Normalize()

	list, _, err := s.repo.Search(ctx, criteria)
	if err != nil || len(list) == 0 {
		return nil, fmt.Errorf("failed to retrieve near miss record: %w", err)
	}

	nm := list[0]

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("NEAR MISS SAFETY INCIDENT REPORT\n"))
	buf.WriteString(fmt.Sprintf("=================================\n\n"))
	buf.WriteString(fmt.Sprintf("Near Miss Number: %s\n", nm.NearMissNumber))
	buf.WriteString(fmt.Sprintf("Classification: %s\n", nm.Classification))
	buf.WriteString(fmt.Sprintf("Severity Level: %s\n", nm.SeverityLevel))
	buf.WriteString(fmt.Sprintf("Status: %s\n\n", nm.StatusCode))
	buf.WriteString(fmt.Sprintf("Title: %s\n", nm.Title))
	buf.WriteString(fmt.Sprintf("Description:\n%s\n", nm.Description))

	return buf.Bytes(), nil
}
