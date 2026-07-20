package grpc

import (
	"context"

	analyticsApp "prahari/services/analytics/internal/application/analytics"
	dashboardsApp "prahari/services/analytics/internal/application/dashboards"
	reportingApp "prahari/services/analytics/internal/application/reporting"
	searchApp "prahari/services/analytics/internal/application/search"
	"prahari/services/analytics/internal/domain/dashboard"
	"prahari/services/analytics/internal/domain/kpi"
	"prahari/services/analytics/internal/domain/metric"
	"prahari/services/analytics/internal/domain/report"
	"prahari/services/analytics/internal/domain/search"
)

type Server struct {
	dashboardsSvc *dashboardsApp.Service
	reportingSvc  *reportingApp.Service
	analyticsSvc  *analyticsApp.Service
	searchSvc     *searchApp.Service
}

func NewServer(
	dashboardsSvc *dashboardsApp.Service,
	reportingSvc *reportingApp.Service,
	analyticsSvc *analyticsApp.Service,
	searchSvc *searchApp.Service,
) *Server {
	return &Server{
		dashboardsSvc: dashboardsSvc,
		reportingSvc:  reportingSvc,
		analyticsSvc:  analyticsSvc,
		searchSvc:     searchSvc,
	}
}

func (s *Server) GetDashboard(ctx context.Context, id string) (*dashboard.Dashboard, error) {
	return s.dashboardsSvc.GetDashboard(ctx, id)
}

func (s *Server) GetKPIs(ctx context.Context, plantID string) ([]*kpi.KPI, error) {
	return s.analyticsSvc.GetKPIs(ctx, plantID)
}

func (s *Server) GenerateReport(ctx context.Context, plantID, title, reportType, userID string) (*report.Report, error) {
	return s.reportingSvc.GenerateExecutiveReport(ctx, plantID, title, reportType, userID)
}

func (s *Server) GetMetrics(ctx context.Context, criteria *search.Criteria) ([]*metric.Metric, int64, error) {
	return s.searchSvc.ExecuteSearch(ctx, criteria)
}

func (s *Server) SearchAnalytics(ctx context.Context, criteria *search.Criteria) ([]*metric.Metric, int64, error) {
	return s.searchSvc.ExecuteSearch(ctx, criteria)
}
