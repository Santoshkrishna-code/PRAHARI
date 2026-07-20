package postgres

import (
	"context"
	"database/sql"

	partDomain "prahari/services/maintenance/internal/domain/sparepart"
)

// PartStore implements material replacements usage logging.
type PartStore struct {
	db *sql.DB
}

// NewPartStore instantiates PartStore.
func NewPartStore(db *sql.DB) *PartStore {
	return &PartStore{db: db}
}

// Create inserts spare part log.
func (s *PartStore) Create(ctx context.Context, sp *partDomain.SparePart) error {
	query := `INSERT INTO maintenance_spare_parts (id, maintenance_id, part_number, quantity_used, unit_cost)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, sp.ID, sp.MaintenanceID, sp.PartNumber, sp.QuantityUsed, sp.UnitCost)
	return err
}

// FindByMaintenanceID returns material replacements inventory lists.
func (s *PartStore) FindByMaintenanceID(ctx context.Context, maintenanceID string) ([]*partDomain.SparePart, error) {
	query := `SELECT id, maintenance_id, part_number, quantity_used, unit_cost FROM maintenance_spare_parts WHERE maintenance_id = $1`
	rows, err := s.db.QueryContext(ctx, query, maintenanceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*partDomain.SparePart
	for rows.Next() {
		sp := &partDomain.SparePart{}
		err = rows.Scan(&sp.ID, &sp.MaintenanceID, &sp.PartNumber, &sp.QuantityUsed, &sp.UnitCost)
		if err != nil {
			return nil, err
		}
		list = append(list, sp)
	}
	return list, nil
}
