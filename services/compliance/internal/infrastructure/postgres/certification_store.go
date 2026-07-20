package postgres

import (
	"context"
	"database/sql"
	"fmt"

	certDomain "prahari/services/compliance/internal/domain/certification"
)

// CertificationStore implements credentials/licenses tracker database.
type CertificationStore struct {
	db *sql.DB
}

// NewCertificationStore instantiates CertificationStore.
func NewCertificationStore(db *sql.DB) *CertificationStore {
	return &CertificationStore{db: db}
}

// Create persists certification metadata trackers.
func (s *CertificationStore) Create(ctx context.Context, c *certDomain.Certification) error {
	query := `INSERT INTO certifications (id, issuer, valid_until, status) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.Issuer, c.ValidUntil, c.Status)
	return err
}

// FindByID returns validity details.
func (s *CertificationStore) FindByID(ctx context.Context, id string) (*certDomain.Certification, error) {
	query := `SELECT id, issuer, valid_until, status FROM certifications WHERE id = $1`
	c := &certDomain.Certification{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.Issuer, &c.ValidUntil, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("certification details not found: %s", id)
		}
		return nil, err
	}
	return c, nil
}
