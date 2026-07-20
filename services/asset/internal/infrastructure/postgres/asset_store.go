package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	assetDomain "prahari/services/asset/internal/domain/asset"
)

// AssetStore implements the core asset repository port.
type AssetStore struct {
	db *sql.DB
}

// NewAssetStore instantiates an AssetStore.
func NewAssetStore(db *sql.DB) *AssetStore {
	return &AssetStore{db: db}
}

// Create inserts an asset record.
func (s *AssetStore) Create(ctx context.Context, a *assetDomain.Asset) error {
	query := `INSERT INTO assets (
		id, asset_number, name, description, serial_number, lifecycle_status, operational_status,
		criticality_code, department_id, location_id, category_id, type_id, manufacturer_id,
		model_number, purchase_date, installation_date, last_maintenance_date, health_score,
		condition_score, remaining_useful_life, failure_probability, created_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)`

	_, err := s.db.ExecContext(ctx, query,
		a.ID, a.AssetNumber, a.Name, a.Description, a.SerialNumber, a.LifecycleStatus, string(a.OperationalStatus),
		string(a.CriticalityCode), a.DepartmentID, a.LocationID, a.CategoryID, a.TypeID, a.ManufacturerID,
		a.ModelNumber, a.PurchaseDate, a.InstallationDate, a.LastMaintenanceDate, a.HealthScore,
		a.ConditionScore, a.RemainingUsefulLife, a.FailureProbability, a.CreatedAt, a.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert asset: %w", err)
	}
	return nil
}

// FindByID retrieves an asset by ID.
func (s *AssetStore) FindByID(ctx context.Context, id string) (*assetDomain.Asset, error) {
	query := `SELECT id, asset_number, name, description, serial_number, lifecycle_status, operational_status,
		criticality_code, department_id, location_id, category_id, type_id, manufacturer_id,
		model_number, purchase_date, installation_date, last_maintenance_date, health_score,
		condition_score, remaining_useful_life, failure_probability, created_at, updated_at
		FROM assets WHERE id = $1 AND is_deleted = false`

	a := &assetDomain.Asset{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&a.ID, &a.AssetNumber, &a.Name, &a.Description, &a.SerialNumber, &a.LifecycleStatus, &a.OperationalStatus,
		&a.CriticalityCode, &a.DepartmentID, &a.LocationID, &a.CategoryID, &a.TypeID, &a.ManufacturerID,
		&a.ModelNumber, &a.PurchaseDate, &a.InstallationDate, &a.LastMaintenanceDate, &a.HealthScore,
		&a.ConditionScore, &a.RemainingUsefulLife, &a.FailureProbability, &a.CreatedAt, &a.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("asset not found: %s", id)
		}
		return nil, fmt.Errorf("postgres: failed to query asset: %w", err)
	}
	return a, nil
}

// FindByNumber retrieves an asset by number.
func (s *AssetStore) FindByNumber(ctx context.Context, number string) (*assetDomain.Asset, error) {
	query := `SELECT id, asset_number, name, description, serial_number, lifecycle_status, operational_status,
		criticality_code, department_id, location_id, category_id, type_id, manufacturer_id,
		model_number, purchase_date, installation_date, last_maintenance_date, health_score,
		condition_score, remaining_useful_life, failure_probability, created_at, updated_at
		FROM assets WHERE asset_number = $1 AND is_deleted = false`

	a := &assetDomain.Asset{}
	err := s.db.QueryRowContext(ctx, query, number).Scan(
		&a.ID, &a.AssetNumber, &a.Name, &a.Description, &a.SerialNumber, &a.LifecycleStatus, &a.OperationalStatus,
		&a.CriticalityCode, &a.DepartmentID, &a.LocationID, &a.CategoryID, &a.TypeID, &a.ManufacturerID,
		&a.ModelNumber, &a.PurchaseDate, &a.InstallationDate, &a.LastMaintenanceDate, &a.HealthScore,
		&a.ConditionScore, &a.RemainingUsefulLife, &a.FailureProbability, &a.CreatedAt, &a.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("asset not found: %s", number)
		}
		return nil, fmt.Errorf("postgres: failed to query asset by number: %w", err)
	}
	return a, nil
}

// Update persists asset modifications.
func (s *AssetStore) Update(ctx context.Context, a *assetDomain.Asset) error {
	query := `UPDATE assets SET
		name = $2, description = $3, lifecycle_status = $4, operational_status = $5,
		health_score = $6, condition_score = $7, remaining_useful_life = $8, failure_probability = $9,
		updated_at = $10 WHERE id = $1`

	_, err := s.db.ExecContext(ctx, query,
		a.ID, a.Name, a.Description, a.LifecycleStatus, string(a.OperationalStatus),
		a.HealthScore, a.ConditionScore, a.RemainingUsefulLife, a.FailureProbability,
		time.Now(),
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to update asset: %w", err)
	}
	return nil
}

// Delete marks an asset deleted.
func (s *AssetStore) Delete(ctx context.Context, id string) error {
	query := `UPDATE assets SET is_deleted = true, updated_at = $2 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id, time.Now())
	if err != nil {
		return fmt.Errorf("postgres: failed to soft-delete asset: %w", err)
	}
	return nil
}

// List returns pages.
func (s *AssetStore) List(ctx context.Context, offset, limit int) ([]*assetDomain.Asset, int, error) {
	var total int
	if err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM assets WHERE is_deleted = false`).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `SELECT id, asset_number, name, description, serial_number, lifecycle_status, operational_status,
		criticality_code, department_id, location_id, category_id, type_id, manufacturer_id,
		model_number, purchase_date, installation_date, last_maintenance_date, health_score,
		condition_score, remaining_useful_life, failure_probability, created_at, updated_at
		FROM assets WHERE is_deleted = false ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var assets []*assetDomain.Asset
	for rows.Next() {
		a := &assetDomain.Asset{}
		err = rows.Scan(
			&a.ID, &a.AssetNumber, &a.Name, &a.Description, &a.SerialNumber, &a.LifecycleStatus, &a.OperationalStatus,
			&a.CriticalityCode, &a.DepartmentID, &a.LocationID, &a.CategoryID, &a.TypeID, &a.ManufacturerID,
			&a.ModelNumber, &a.PurchaseDate, &a.InstallationDate, &a.LastMaintenanceDate, &a.HealthScore,
			&a.ConditionScore, &a.RemainingUsefulLife, &a.FailureProbability, &a.CreatedAt, &a.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		assets = append(assets, a)
	}
	return assets, total, nil
}
