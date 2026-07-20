package postgres

import (
	"context"
	"database/sql"
	"fmt"

	typeDomain "prahari/services/permit/internal/domain/permittype"
)

// TypeStore implements permit template queries.
type TypeStore struct {
	db *sql.DB
}

// NewTypeStore instantiates a TypeStore.
func NewTypeStore(db *sql.DB) *TypeStore {
	return &TypeStore{db: db}
}

// Create stores template metadata.
func (s *TypeStore) Create(ctx context.Context, pt *typeDomain.PermitType) error {
	query := `INSERT INTO permit_types (id, code, name, description, default_duration_hours, preconditions, is_active)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, pt.ID, pt.Code, pt.Name, pt.Description, pt.DefaultDurationHours, pt.Preconditions, pt.IsActive)
	return err
}

// FindByID returns permit category by ID.
func (s *TypeStore) FindByID(ctx context.Context, id string) (*typeDomain.PermitType, error) {
	query := `SELECT id, code, name, description, default_duration_hours, preconditions, is_active FROM permit_types WHERE id = $1`
	pt := &typeDomain.PermitType{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&pt.ID, &pt.Code, &pt.Name, &pt.Description, &pt.DefaultDurationHours, &pt.Preconditions, &pt.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("permit type not found: %s", id)
		}
		return nil, err
	}
	return pt, nil
}

// ListActive returns templates marked active.
func (s *TypeStore) ListActive(ctx context.Context) ([]*typeDomain.PermitType, error) {
	query := `SELECT id, code, name, description, default_duration_hours, preconditions, is_active FROM permit_types WHERE is_active = true`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []*typeDomain.PermitType
	for rows.Next() {
		pt := &typeDomain.PermitType{}
		if err := rows.Scan(&pt.ID, &pt.Code, &pt.Name, &pt.Description, &pt.DefaultDurationHours, &pt.Preconditions, &pt.IsActive); err != nil {
			return nil, err
		}
		types = append(types, pt)
	}
	return types, nil
}
