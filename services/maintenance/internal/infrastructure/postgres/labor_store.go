package postgres

import (
	"context"
	"database/sql"

	laborDomain "prahari/services/maintenance/internal/domain/labor"
)

// LaborStore implements labor billings log storage.
type LaborStore struct {
	db *sql.DB
}

// NewLaborStore instantiates LaborStore.
func NewLaborStore(db *sql.DB) *LaborStore {
	return &LaborStore{db: db}
}

// Create persists labor metrics.
func (s *LaborStore) Create(ctx context.Context, l *laborDomain.Labor) error {
	query := `INSERT INTO maintenance_labor (id, maintenance_id, technician_id, hours_worked, hourly_rate)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, l.ID, l.MaintenanceID, l.TechnicianID, l.HoursWorked, l.HourlyRate)
	return err
}

// FindByMaintenanceID returns labor hours worked records.
func (s *LaborStore) FindByMaintenanceID(ctx context.Context, maintenanceID string) ([]*laborDomain.Labor, error) {
	query := `SELECT id, maintenance_id, technician_id, hours_worked, hourly_rate FROM maintenance_labor WHERE maintenance_id = $1`
	rows, err := s.db.QueryContext(ctx, query, maintenanceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*laborDomain.Labor
	for rows.Next() {
		l := &laborDomain.Labor{}
		err = rows.Scan(&l.ID, &l.MaintenanceID, &l.TechnicianID, &l.HoursWorked, &l.HourlyRate)
		if err != nil {
			return nil, err
		}
		list = append(list, l)
	}
	return list, nil
}
