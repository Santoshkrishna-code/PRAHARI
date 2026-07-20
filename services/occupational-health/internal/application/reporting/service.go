package reporting

import (
	"context"
)

// Repository retrieves analytics summaries.
type Repository interface {
	GetWorkforceHealthMetrics(ctx context.Context) (map[string]interface{}, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetDashboardReport(ctx context.Context) (map[string]interface{}, error) {
	return s.repo.GetWorkforceHealthMetrics(ctx)
}
