package reporting

import (
	"context"
	"time"
)

// PermitCounter defines the reporting metrics port.
type PermitCounter interface {
	CountActive(ctx context.Context) (int, error)
	CountExpired(ctx context.Context) (int, error)
	CountOpen(ctx context.Context) (int, error)
	AverageApprovalTimeHours(ctx context.Context) (float64, error)
	AverageCompletionTimeHours(ctx context.Context) (float64, error)
	ComplianceRate(ctx context.Context) (float64, error)
}

// MetricsReport wraps standard PTW stats.
type MetricsReport struct {
	ActivePermits              int       `json:"active_permits"`
	ExpiredPermits             int       `json:"expired_permits"`
	OpenPermits                int       `json:"open_permits"`
	AverageApprovalTimeHours   float64   `json:"average_approval_time_hours"`
	AverageCompletionTimeHours float64   `json:"average_completion_time_hours"`
	ComplianceRate             float64   `json:"compliance_rate"`
	GeneratedAt                time.Time `json:"generated_at"`
}

// Service calculates metrics reports.
type Service struct {
	counter PermitCounter
}

// NewService instantiates a Reporting Service.
func NewService(counter PermitCounter) *Service {
	return &Service{counter: counter}
}

// GenerateMetricsReport aggregates all real-time counters.
func (s *Service) GenerateMetricsReport(ctx context.Context) (*MetricsReport, error) {
	active, _ := s.counter.CountActive(ctx)
	expired, _ := s.counter.CountExpired(ctx)
	open, _ := s.counter.CountOpen(ctx)
	avgApp, _ := s.counter.AverageApprovalTimeHours(ctx)
	avgComp, _ := s.counter.AverageCompletionTimeHours(ctx)
	compliance, _ := s.counter.ComplianceRate(ctx)

	return &MetricsReport{
		ActivePermits:              active,
		ExpiredPermits:             expired,
		OpenPermits:                open,
		AverageApprovalTimeHours:   avgApp,
		AverageCompletionTimeHours: avgComp,
		ComplianceRate:             compliance,
		GeneratedAt:                time.Now(),
	}, nil
}
