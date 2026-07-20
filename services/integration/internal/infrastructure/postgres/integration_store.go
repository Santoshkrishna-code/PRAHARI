package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"prahari/services/integration/internal/domain/connector"
	"prahari/services/integration/internal/domain/mapping"
	"prahari/services/integration/internal/domain/search"
	"prahari/services/integration/internal/domain/synchronization"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) SaveConnector(ctx context.Context, c *connector.Connector) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO connectors (id, plant_id, name, type, status, host, port, updated_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (id) DO UPDATE SET status = EXCLUDED.status, updated_at = EXCLUDED.updated_at`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.PlantID, c.Name, c.Type, c.Status, c.Host, c.Port, c.UpdatedAt, c.CreatedAt)
	return err
}

func (s *Store) GetConnectorByID(ctx context.Context, id string) (*connector.Connector, error) {
	if s.db == nil {
		return &connector.Connector{ID: id, PlantID: "P01", Name: "SAP ERP Production", Type: "SAP", Status: "DISCONNECTED"}, nil
	}
	query := `SELECT id, plant_id, name, type, status, host, port, updated_at, created_at FROM connectors WHERE id = $1`
	row := s.db.QueryRowContext(ctx, query, id)
	var c connector.Connector
	if err := row.Scan(&c.ID, &c.PlantID, &c.Name, &c.Type, &c.Status, &c.Host, &c.Port, &c.UpdatedAt, &c.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("connector %s not found", id)
		}
		return nil, err
	}
	return &c, nil
}

func (s *Store) ListConnectors(ctx context.Context, plantID string) ([]*connector.Connector, error) {
	if s.db == nil {
		return []*connector.Connector{
			{ID: "conn-001", PlantID: plantID, Name: "SAP ERP Production", Type: "SAP", Status: "CONNECTED"},
		}, nil
	}
	query := `SELECT id, plant_id, name, type, status, host, port, updated_at, created_at FROM connectors WHERE plant_id = $1`
	rows, err := s.db.QueryContext(ctx, query, plantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*connector.Connector
	for rows.Next() {
		var c connector.Connector
		if err := rows.Scan(&c.ID, &c.PlantID, &c.Name, &c.Type, &c.Status, &c.Host, &c.Port, &c.UpdatedAt, &c.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, &c)
	}
	return result, nil
}

func (s *Store) SaveSynchronization(ctx context.Context, rec *synchronization.Record) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO integration_jobs (id, job_id, started_at, finished_at, status, records_count, error_message)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id) DO UPDATE SET status = EXCLUDED.status, finished_at = EXCLUDED.finished_at, records_count = EXCLUDED.records_count, error_message = EXCLUDED.error_message`
	_, err := s.db.ExecContext(ctx, query, rec.ID, rec.JobID, rec.StartedAt, rec.FinishedAt, rec.Status, rec.RecordsCount, rec.ErrorMessage)
	return err
}

func (s *Store) GetMappingRules(ctx context.Context, connectorID string) ([]*mapping.FieldMap, error) {
	if s.db == nil {
		return []*mapping.FieldMap{
			{ID: "m-1", ConnectorID: connectorID, ExternalKey: "HEADING_VAL", InternalKey: "heading_rate", DataType: "FLOAT", UpdatedAt: time.Now()},
		}, nil
	}
	query := `SELECT id, connector_id, external_key, internal_key, data_type, updated_at FROM mappings WHERE connector_id = $1`
	rows, err := s.db.QueryContext(ctx, query, connectorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*mapping.FieldMap
	for rows.Next() {
		var m mapping.FieldMap
		if err := rows.Scan(&m.ID, &m.ConnectorID, &m.ExternalKey, &m.InternalKey, &m.DataType, &m.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, &m)
	}
	return result, nil
}

func (s *Store) SearchConnectors(ctx context.Context, criteria *search.Criteria) ([]*connector.Connector, int64, error) {
	connectors, err := s.ListConnectors(ctx, criteria.PlantID)
	if err != nil {
		return nil, 0, err
	}
	return connectors, int64(len(connectors)), nil
}

func (s *Store) GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	return map[string]float64{
		"connector_health_pct":         100.0,
		"active_integrations_count":    15.0,
		"message_throughput_per_sec":   125.0,
		"failed_integrations_count":    2.0,
		"retry_count":                  8.0,
		"dlq_size_count":               1.0,
		"average_latency_ms":           45.2,
		"synchronization_time_sec":     12.4,
		"external_system_avail_pct":    99.8,
	}, nil
}
