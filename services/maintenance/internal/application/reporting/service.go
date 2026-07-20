package reporting

import (
	"context"
	"time"
)

// MaintenanceCounter defines metric aggregates counts ports.
type MaintenanceCounter interface {
	CountTotal(ctx context.Context) (int, error)
	CountOpen(ctx context.Context) (int, error)
	AverageDowntime(ctx context.Context) (float64, error)
	MTTR(ctx context.Context) (float64, error)
	MTBF(ctx context.Context) (float64, error)
}

// DashboardReport wraps standard KPI summaries.
type DashboardReport struct {
	TotalWorkOrders     int       `json:"total_work_orders"`
	OpenWorkOrders      int       `json:"open_work_orders"`
	AverageDowntimeHrs  float64   `json:"average_downtime_hrs"`
	MeanTimeToRepairHrs float64   `json:"mean_time_to_repair_hrs"`
	MeanTimeBetweenFailuresHrs float64 `json:"mean_time_between_failures_hrs"`
	GeneratedAt         time.Time `json:"generated_at"`
}

// Service calculates reporting numbers.
type Service struct {
	counter MaintenanceCounter
}

// NewService instantiates Reporting Service.
func NewService(counter MaintenanceCounter) *Service {
	return &Service{counter: counter}
}

// GenerateDashboardReport aggregates real-time metrics.
func (s *Service) GenerateDashboardReport(ctx context.Context) (*DashboardReport, error) {
	total, _ := s.counter.CountTotal(ctx)
	open, _ := s.counter.CountOpen(ctx)
	downtime, _ := s.counter.AverageDowntime(ctx)
	mttr, _ := s.counter.MTTR(ctx)
	mtbf, _ := s.counter.MTBF(ctx)

	return &DashboardReport{
		TotalWorkOrders:            total,
		OpenWorkOrders:             open,
		AverageDowntimeHrs:         downtime,
		MeanTimeToRepairHrs:        mttr,
		MeanTimeBetweenFailuresHrs: mtbf,
		GeneratedAt:                time.Now(),
	}, nil
}
