package postgres

import (
	"context"
	"database/sql"

	certDomain "prahari/services/training/internal/domain/certification"
)

// CertStore implements certifications database.
type CertStore struct {
	db *sql.DB
}

// NewCertStore instantiates CertStore.
func NewCertStore(db *sql.DB) *CertStore {
	return &CertStore{db: db}
}

// Create persists certification tracks.
func (s *CertStore) Create(ctx context.Context, c *certDomain.Certification) error {
	query := `INSERT INTO certifications (id, trainee_id, course_id, issuer, valid_until)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.TraineeID, c.CourseID, c.Issuer, c.ValidUntil)
	return err
}
