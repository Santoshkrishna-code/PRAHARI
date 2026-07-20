package reporting

import (
	"context"
	"time"
)

// TrainingCounter defines metrics counts ports.
type TrainingCounter interface {
	CountTotal(ctx context.Context) (int, error)
	CountCertified(ctx context.Context) (int, error)
}

// DashboardReport wraps EHS competency KPIs.
type DashboardReport struct {
	TotalTrainings int       `json:"total_trainings"`
	CertifiedCount int       `json:"certified_count"`
	GeneratedAt    time.Time `json:"generated_at"`
}

// Service calculates reporting numbers.
type Service struct {
	counter TrainingCounter
}

// NewService instantiates Reporting Service.
func NewService(counter TrainingCounter) *Service {
	return &Service{counter: counter}
}

// GenerateDashboardReport aggregates real-time metrics.
func (s *Service) GenerateDashboardReport(ctx context.Context) (*DashboardReport, error) {
	total, _ := s.counter.CountTotal(ctx)
	cert, _ := s.counter.CountCertified(ctx)

	return &DashboardReport{
		TotalTrainings: total,
		CertifiedCount: cert,
		GeneratedAt:    time.Now(),
	}, nil
}
