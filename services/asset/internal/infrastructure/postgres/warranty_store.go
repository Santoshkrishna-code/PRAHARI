package postgres

import (
	"context"
	"database/sql"

	warrantyDomain "prahari/services/asset/internal/domain/warranty"
)

// WarrantyStore implements warranty contracts tracking.
type WarrantyStore struct {
	db *sql.DB
}

// NewWarrantyStore instantiates WarrantyStore.
func NewWarrantyStore(db *sql.DB) *WarrantyStore {
	return &WarrantyStore{db: db}
}

// Create persists contract details.
func (s *WarrantyStore) Create(ctx context.Context, w *warrantyDomain.Warranty) error {
	query := `INSERT INTO asset_warranties (id, asset_id, start_date, end_date, coverage_detail, contact_person, contact_email)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, w.ID, w.AssetID, w.StartDate, w.EndDate, w.CoverageDetail, w.ContactPerson, w.ContactEmail)
	return err
}

// FindByAssetID returns warranty.
func (s *WarrantyStore) FindByAssetID(ctx context.Context, assetID string) (*warrantyDomain.Warranty, error) {
	query := `SELECT id, asset_id, start_date, end_date, coverage_detail, contact_person, contact_email FROM asset_warranties WHERE asset_id = $1`
	w := &warrantyDomain.Warranty{}
	err := s.db.QueryRowContext(ctx, query, assetID).Scan(&w.ID, &w.AssetID, &w.StartDate, &w.EndDate, &w.CoverageDetail, &w.ContactPerson, &w.ContactEmail)
	if err != nil {
		return nil, err
	}
	return w, nil
}
