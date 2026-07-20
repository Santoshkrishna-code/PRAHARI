package postgres

import (
	"context"
	"database/sql"

	hazardDomain "prahari/services/permit/internal/domain/hazard"
)

// HazardStore implements hazard logging tables.
type HazardStore struct {
	db *sql.DB
}

// NewHazardStore instantiates HazardStore.
func NewHazardStore(db *sql.DB) *HazardStore {
	return &HazardStore{db: db}
}

// Create inserts identified threat points.
func (s *HazardStore) Create(ctx context.Context, h *hazardDomain.Hazard) error {
	query := `INSERT INTO permit_hazards (id, permit_id, type, description, control_measure, identified_by, identified_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, h.ID, h.PermitID, h.Type, h.Description, h.ControlMeasure, h.IdentifiedBy, h.IdentifiedAt)
	return err
}

// FindByPermitID returns identified points.
func (s *HazardStore) FindByPermitID(ctx context.Context, permitID string) ([]*hazardDomain.Hazard, error) {
	query := `SELECT id, permit_id, type, description, control_measure, identified_by, identified_at FROM permit_hazards WHERE permit_id = $1`
	rows, err := s.db.QueryContext(ctx, query, permitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hazards []*hazardDomain.Hazard
	for rows.Next() {
		h := &hazardDomain.Hazard{}
		err = rows.Scan(&h.ID, &h.PermitID, &h.Type, &h.Description, &h.ControlMeasure, &h.IdentifiedBy, &h.IdentifiedAt)
		if err != nil {
			return nil, err
		}
		hazards = append(hazards, h)
	}
	return hazards, nil
}
