package reporting

import (
	"context"
	"fmt"
	"time"
)

// IncidentCounter defines the persistence port for aggregate reporting queries.
type IncidentCounter interface {
	CountByCategory(ctx context.Context, from, to time.Time) (map[string]int, error)
	CountBySeverity(ctx context.Context, from, to time.Time) (map[string]int, error)
	CountByStatus(ctx context.Context, from, to time.Time) (map[string]int, error)
	CountByDepartment(ctx context.Context, from, to time.Time) (map[string]int, error)
	AverageResolutionHours(ctx context.Context, from, to time.Time) (float64, error)
	AverageInvestigationHours(ctx context.Context, from, to time.Time) (float64, error)
	CAPACompletionRate(ctx context.Context, from, to time.Time) (float64, error)
}

// SummaryReport represents an aggregated summary of incident data over a time range.
type SummaryReport struct {
	DateFrom     time.Time      `json:"date_from"`
	DateTo       time.Time      `json:"date_to"`
	ByCategory   map[string]int `json:"by_category"`
	BySeverity   map[string]int `json:"by_severity"`
	ByStatus     map[string]int `json:"by_status"`
	ByDepartment map[string]int `json:"by_department"`
	GeneratedAt  time.Time      `json:"generated_at"`
}

// MetricsReport represents calculated performance metrics over a time range.
type MetricsReport struct {
	DateFrom                 time.Time `json:"date_from"`
	DateTo                   time.Time `json:"date_to"`
	AverageResolutionHours   float64   `json:"average_resolution_hours"`
	AverageInvestigationHours float64  `json:"average_investigation_hours"`
	CAPACompletionRate       float64   `json:"capa_completion_rate"`
	GeneratedAt              time.Time `json:"generated_at"`
}

// Service orchestrates reporting operations with aggregate queries.
type Service struct {
	counter IncidentCounter
}

// NewService constructs a Service with the reporting repository injected.
func NewService(counter IncidentCounter) *Service {
	return &Service{counter: counter}
}

// GenerateSummaryReport aggregates incident counts across multiple dimensions.
func (s *Service) GenerateSummaryReport(ctx context.Context, from, to time.Time) (*SummaryReport, error) {
	byCategory, err := s.counter.CountByCategory(ctx, from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to count by category: %w", err)
	}

	bySeverity, err := s.counter.CountBySeverity(ctx, from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to count by severity: %w", err)
	}

	byStatus, err := s.counter.CountByStatus(ctx, from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to count by status: %w", err)
	}

	byDept, err := s.counter.CountByDepartment(ctx, from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to count by department: %w", err)
	}

	return &SummaryReport{
		DateFrom:     from,
		DateTo:       to,
		ByCategory:   byCategory,
		BySeverity:   bySeverity,
		ByStatus:     byStatus,
		ByDepartment: byDept,
		GeneratedAt:  time.Now(),
	}, nil
}

// GenerateMetricsReport calculates performance metrics over a time range.
func (s *Service) GenerateMetricsReport(ctx context.Context, from, to time.Time) (*MetricsReport, error) {
	avgResolution, err := s.counter.AverageResolutionHours(ctx, from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate avg resolution: %w", err)
	}

	avgInvestigation, err := s.counter.AverageInvestigationHours(ctx, from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate avg investigation: %w", err)
	}

	capaRate, err := s.counter.CAPACompletionRate(ctx, from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate CAPA rate: %w", err)
	}

	return &MetricsReport{
		DateFrom:                  from,
		DateTo:                    to,
		AverageResolutionHours:    avgResolution,
		AverageInvestigationHours: avgInvestigation,
		CAPACompletionRate:        capaRate,
		GeneratedAt:               time.Now(),
	}, nil
}
