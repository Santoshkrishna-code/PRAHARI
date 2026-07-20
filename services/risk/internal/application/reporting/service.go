package reporting

import (
	"context"
	"time"
)

// RiskCounter defines metrics counts ports.
type RiskCounter interface {
	CountTotal(ctx context.Context) (int, error)
	CountHigh(ctx context.Context) (int, error)
	CountCritical(ctx context.Context) (int, error)
}

// DashboardReport wraps process safety metrics.
type DashboardReport struct {
	TotalRisks    int       `json:"total_risks"`
	HighCount     int       `json:"high_count"`
	CriticalCount int       `json:"critical_count"`
	GeneratedAt   time.Time `json:"generated_at"`
}

// Service calculates reporting numbers.
type Service struct {
	counter RiskCounter
}

// NewService instantiates Reporting Service.
func NewService(counter RiskCounter) *Service {
	return &Service{counter: counter}
}

// GenerateDashboardReport aggregates real-time metrics.
func (s *Service) GenerateDashboardReport(ctx context.Context) (*DashboardReport, error) {
	total, _ := s.counter.CountTotal(ctx)
	high, _ := s.counter.CountHigh(ctx)
	crit, _ := s.counter.CountCritical(ctx)

	return &DashboardReport{
		TotalRisks:    total,
		HighCount:     high,
		CriticalCount: crit,
		GeneratedAt:   time.Now(),
	}, nil
}
