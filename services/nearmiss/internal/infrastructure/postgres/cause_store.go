package postgres

import (
	"context"
	"database/sql"

	causeDomain "prahari/services/nearmiss/internal/domain/cause"
)

// CauseStore implements root causes classification mapping.
type CauseStore struct {
	db *sql.DB
}

// NewCauseStore instantiates CauseStore.
func NewCauseStore(db *sql.DB) *CauseStore {
	return &CauseStore{db: db}
}

// Create persists cause.
func (s *CauseStore) Create(ctx context.Context, c *causeDomain.Cause) error {
	query := `INSERT INTO near_miss_causes (id, near_miss_id, root_cause, description) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.NearMissID, c.RootCause, c.Description)
	return err
}

// FindByNearMissID returns root causes lists.
func (s *CauseStore) FindByNearMissID(ctx context.Context, nearmissID string) ([]*causeDomain.Cause, error) {
	query := `SELECT id, near_miss_id, root_cause, description FROM near_miss_causes WHERE near_miss_id = $1`
	rows, err := s.db.QueryContext(ctx, query, nearmissID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*causeDomain.Cause
	for rows.Next() {
		c := &causeDomain.Cause{}
		err = rows.Scan(&c.ID, &c.NearMissID, &c.RootCause, &c.Description)
		if err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, nil
}
