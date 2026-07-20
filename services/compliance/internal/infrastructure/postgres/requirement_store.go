package postgres

import (
	"context"
	"database/sql"

	reqDomain "prahari/services/compliance/internal/domain/requirement"
)

// RequirementStore implements standards rules clauses.
type RequirementStore struct {
	db *sql.DB
}

// NewRequirementStore instantiates RequirementStore.
func NewRequirementStore(db *sql.DB) *RequirementStore {
	return &RequirementStore{db: db}
}

// Create persists requirement rules clauses.
func (s *RequirementStore) Create(ctx context.Context, r *reqDomain.Requirement) error {
	query := `INSERT INTO requirements (id, obligation_id, clause, description) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, r.ID, r.ObligationID, r.Clause, r.Description)
	return err
}

// FindByObligationID returns requirement checklist clauses.
func (s *RequirementStore) FindByObligationID(ctx context.Context, obligationID string) ([]*reqDomain.Requirement, error) {
	query := `SELECT id, obligation_id, clause, description FROM requirements WHERE obligation_id = $1`
	rows, err := s.db.QueryContext(ctx, query, obligationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*reqDomain.Requirement
	for rows.Next() {
		r := &reqDomain.Requirement{}
		err = rows.Scan(&r.ID, &r.ObligationID, &r.Clause, &r.Description)
		if err != nil {
			return nil, err
		}
		list = append(list, r)
	}
	return list, nil
}
