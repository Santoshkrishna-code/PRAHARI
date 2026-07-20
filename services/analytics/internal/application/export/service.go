package export

import (
	"context"
	"fmt"

	"prahari/services/analytics/internal/domain/metric"
	"prahari/services/analytics/internal/domain/search"
)

type Repository interface {
	SearchMetrics(ctx context.Context, criteria *search.Criteria) ([]*metric.Metric, int64, error)
	GetMetricByID(ctx context.Context, id string) (*metric.Metric, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportCSV(ctx context.Context, criteria *search.Criteria) ([]byte, error) {
	metrics, _, err := s.repo.SearchMetrics(ctx, criteria)
	if err != nil {
		return nil, err
	}

	csvData := "ID,PlantID,MetricKey,Value,Timestamp\n"
	for _, m := range metrics {
		csvData += fmt.Sprintf("%s,%s,%s,%f,%s\n", m.ID, m.PlantID, m.Key, m.Value, m.Timestamp)
	}

	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	m, err := s.repo.GetMetricByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock Metric Summary\nID: %s\nPlantID: %s\nMetricKey: %s\nValue: %f\n",
		m.ID, m.PlantID, m.Key, m.Value)
	return []byte(pdfDoc), nil
}
