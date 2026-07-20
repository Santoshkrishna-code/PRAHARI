package postgres

import (
	"context"
	"database/sql"
	"fmt"

	severityDomain "prahari/services/nearmiss/internal/domain/severity"
)

// SeverityStore implements severity definitions operations.
type SeverityStore struct {
	db *sql.DB
}

// NewSeverityStore instantiates SeverityStore.
func NewSeverityStore(db *sql.DB) *SeverityStore {
	return &SeverityStore{db: db}
}

// Create persists severity scale.
func (s *SeverityStore) Create(ctx context.Context, sv *severityDomain.Severity) error {
	query := `INSERT INTO near_miss_severity (id, level, score, description) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, sv.ID, sv.Level, sv.Score, sv.Description)
	return err
}

// FindByID returns severity level details.
func (s *SeverityStore) FindByID(ctx context.Context, id string) (*severityDomain.Severity, error) {
	query := `SELECT id, level, score, description FROM near_miss_severity WHERE id = $1`
	sv := &severityDomain.Severity{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&sv.ID, &sv.Level, &sv.Score, &sv.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("severity not found: %s", id)
		}
		return nil, err
	}
	return sv, nil
}
