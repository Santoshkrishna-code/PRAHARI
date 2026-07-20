package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"prahari/templates/service-template/internal/domain"
)

type PostgresPermitRepository struct {
	db *sql.DB
}

func NewPostgresPermitRepository(db *sql.DB) domain.PermitRepository {
	return &PostgresPermitRepository{db: db}
}

func (r *PostgresPermitRepository) GetByID(ctx context.Context, id string) (*domain.Permit, error) {
	query := `SELECT id, worker_id, zone_id, status, created_at, expires_at, approved_by 
	          FROM permits WHERE id = $1`
	
	row := r.db.QueryRowContext(ctx, query, id)
	
	var p domain.Permit
	err := row.Scan(&p.ID, &p.WorkerID, &p.ZoneID, &p.Status, &p.CreatedAt, &p.ExpiresAt, &p.ApprovedBy)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("permit %s not found: %w", id, sql.ErrNoRows)
		}
		return nil, fmt.Errorf("database query error: %w", err)
	}
	
	return &p, nil
}

func (r *PostgresPermitRepository) Create(ctx context.Context, p *domain.Permit) error {
	query := `INSERT INTO permits (id, worker_id, zone_id, status, created_at, expires_at, approved_by) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7)`
	
	_, err := r.db.ExecContext(ctx, query, p.ID, p.WorkerID, p.ZoneID, p.Status, p.CreatedAt, p.ExpiresAt, p.ApprovedBy)
	if err != nil {
		return fmt.Errorf("failed to insert permit: %w", err)
	}
	
	return nil
}

func (r *PostgresPermitRepository) Update(ctx context.Context, p *domain.Permit) error {
	query := `UPDATE permits SET status = $1, approved_by = $2, expires_at = $3 WHERE id = $4`
	
	result, err := r.db.ExecContext(ctx, query, p.Status, p.ApprovedBy, p.ExpiresAt, p.ID)
	if err != nil {
		return fmt.Errorf("failed to update permit: %w", err)
	}
	
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("permit %s not found for update", p.ID)
	}
	
	return nil
}
