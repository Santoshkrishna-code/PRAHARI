package reporting

import (
	"context"
	"time"
)

// HazardCounter defines metrics aggregates ports.
type HazardCounter interface {
	CountTotal(ctx context.Context) (int, error)
	CountOpen(ctx context.Context) (int, error)
	CountHighRisk(ctx context.Context) (int, error)
}

// DashboardReport wraps safety dashboard metrics.
type DashboardReport struct {
	TotalHazards     int       `json:"total_hazards"`
	OpenHazards      int       `json:"open_hazards"`
	HighRiskHazards  int       `json:"high_risk_hazards"`
	GeneratedAt      time.Time `json:"generated_at"`
}

// Service calculates reporting numbers.
type Service struct {
	counter HazardCounter
}

// NewService instantiates Reporting Service.
func NewService(counter HazardCounter) *Service {
	return &Service{counter: counter}
}

// GenerateDashboardReport aggregates real-time metrics.
func (s *Service) GenerateDashboardReport(ctx context.Context) (*DashboardReport, error) {
	total, _ := s.counter.CountTotal(ctx)
	open, _ := s.counter.CountOpen(ctx)
	highRisk, _ := s.counter.CountHighRisk(ctx)

	return &DashboardReport{
		TotalHazards:    total,
		OpenHazards:     open,
		HighRiskHazards: highRisk,
		GeneratedAt:     time.Now(),
	}, nil
}
