package postgres

import (
	"context"
	"database/sql"

	certDomain "prahari/services/contractor/internal/domain/certification"
)

// CertStore implements certifications validation queries.
type CertStore struct {
	db *sql.DB
}

// NewCertStore instantiates CertStore.
func NewCertStore(db *sql.DB) *CertStore {
	return &CertStore{db: db}
}

// Create inserts certification record.
func (s *CertStore) Create(ctx context.Context, c *certDomain.Certification) error {
	query := `INSERT INTO contractor_certifications (id, worker_id, cert_number, title, expiry_date, issuer)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.WorkerID, c.CertNumber, c.Title, c.ExpiryDate, c.Issuer)
	return err
}

// FindByWorkerID returns crew member certifications.
func (s *CertStore) FindByWorkerID(ctx context.Context, workerID string) ([]*certDomain.Certification, error) {
	query := `SELECT id, worker_id, cert_number, title, expiry_date, issuer FROM contractor_certifications WHERE worker_id = $1`
	rows, err := s.db.QueryContext(ctx, query, workerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*certDomain.Certification
	for rows.Next() {
		c := &certDomain.Certification{}
		err = rows.Scan(&c.ID, &c.WorkerID, &c.CertNumber, &c.Title, &c.ExpiryDate, &c.Issuer)
		if err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, nil
}
