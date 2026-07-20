package reporting

import (
	"context"
	"time"
)

// ComplianceCounter defines metrics counts ports.
type ComplianceCounter interface {
	CountTotal(ctx context.Context) (int, error)
	CountCompliant(ctx context.Context) (int, error)
	CountNonCompliant(ctx context.Context) (int, error)
}

// DashboardReport wraps process safety metrics.
type DashboardReport struct {
	TotalObligations int       `json:"total_obligations"`
	CompliantCount   int       `json:"compliant_count"`
	NonCompliant     int       `json:"non_compliant"`
	GeneratedAt      time.Time `json:"generated_at"`
}

// Service calculates reporting numbers.
type Service struct {
	counter ComplianceCounter
}

// NewService instantiates Reporting Service.
func NewService(counter ComplianceCounter) *Service {
	return &Service{counter: counter}
}

// GenerateDashboardReport aggregates real-time metrics.
func (s *Service) GenerateDashboardReport(ctx context.Context) (*DashboardReport, error) {
	total, _ := s.counter.CountTotal(ctx)
	comp, _ := s.counter.CountCompliant(ctx)
	non, _ := s.counter.CountNonCompliant(ctx)

	return &DashboardReport{
		TotalObligations: total,
		CompliantCount:   comp,
		NonCompliant:     non,
		GeneratedAt:      time.Now(),
	}, nil
}
