package postgres

import (
	"context"
	"database/sql"

	riskDomain "prahari/services/risk/internal/domain/risk"
)

// HistoryStore implements risk register mutation backups catalog.
type HistoryStore struct {
	db *sql.DB
}

// NewHistoryStore instantiates HistoryStore.
func NewHistoryStore(db *sql.DB) *HistoryStore {
	return &HistoryStore{db: db}
}

// CreateBackup persists snapshot backups.
func (s *HistoryStore) CreateBackup(ctx context.Context, r *riskDomain.Risk) error {
	query := `INSERT INTO risk_history (id, risk_number, title, description) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, r.ID, r.RiskNumber, r.Title, r.Description)
	return err
}
