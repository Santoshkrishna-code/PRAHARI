package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	maintenanceDomain "prahari/services/maintenance/internal/domain/maintenance"
)

// MaintenanceStore implements the core maintenance repository port.
type MaintenanceStore struct {
	db *sql.DB
}

// NewMaintenanceStore instantiates a MaintenanceStore.
func NewMaintenanceStore(db *sql.DB) *MaintenanceStore {
	return &MaintenanceStore{db: db}
}

// Create inserts a maintenance record.
func (s *MaintenanceStore) Create(ctx context.Context, m *maintenanceDomain.Maintenance) error {
	query := `INSERT INTO maintenance (
		id, maintenance_number, asset_id, maintenance_type, priority, department_id,
		title, description, status_code, total_estimated_cost, total_actual_cost, created_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	_, err := s.db.ExecContext(ctx, query,
		m.ID, m.MaintenanceNumber, m.AssetID, m.MaintenanceType, string(m.Priority), m.DepartmentID,
		m.Title, m.Description, m.StatusCode, m.TotalEstimatedCost, m.TotalActualCost, m.CreatedAt, m.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert maintenance: %w", err)
	}
	return nil
}

// FindByID retrieves a maintenance by ID.
func (s *MaintenanceStore) FindByID(ctx context.Context, id string) (*maintenanceDomain.Maintenance, error) {
	query := `SELECT id, maintenance_number, asset_id, maintenance_type, priority, department_id,
		title, description, status_code, total_estimated_cost, total_actual_cost, created_at, updated_at
		FROM maintenance WHERE id = $1 AND is_deleted = false`

	m := &maintenanceDomain.Maintenance{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&m.ID, &m.MaintenanceNumber, &m.AssetID, &m.MaintenanceType, &m.Priority, &m.DepartmentID,
		&m.Title, &m.Description, &m.StatusCode, &m.TotalEstimatedCost, &m.TotalActualCost, &m.CreatedAt, &m.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("maintenance record not found: %s", id)
		}
		return nil, fmt.Errorf("postgres: failed to query maintenance: %w", err)
	}
	return m, nil
}

// FindByNumber retrieves a maintenance by number.
func (s *MaintenanceStore) FindByNumber(ctx context.Context, number string) (*maintenanceDomain.Maintenance, error) {
	query := `SELECT id, maintenance_number, asset_id, maintenance_type, priority, department_id,
		title, description, status_code, total_estimated_cost, total_actual_cost, created_at, updated_at
		FROM maintenance WHERE maintenance_number = $1 AND is_deleted = false`

	m := &maintenanceDomain.Maintenance{}
	err := s.db.QueryRowContext(ctx, query, number).Scan(
		&m.ID, &m.MaintenanceNumber, &m.AssetID, &m.MaintenanceType, &m.Priority, &m.DepartmentID,
		&m.Title, &m.Description, &m.StatusCode, &m.TotalEstimatedCost, &m.TotalActualCost, &m.CreatedAt, &m.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("maintenance not found: %s", number)
		}
		return nil, fmt.Errorf("postgres: failed to query maintenance by number: %w", err)
	}
	return m, nil
}

// Update persists modifications.
func (s *MaintenanceStore) Update(ctx context.Context, m *maintenanceDomain.Maintenance) error {
	query := `UPDATE maintenance SET
		title = $2, description = $3, status_code = $4, total_estimated_cost = $5,
		total_actual_cost = $6, updated_at = $7 WHERE id = $1`

	_, err := s.db.ExecContext(ctx, query,
		m.ID, m.Title, m.Description, m.StatusCode, m.TotalEstimatedCost, m.TotalActualCost, time.Now(),
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to update maintenance: %w", err)
	}
	return nil
}

// Delete marks a maintenance record deleted.
func (s *MaintenanceStore) Delete(ctx context.Context, id string) error {
	query := `UPDATE maintenance SET is_deleted = true, updated_at = $2 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id, time.Now())
	if err != nil {
		return fmt.Errorf("postgres: failed to soft-delete maintenance: %w", err)
	}
	return nil
}

// List returns pages.
func (s *MaintenanceStore) List(ctx context.Context, offset, limit int) ([]*maintenanceDomain.Maintenance, int, error) {
	var total int
	if err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM maintenance WHERE is_deleted = false`).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `SELECT id, maintenance_number, asset_id, maintenance_type, priority, department_id,
		title, description, status_code, total_estimated_cost, total_actual_cost, created_at, updated_at
		FROM maintenance WHERE is_deleted = false ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var maintenanceList []*maintenanceDomain.Maintenance
	for rows.Next() {
		m := &maintenanceDomain.Maintenance{}
		err = rows.Scan(
			&m.ID, &m.MaintenanceNumber, &m.AssetID, &m.MaintenanceType, &m.Priority, &m.DepartmentID,
			&m.Title, &m.Description, &m.StatusCode, &m.TotalEstimatedCost, &m.TotalActualCost, &m.CreatedAt, &m.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		maintenanceList = append(maintenanceList, m)
	}
	return maintenanceList, total, nil
}
