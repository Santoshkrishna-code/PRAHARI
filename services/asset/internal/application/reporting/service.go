package reporting

import (
	"context"
	"time"
)

// AssetCounter defines metric aggregates counts ports.
type AssetCounter interface {
	CountTotal(ctx context.Context) (int, error)
	CountCritical(ctx context.Context) (int, error)
	CountOperational(ctx context.Context) (int, error)
	CountInMaintenance(ctx context.Context) (int, error)
	AverageHealthScore(ctx context.Context) (float64, error)
	AverageAge(ctx context.Context) (float64, error)
}

// DashboardReport wraps standard KPI summaries.
type DashboardReport struct {
	TotalAssets            int       `json:"total_assets"`
	CriticalAssets         int       `json:"critical_assets"`
	OperationalAssets      int       `json:"operational_assets"`
	AssetsInMaintenance    int       `json:"assets_in_maintenance"`
	AverageHealthScore     float64   `json:"average_health_score"`
	AverageAgeYears        float64   `json:"average_age_years"`
	GeneratedAt            time.Time `json:"generated_at"`
}

// Service calculates reporting numbers.
type Service struct {
	counter AssetCounter
}

// NewService instantiates Reporting Service.
func NewService(counter AssetCounter) *Service {
	return &Service{counter: counter}
}

// GenerateDashboardReport aggregates real-time metrics.
func (s *Service) GenerateDashboardReport(ctx context.Context) (*DashboardReport, error) {
	total, _ := s.counter.CountTotal(ctx)
	critical, _ := s.counter.CountCritical(ctx)
	operational, _ := s.counter.CountOperational(ctx)
	maintenance, _ := s.counter.CountInMaintenance(ctx)
	health, _ := s.counter.AverageHealthScore(ctx)
	age, _ := s.counter.AverageAge(ctx)

	return &DashboardReport{
		TotalAssets:         total,
		CriticalAssets:      critical,
		OperationalAssets:   operational,
		AssetsInMaintenance: maintenance,
		AverageHealthScore:  health,
		AverageAgeYears:     age,
		GeneratedAt:         time.Now(),
	}, nil
}
