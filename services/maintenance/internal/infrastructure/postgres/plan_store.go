package postgres

import (
	"context"
	"database/sql"
	"fmt"

	planDomain "prahari/services/maintenance/internal/domain/maintenanceplan"
)

// PlanStore implements preventive recurrence plan queries.
type PlanStore struct {
	db *sql.DB
}

// NewPlanStore instantiates a PlanStore.
func NewPlanStore(db *sql.DB) *PlanStore {
	return &PlanStore{db: db}
}

// Create inserts plan.
func (s *PlanStore) Create(ctx context.Context, p *planDomain.MaintenancePlan) error {
	query := `INSERT INTO maintenance_plans (id, asset_id, title, interval_code, is_active, last_run_date, next_run_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, p.ID, p.AssetID, p.Title, string(p.Interval), p.IsActive, p.LastRunDate, p.NextRunDate)
	return err
}

// FindByID returns plan.
func (s *PlanStore) FindByID(ctx context.Context, id string) (*planDomain.MaintenancePlan, error) {
	query := `SELECT id, asset_id, title, interval_code, is_active, last_run_date, next_run_date FROM maintenance_plans WHERE id = $1`
	p := &planDomain.MaintenancePlan{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&p.ID, &p.AssetID, &p.Title, &p.Interval, &p.IsActive, &p.LastRunDate, &p.NextRunDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("plan not found: %s", id)
		}
		return nil, err
	}
	return p, nil
}
