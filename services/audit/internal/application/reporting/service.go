package reporting

import (
	"context"
	"time"
)

// AuditCounter defines metrics counts ports.
type AuditCounter interface {
	CountTotal(ctx context.Context) (int, error)
	CountOpen(ctx context.Context) (int, error)
	CountClosed(ctx context.Context) (int, error)
}

// DashboardReport wraps EHS KPIs.
type DashboardReport struct {
	TotalAudits int       `json:"total_audits"`
	OpenCount   int       `json:"open_count"`
	ClosedCount int       `json:"closed_count"`
	GeneratedAt time.Time `json:"generated_at"`
}

// Service calculates reporting numbers.
type Service struct {
	counter AuditCounter
}

// NewService instantiates Reporting Service.
func NewService(counter AuditCounter) *Service {
	return &Service{counter: counter}
}

// GenerateDashboardReport aggregates real-time metrics.
func (s *Service) GenerateDashboardReport(ctx context.Context) (*DashboardReport, error) {
	total, _ := s.counter.CountTotal(ctx)
	open, _ := s.counter.CountOpen(ctx)
	closed, _ := s.counter.CountClosed(ctx)

	return &DashboardReport{
		TotalAudits: total,
		OpenCount:   open,
		ClosedCount: closed,
		GeneratedAt: time.Now(),
	}, nil
}
