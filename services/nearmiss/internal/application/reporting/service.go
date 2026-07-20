package reporting

import (
	"context"
	"time"
)

// NearMissCounter defines metrics aggregates ports.
type NearMissCounter interface {
	CountTotal(ctx context.Context) (int, error)
	CountOpen(ctx context.Context) (int, error)
	CountEscalated(ctx context.Context) (int, error)
}

// DashboardReport wraps safety dashboard metrics.
type DashboardReport struct {
	TotalNearMisses  int       `json:"total_near_misses"`
	OpenNearMisses   int       `json:"open_near_misses"`
	EscalatedCount   int       `json:"escalated_count"`
	GeneratedAt      time.Time `json:"generated_at"`
}

// Service calculates reporting numbers.
type Service struct {
	counter NearMissCounter
}

// NewService instantiates Reporting Service.
func NewService(counter NearMissCounter) *Service {
	return &Service{counter: counter}
}

// GenerateDashboardReport aggregates real-time metrics.
func (s *Service) GenerateDashboardReport(ctx context.Context) (*DashboardReport, error) {
	total, _ := s.counter.CountTotal(ctx)
	open, _ := s.counter.CountOpen(ctx)
	escalated, _ := s.counter.CountEscalated(ctx)

	return &DashboardReport{
		TotalNearMisses: total,
		OpenNearMisses:  open,
		EscalatedCount:  escalated,
		GeneratedAt:     time.Now(),
	}, nil
}
