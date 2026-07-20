package reporting

import (
	"context"
	"time"
)

// ObservationCounter defines metrics aggregates ports.
type ObservationCounter interface {
	CountTotal(ctx context.Context) (int, error)
	CountSafe(ctx context.Context) (int, error)
	CountUnsafe(ctx context.Context) (int, error)
}

// DashboardReport wraps BBS dashboard metrics.
type DashboardReport struct {
	TotalObservations int       `json:"total_observations"`
	SafeCount         int       `json:"safe_count"`
	UnsafeCount       int       `json:"unsafe_count"`
	GeneratedAt       time.Time `json:"generated_at"`
}

// Service calculates reporting numbers.
type Service struct {
	counter ObservationCounter
}

// NewService instantiates Reporting Service.
func NewService(counter ObservationCounter) *Service {
	return &Service{counter: counter}
}

// GenerateDashboardReport aggregates real-time metrics.
func (s *Service) GenerateDashboardReport(ctx context.Context) (*DashboardReport, error) {
	total, _ := s.counter.CountTotal(ctx)
	safe, _ := s.counter.CountSafe(ctx)
	unsafe, _ := s.counter.CountUnsafe(ctx)

	return &DashboardReport{
		TotalObservations: total,
		SafeCount:         safe,
		UnsafeCount:       unsafe,
		GeneratedAt:       time.Now(),
	}, nil
}
