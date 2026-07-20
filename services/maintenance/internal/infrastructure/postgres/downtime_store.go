package postgres

import (
	"context"
	"database/sql"

	downtimeDomain "prahari/services/maintenance/internal/domain/downtime"
)

// DowntimeStore implements downtime logs storage.
type DowntimeStore struct {
	db *sql.DB
}

// NewDowntimeStore instantiates DowntimeStore.
func NewDowntimeStore(db *sql.DB) *DowntimeStore {
	return &DowntimeStore{db: db}
}

// Create persists downtime record.
func (s *DowntimeStore) Create(ctx context.Context, d *downtimeDomain.Downtime) error {
	query := `INSERT INTO maintenance_downtime (id, maintenance_id, asset_id, start_date, end_date, reason_code)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, d.ID, d.MaintenanceID, d.AssetID, d.StartDate, d.EndDate, d.ReasonCode)
	return err
}

// FindByAssetID returns downtime logs for an asset profile.
func (s *DowntimeStore) FindByAssetID(ctx context.Context, assetID string) ([]*downtimeDomain.Downtime, error) {
	query := `SELECT id, maintenance_id, asset_id, start_date, end_date, reason_code FROM maintenance_downtime WHERE asset_id = $1`
	rows, err := s.db.QueryContext(ctx, query, assetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*downtimeDomain.Downtime
	for rows.Next() {
		d := &downtimeDomain.Downtime{}
		err = rows.Scan(&d.ID, &d.MaintenanceID, &d.AssetID, &d.StartDate, &d.EndDate, &d.ReasonCode)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}
