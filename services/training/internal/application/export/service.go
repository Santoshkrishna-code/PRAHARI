package export

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"

	trainingDomain "prahari/services/training/internal/domain/training"
	searchDomain "prahari/services/training/internal/domain/search"
)

// SearchRepository lists matches.
type SearchRepository interface {
	Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*trainingDomain.Training, int, error)
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
		return nil, fmt.Errorf("failed to retrieve training records: %w", err)
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	header := []string{"Training Number", "Status", "Title"}
	_ = writer.Write(header)

	for _, t := range list {
		row := []string{
			t.TrainingNumber,
			t.StatusCode,
			t.Title,
		}
		_ = writer.Write(row)
	}

	writer.Flush()
	return buf.Bytes(), nil
}

// ExportPDF prints layout summary text bytes.
func (s *Service) ExportPDF(ctx context.Context, trainingID string) ([]byte, error) {
	criteria := &searchDomain.Criteria{
		TrainingNumber: trainingID,
		PageSize:       1,
	}
	criteria.Normalize()

	list, _, err := s.repo.Search(ctx, criteria)
	if err != nil || len(list) == 0 {
		return nil, fmt.Errorf("failed to retrieve training record: %w", err)
	}

	t := list[0]

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("ENTERPRISE WORKFORCE COMPETENCY TRAINING REPORT\n"))
	buf.WriteString(fmt.Sprintf("==================================================\n\n"))
	buf.WriteString(fmt.Sprintf("Training Number: %s\n", t.TrainingNumber))
	buf.WriteString(fmt.Sprintf("Status: %s\n\n", t.StatusCode))
	buf.WriteString(fmt.Sprintf("Title: %s\n", t.Title))
	buf.WriteString(fmt.Sprintf("Description:\n%s\n", t.Description))

	return buf.Bytes(), nil
}
