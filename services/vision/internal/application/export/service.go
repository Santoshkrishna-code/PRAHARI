package export

import (
	"context"
	"fmt"

	"prahari/services/vision/internal/domain/detection"
	"prahari/services/vision/internal/domain/search"
)

type Repository interface {
	SearchDetections(ctx context.Context, criteria *search.Criteria) ([]*detection.Detection, int64, error)
	GetDetectionByID(ctx context.Context, id string) (*detection.Detection, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportCSV(ctx context.Context, criteria *search.Criteria) ([]byte, error) {
	detections, _, err := s.repo.SearchDetections(ctx, criteria)
	if err != nil {
		return nil, err
	}

	csvData := "ID,JobID,Label,Confidence,Timestamp\n"
	for _, d := range detections {
		csvData += fmt.Sprintf("%s,%s,%s,%f,%s\n", d.ID, d.JobID, d.Label, d.Confidence, d.Timestamp)
	}

	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	d, err := s.repo.GetDetectionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock Detection Metadata\nID: %s\nJobID: %s\nLabel: %s\nConfidence: %f\n",
		d.ID, d.JobID, d.Label, d.Confidence)
	return []byte(pdfDoc), nil
}
