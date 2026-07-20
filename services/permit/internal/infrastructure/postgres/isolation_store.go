package postgres

import (
	"context"
	"database/sql"

	isolationDomain "prahari/services/permit/internal/domain/isolation"
)

// IsolationStore implements lock verification storage.
type IsolationStore struct {
	db *sql.DB
}

// NewIsolationStore instantiates IsolationStore.
func NewIsolationStore(db *sql.DB) *IsolationStore {
	return &IsolationStore{db: db}
}

// Create persists isolation lock points.
func (s *IsolationStore) Create(ctx context.Context, i *isolationDomain.Isolation) error {
	query := `INSERT INTO permit_isolations (id, permit_id, isolation_type, equipment_id, equipment_description, isolation_point, lock_number, tag_number, isolated_by, isolated_at, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := s.db.ExecContext(ctx, query,
		i.ID, i.PermitID, i.IsolationType, i.EquipmentID, i.EquipmentDescription, i.IsolationPoint, i.LockNumber, i.TagNumber, i.IsolatedBy, i.IsolatedAt, i.Status,
	)
	return err
}

// FindByPermitID returns applied locks.
func (s *IsolationStore) FindByPermitID(ctx context.Context, permitID string) ([]*isolationDomain.Isolation, error) {
	query := `SELECT id, permit_id, isolation_type, equipment_id, equipment_description, isolation_point, lock_number, tag_number, isolated_by, isolated_at, verified_by, verified_at, removed_by, removed_at, status
		FROM permit_isolations WHERE permit_id = $1`
	rows, err := s.db.QueryContext(ctx, query, permitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var isolations []*isolationDomain.Isolation
	for rows.Next() {
		i := &isolationDomain.Isolation{}
		err = rows.Scan(&i.ID, &i.PermitID, &i.IsolationType, &i.EquipmentID, &i.EquipmentDescription, &i.IsolationPoint, &i.LockNumber, &i.TagNumber, &i.IsolatedBy, &i.IsolatedAt, &i.VerifiedBy, &i.VerifiedAt, &i.RemovedBy, &i.RemovedAt, &i.Status)
		if err != nil {
			return nil, err
		}
		isolations = append(isolations, i)
	}
	return isolations, nil
}

// Update registers lock releases.
func (s *IsolationStore) Update(ctx context.Context, i *isolationDomain.Isolation) error {
	query := `UPDATE permit_isolations SET verified_by = $2, verified_at = $3, removed_by = $4, removed_at = $5, status = $6 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, i.ID, i.VerifiedBy, i.VerifiedAt, i.RemovedBy, i.RemovedAt, i.Status)
	return err
}
