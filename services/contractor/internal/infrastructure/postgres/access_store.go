package postgres

import (
	"context"
	"database/sql"

	accessDomain "prahari/services/contractor/internal/domain/siteaccess"
)

// AccessStore implements site access locations check logs.
type AccessStore struct {
	db *sql.DB
}

// NewAccessStore instantiates AccessStore.
func NewAccessStore(db *sql.DB) *AccessStore {
	return &AccessStore{db: db}
}

// Create persists gate entrance clearance parameters.
func (s *AccessStore) Create(ctx context.Context, sa *accessDomain.SiteAccess) error {
	query := `INSERT INTO contractor_site_access (id, worker_id, allowed_locations, badge_number, access_start, access_end)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, sa.ID, sa.WorkerID, sa.AllowedLocations, sa.BadgeNumber, sa.AccessStart, sa.AccessEnd)
	return err
}

// FindByWorkerID returns gate clearances checks.
func (s *AccessStore) FindByWorkerID(ctx context.Context, workerID string) ([]*accessDomain.SiteAccess, error) {
	query := `SELECT id, worker_id, allowed_locations, badge_number, access_start, access_end FROM contractor_site_access WHERE worker_id = $1`
	rows, err := s.db.QueryContext(ctx, query, workerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*accessDomain.SiteAccess
	for rows.Next() {
		sa := &accessDomain.SiteAccess{}
		err = rows.Scan(&sa.ID, &sa.WorkerID, &sa.AllowedLocations, &sa.BadgeNumber, &sa.AccessStart, &sa.AccessEnd)
		if err != nil {
			return nil, err
		}
		list = append(list, sa)
	}
	return list, nil
}
