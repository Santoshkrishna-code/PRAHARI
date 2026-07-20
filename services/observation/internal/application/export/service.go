package export

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"

	observationDomain "prahari/services/observation/internal/domain/observation"
	searchDomain "prahari/services/observation/internal/domain/search"
)

// SearchRepository lists matches.
type SearchRepository interface {
	Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*observationDomain.Observation, int, error)
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
		return nil, fmt.Errorf("failed to retrieve safety observation records: %w", err)
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	header := []string{"Observation Number", "Type", "Status", "Title"}
	_ = writer.Write(header)

	for _, o := range list {
		row := []string{
			o.ObservationNumber,
			o.ObservationType,
			o.StatusCode,
			o.Title,
		}
		_ = writer.Write(row)
	}

	writer.Flush()
	return buf.Bytes(), nil
}

// ExportPDF prints layout summary text bytes.
func (s *Service) ExportPDF(ctx context.Context, observationID string) ([]byte, error) {
	criteria := &searchDomain.Criteria{
		ObservationNumber: observationID,
		PageSize:          1,
	}
	criteria.Normalize()

	list, _, err := s.repo.Search(ctx, criteria)
	if err != nil || len(list) == 0 {
		return nil, fmt.Errorf("failed to retrieve observation record: %w", err)
	}

	o := list[0]

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("BEHAVIOR-BASED SAFETY OBSERVATION SHEET\n"))
	buf.WriteString(fmt.Sprintf("=========================================\n\n"))
	buf.WriteString(fmt.Sprintf("Observation Number: %s\n", o.ObservationNumber))
	buf.WriteString(fmt.Sprintf("Observation Type: %s\n", o.ObservationType))
	buf.WriteString(fmt.Sprintf("Status: %s\n\n", o.StatusCode))
	buf.WriteString(fmt.Sprintf("Title: %s\n", o.Title))
	buf.WriteString(fmt.Sprintf("Description:\n%s\n", o.Description))

	return buf.Bytes(), nil
}
