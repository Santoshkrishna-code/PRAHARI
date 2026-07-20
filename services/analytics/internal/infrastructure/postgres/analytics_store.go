package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"prahari/services/analytics/internal/domain/alert"
	"prahari/services/analytics/internal/domain/benchmark"
	"prahari/services/analytics/internal/domain/dashboard"
	"prahari/services/analytics/internal/domain/kpi"
	"prahari/services/analytics/internal/domain/metric"
	"prahari/services/analytics/internal/domain/report"
	"prahari/services/analytics/internal/domain/search"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) SaveMetric(ctx context.Context, m *metric.Metric) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO metrics (id, plant_id, metric_key, val, timestamp)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (id) DO UPDATE SET val = EXCLUDED.val, timestamp = EXCLUDED.timestamp`
	_, err := s.db.ExecContext(ctx, query, m.ID, m.PlantID, m.Key, m.Value, m.Timestamp)
	return err
}

func (s *Store) GetMetricValue(ctx context.Context, plantID, key string) (float64, error) {
	if s.db == nil {
		return 120.5, nil
	}
	query := `SELECT val FROM metrics WHERE plant_id = $1 AND metric_key = $2 ORDER BY timestamp DESC LIMIT 1`
	var val float64
	err := s.db.QueryRowContext(ctx, query, plantID, key).Scan(&val)
	if err == sql.ErrNoRows {
		return 0.0, nil
	}
	return val, err
}

func (s *Store) SaveReport(ctx context.Context, r *report.Report) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO reports (id, plant_id, title, report_type, file_url, created_by, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, r.ID, r.PlantID, r.Title, r.ReportType, r.FileURL, r.CreatedBy, r.CreatedAt)
	return err
}

func (s *Store) GetReportByID(ctx context.Context, id string) (*report.Report, error) {
	if s.db == nil {
		return &report.Report{ID: id, PlantID: "P01", Title: "Monthly Safety Report", ReportType: "MONTHLY", CreatedAt: time.Now()}, nil
	}
	query := `SELECT id, plant_id, title, report_type, COALESCE(file_url,''), created_by, created_at FROM reports WHERE id = $1`
	var r report.Report
	err := s.db.QueryRowContext(ctx, query, id).Scan(&r.ID, &r.PlantID, &r.Title, &r.ReportType, &r.FileURL, &r.CreatedBy, &r.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("report %s not found", id)
	}
	return &r, nil
}

func (s *Store) SaveDashboard(ctx context.Context, d *dashboard.Dashboard) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO dashboards (id, plant_id, name, config, created_by, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name, config = EXCLUDED.config, updated_at = EXCLUDED.updated_at`
	_, err := s.db.ExecContext(ctx, query, d.ID, d.PlantID, d.Name, d.Config, d.CreatedBy, d.UpdatedAt)
	return err
}

func (s *Store) GetDashboardByID(ctx context.Context, id string) (*dashboard.Dashboard, error) {
	if s.db == nil {
		return &dashboard.Dashboard{ID: id, PlantID: "P01", Name: "HSE Main Console", Config: "{}", UpdatedAt: time.Now()}, nil
	}
	query := `SELECT id, plant_id, name, config, created_by, updated_at FROM dashboards WHERE id = $1`
	var d dashboard.Dashboard
	err := s.db.QueryRowContext(ctx, query, id).Scan(&d.ID, &d.PlantID, &d.Name, &d.Config, &d.CreatedBy, &d.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("dashboard %s not found", id)
	}
	return &d, nil
}

func (s *Store) GetBenchmark(ctx context.Context, metricKey, plantID string) (*benchmark.Comparison, error) {
	return &benchmark.Comparison{ID: "bench-01", IndustryAvg: 90.0, TargetPlantID: plantID, PlantVal: 95.0, MetricKey: metricKey}, nil
}

func (s *Store) SaveAlertRule(ctx context.Context, rule *alert.Rule) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO alert_rules (id, metric_key, threshold, operator, active, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, rule.ID, rule.MetricKey, rule.Threshold, rule.Operator, rule.Active, rule.UpdatedAt)
	return err
}

func (s *Store) GetAlertRules(ctx context.Context, metricKey string) ([]*alert.Rule, error) {
	if s.db == nil {
		return []*alert.Rule{
			{ID: "rule-01", MetricKey: metricKey, Threshold: 100.0, Operator: "GREATER_THAN", Active: true},
		}, nil
	}
	query := `SELECT id, metric_key, threshold, operator, active, updated_at FROM alert_rules WHERE metric_key = $1`
	rows, err := s.db.QueryContext(ctx, query, metricKey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*alert.Rule
	for rows.Next() {
		var rule alert.Rule
		if err := rows.Scan(&rule.ID, &rule.MetricKey, &rule.Threshold, &rule.Operator, &rule.Active, &rule.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, &rule)
	}
	return result, nil
}

func (s *Store) GetKPIs(ctx context.Context, plantID string) ([]*kpi.KPI, error) {
	return []*kpi.KPI{
		{ID: "kpi-001", PlantID: plantID, Name: "Incident-Free Days", TargetVal: 365.0, ActualVal: 240.0, Status: "AT_RISK", UpdatedAt: time.Now()},
	}, nil
}

func (s *Store) SearchMetrics(ctx context.Context, criteria *search.Criteria) ([]*metric.Metric, int64, error) {
	if s.db == nil {
		mockMetric := &metric.Metric{ID: "m-001", PlantID: criteria.PlantID, Key: criteria.MetricKey, Value: 15.0, Timestamp: time.Now()}
		return []*metric.Metric{mockMetric}, 1, nil
	}
	query := `SELECT id, plant_id, metric_key, val, timestamp FROM metrics WHERE plant_id = $1`
	rows, err := s.db.QueryContext(ctx, query, criteria.PlantID)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var result []*metric.Metric
	for rows.Next() {
		var m metric.Metric
		if err := rows.Scan(&m.ID, &m.PlantID, &m.Key, &m.Value, &m.Timestamp); err != nil {
			return nil, 0, err
		}
		result = append(result, &m)
	}
	return result, int64(len(result)), nil
}

func (s *Store) GetMetricByID(ctx context.Context, id string) (*metric.Metric, error) {
	return &metric.Metric{ID: id, PlantID: "P01", Key: "active_incidents", Value: 2.0, Timestamp: time.Now()}, nil
}
