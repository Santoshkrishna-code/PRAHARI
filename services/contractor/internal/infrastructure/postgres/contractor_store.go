package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	contractorDomain "prahari/services/contractor/internal/domain/contractor"
)

// ContractorStore implements the core contractor repository port.
type ContractorStore struct {
	db *sql.DB
}

// NewContractorStore instantiates a ContractorStore.
func NewContractorStore(db *sql.DB) *ContractorStore {
	return &ContractorStore{db: db}
}

// Create inserts a contractor record.
func (s *ContractorStore) Create(ctx context.Context, c *contractorDomain.Contractor) error {
	query := `INSERT INTO contractor (
		id, contractor_number, company_name, tax_id, status_code, department_id,
		registration_date, insurance_expiry, created_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := s.db.ExecContext(ctx, query,
		c.ID, c.ContractorNumber, c.CompanyName, c.TaxID, c.StatusCode, c.DepartmentID,
		c.RegistrationDate, c.InsuranceExpiry, c.CreatedAt, c.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert contractor: %w", err)
	}
	return nil
}

// FindByID retrieves a contractor by ID.
func (s *ContractorStore) FindByID(ctx context.Context, id string) (*contractorDomain.Contractor, error) {
	query := `SELECT id, contractor_number, company_name, tax_id, status_code, department_id,
		registration_date, insurance_expiry, created_at, updated_at
		FROM contractor WHERE id = $1 AND is_deleted = false`

	c := &contractorDomain.Contractor{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&c.ID, &c.ContractorNumber, &c.CompanyName, &c.TaxID, &c.StatusCode, &c.DepartmentID,
		&c.RegistrationDate, &c.InsuranceExpiry, &c.CreatedAt, &c.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("contractor record not found: %s", id)
		}
		return nil, fmt.Errorf("postgres: failed to query contractor: %w", err)
	}
	return c, nil
}

// FindByNumber retrieves a contractor by number.
func (s *ContractorStore) FindByNumber(ctx context.Context, number string) (*contractorDomain.Contractor, error) {
	query := `SELECT id, contractor_number, company_name, tax_id, status_code, department_id,
		registration_date, insurance_expiry, created_at, updated_at
		FROM contractor WHERE contractor_number = $1 AND is_deleted = false`

	c := &contractorDomain.Contractor{}
	err := s.db.QueryRowContext(ctx, query, number).Scan(
		&c.ID, &c.ContractorNumber, &c.CompanyName, &c.TaxID, &c.StatusCode, &c.DepartmentID,
		&c.RegistrationDate, &c.InsuranceExpiry, &c.CreatedAt, &c.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("contractor not found: %s", number)
		}
		return nil, fmt.Errorf("postgres: failed to query contractor by number: %w", err)
	}
	return c, nil
}

// Update persists modifications.
func (s *ContractorStore) Update(ctx context.Context, c *contractorDomain.Contractor) error {
	query := `UPDATE contractor SET
		company_name = $2, tax_id = $3, status_code = $4, insurance_expiry = $5,
		updated_at = $6 WHERE id = $1`

	_, err := s.db.ExecContext(ctx, query,
		c.ID, c.CompanyName, c.TaxID, c.StatusCode, c.InsuranceExpiry, time.Now(),
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to update contractor: %w", err)
	}
	return nil
}

// Delete marks a contractor record deleted.
func (s *ContractorStore) Delete(ctx context.Context, id string) error {
	query := `UPDATE contractor SET is_deleted = true, updated_at = $2 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id, time.Now())
	if err != nil {
		return fmt.Errorf("postgres: failed to soft-delete contractor: %w", err)
	}
	return nil
}

// List returns pages.
func (s *ContractorStore) List(ctx context.Context, offset, limit int) ([]*contractorDomain.Contractor, int, error) {
	var total int
	if err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM contractor WHERE is_deleted = false`).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `SELECT id, contractor_number, company_name, tax_id, status_code, department_id,
		registration_date, insurance_expiry, created_at, updated_at
		FROM contractor WHERE is_deleted = false ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var list []*contractorDomain.Contractor
	for rows.Next() {
		c := &contractorDomain.Contractor{}
		err = rows.Scan(
			&c.ID, &c.ContractorNumber, &c.CompanyName, &c.TaxID, &c.StatusCode, &c.DepartmentID,
			&c.RegistrationDate, &c.InsuranceExpiry, &c.CreatedAt, &c.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, c)
	}
	return list, total, nil
}
