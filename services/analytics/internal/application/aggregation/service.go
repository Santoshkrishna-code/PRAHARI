package aggregation

import (
	"context"
	"fmt"
	"time"

	"prahari/services/analytics/internal/domain/metric"
)

type Repository interface {
	SaveMetric(ctx context.Context, m *metric.Metric) error
	GetMetricValue(ctx context.Context, plantID, key string) (float64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AggregateEvent(ctx context.Context, plantID, key string, delta float64) error {
	current, _ := s.repo.GetMetricValue(ctx, plantID, key)
	newValue := current + delta

	m := &metric.Metric{
		ID:        fmt.Sprintf("m-%d", time.Now().UnixNano()),
		PlantID:   plantID,
		Key:       key,
		Value:     newValue,
		Timestamp: time.Now(),
	}

	return s.repo.SaveMetric(ctx, m)
}
