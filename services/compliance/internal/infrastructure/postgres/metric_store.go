package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

// MetricStore implements process safety dashboard metrics index logs.
type MetricStore struct {
	db *sql.DB
}

// NewMetricStore instantiates MetricStore.
func NewMetricStore(db *sql.DB) *MetricStore {
	return &MetricStore{db: db}
}

// LogMetric records real-time index metrics scores.
func (s *MetricStore) LogMetric(ctx context.Context, key string, score float64) error {
	query := `INSERT INTO metrics (metric_key, score) VALUES ($1, $2)`
	_, err := s.db.ExecContext(ctx, query, key, score)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert metric: %w", err)
	}
	return nil
}
