package reporting

import (
	"context"
	"time"
)

// InspectionCounter defines metric aggregates counts ports.
type InspectionCounter interface {
	CountTotal(ctx context.Context) (int, error)
	CountOverdue(ctx context.Context) (int, error)
	CountUpcoming(ctx context.Context) (int, error)
	CountOpenNCRs(ctx context.Context) (int, error)
	CountCriticalFindings(ctx context.Context) (int, error)
	AverageComplianceScore(ctx context.Context) (float64, error)
	CAPACompletionRate(ctx context.Context) (float64, error)
}

// MetricsReport wraps standard KPI summaries.
type MetricsReport struct {
	TotalInspections       int       `json:"total_inspections"`
	OverdueInspections     int       `json:"overdue_inspections"`
	UpcomingInspections    int       `json:"upcoming_inspections"`
	OpenNCRs               int       `json:"open_ncrs"`
	CriticalFindings       int       `json:"critical_findings"`
	AverageComplianceScore float64   `json:"average_compliance_score"`
	CAPACompletionRate     float64   `json:"capa_completion_rate"`
	GeneratedAt            time.Time `json:"generated_at"`
}

// Service calculates reporting numbers.
type Service struct {
	counter InspectionCounter
}

// NewService instantiates Reporting Service.
func NewService(counter InspectionCounter) *Service {
	return &Service{counter: counter}
}

// GenerateMetricsReport aggregates real-time metrics.
func (s *Service) GenerateMetricsReport(ctx context.Context) (*MetricsReport, error) {
	total, _ := s.counter.CountTotal(ctx)
	overdue, _ := s.counter.CountOverdue(ctx)
	upcoming, _ := s.counter.CountUpcoming(ctx)
	ncrs, _ := s.counter.CountOpenNCRs(ctx)
	critical, _ := s.counter.CountCriticalFindings(ctx)
	compliance, _ := s.counter.AverageComplianceScore(ctx)
	capaRate, _ := s.counter.CAPACompletionRate(ctx)

	return &MetricsReport{
		TotalInspections:       total,
		OverdueInspections:     overdue,
		UpcomingInspections:    upcoming,
		OpenNCRs               ncrs,
		CriticalFindings:       critical,
		AverageComplianceScore: compliance,
		CAPACompletionRate:     capaRate,
		GeneratedAt:            time.Now(),
	}, nil
}
