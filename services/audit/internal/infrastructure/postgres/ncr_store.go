package postgres

import (
	"context"
	"database/sql"

	ncrDomain "prahari/services/audit/internal/domain/nonconformity"
)

// NCRStore implements Non-Conformity logs database.
type NCRStore struct {
	db *sql.DB
}

// NewNCRStore instantiates NCRStore.
func NewNCRStore(db *sql.DB) *NCRStore {
	return &NCRStore{db: db}
}

// Create persists NCR details.
func (s *NCRStore) Create(ctx context.Context, nc *ncrDomain.NonConformity) error {
	query := `INSERT INTO nonconformities (id, finding_id, severity, description) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, nc.ID, nc.FindingID, nc.Severity, nc.Description)
	return err
}
