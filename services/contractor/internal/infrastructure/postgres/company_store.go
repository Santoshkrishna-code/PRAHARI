package postgres

import (
	"context"
	"database/sql"
	"fmt"

	companyDomain "prahari/services/contractor/internal/domain/company"
)

// CompanyStore implements contractor companies persistence operations.
type CompanyStore struct {
	db *sql.DB
}

// NewCompanyStore instantiates CompanyStore.
func NewCompanyStore(db *sql.DB) *CompanyStore {
	return &CompanyStore{db: db}
}

// Create persists a company.
func (s *CompanyStore) Create(ctx context.Context, c *companyDomain.Company) error {
	query := `INSERT INTO contractor_companies (id, name, description, is_active) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.Name, c.Description, c.IsActive)
	return err
}

// FindByID returns a company.
func (s *CompanyStore) FindByID(ctx context.Context, id string) (*companyDomain.Company, error) {
	query := `SELECT id, name, description, is_active FROM contractor_companies WHERE id = $1`
	c := &companyDomain.Company{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.Name, &c.Description, &c.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("company not found: %s", id)
		}
		return nil, err
	}
	return c, nil
}
