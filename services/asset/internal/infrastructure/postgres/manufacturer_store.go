package postgres

import (
	"context"
	"database/sql"
	"fmt"

	manufacturerDomain "prahari/services/asset/internal/domain/manufacturer"
)

// ManufacturerStore implements manufacturer contact directory.
type ManufacturerStore struct {
	db *sql.DB
}

// NewManufacturerStore instantiates ManufacturerStore.
func NewManufacturerStore(db *sql.DB) *ManufacturerStore {
	return &ManufacturerStore{db: db}
}

// Create inserts details.
func (s *ManufacturerStore) Create(ctx context.Context, m *manufacturerDomain.Manufacturer) error {
	query := `INSERT INTO asset_manufacturers (id, name, contact_email, contact_phone, support_website)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, m.ID, m.Name, m.ContactEmail, m.ContactPhone, m.SupportWebsite)
	return err
}

// FindByID returns manufacturer details.
func (s *ManufacturerStore) FindByID(ctx context.Context, id string) (*manufacturerDomain.Manufacturer, error) {
	query := `SELECT id, name, contact_email, contact_phone, support_website FROM asset_manufacturers WHERE id = $1`
	m := &manufacturerDomain.Manufacturer{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&m.ID, &m.Name, &m.ContactEmail, &m.ContactPhone, &m.SupportWebsite)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("manufacturer not found: %s", id)
		}
		return nil, err
	}
	return m, nil
}
