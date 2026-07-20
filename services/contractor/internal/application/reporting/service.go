package reporting

import (
	"context"
	"time"
)

// ContractorCounter defines metrics aggregates ports.
type ContractorCounter interface {
	CountTotal(ctx context.Context) (int, error)
	CountActive(ctx context.Context) (int, error)
	CountComplianceViolations(ctx context.Context) (int, error)
}

// DashboardReport wraps standard KPI summaries.
type DashboardReport struct {
	TotalContractors     int       `json:"total_contractors"`
	ActiveContractors     int       `json:"active_contractors"`
	ComplianceViolations int       `json:"compliance_violations"`
	GeneratedAt          time.Time `json:"generated_at"`
}

// Service calculates reporting numbers.
type Service struct {
	counter ContractorCounter
}

// NewService instantiates Reporting Service.
func NewService(counter ContractorCounter) *Service {
	return &Service{counter: counter}
}

// GenerateDashboardReport aggregates real-time metrics.
func (s *Service) GenerateDashboardReport(ctx context.Context) (*DashboardReport, error) {
	total, _ := s.counter.CountTotal(ctx)
	active, _ := s.counter.CountActive(ctx)
	violations, _ := s.counter.CountComplianceViolations(ctx)

	return &DashboardReport{
		TotalContractors:     total,
		ActiveContractors:     active,
		ComplianceViolations: violations,
		GeneratedAt:          time.Now(),
	}, nil
}
