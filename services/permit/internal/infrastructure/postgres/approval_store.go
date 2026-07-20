package postgres

import (
	"context"
	"database/sql"

	approvalDomain "prahari/services/permit/internal/domain/approval"
)

// ApprovalStore implements signature persistent tables.
type ApprovalStore struct {
	db *sql.DB
}

// NewApprovalStore instantiates ApprovalStore.
func NewApprovalStore(db *sql.DB) *ApprovalStore {
	return &ApprovalStore{db: db}
}

// Create registers signature steps.
func (s *ApprovalStore) Create(ctx context.Context, a *approvalDomain.Approval) error {
	query := `INSERT INTO permit_approvals (id, permit_id, approver_id, approver_role, decision, comments, signature_hash, decided_at, sequence_order)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := s.db.ExecContext(ctx, query,
		a.ID, a.PermitID, a.ApproverID, a.ApproverRole, a.Decision, a.Comments, a.SignatureHash, a.DecidedAt, a.SequenceOrder,
	)
	return err
}

// FindByID returns signature detail.
func (s *ApprovalStore) FindByID(ctx context.Context, id string) (*approvalDomain.Approval, error) {
	query := `SELECT id, permit_id, approver_id, approver_role, decision, comments, signature_hash, decided_at, sequence_order
		FROM permit_approvals WHERE id = $1`
	a := &approvalDomain.Approval{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&a.ID, &a.PermitID, &a.ApproverID, &a.ApproverRole, &a.Decision, &a.Comments, &a.SignatureHash, &a.DecidedAt, &a.SequenceOrder)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// FindByPermitID returns all signature steps for a permit.
func (s *ApprovalStore) FindByPermitID(ctx context.Context, permitID string) ([]*approvalDomain.Approval, error) {
	query := `SELECT id, permit_id, approver_id, approver_role, decision, comments, signature_hash, decided_at, sequence_order
		FROM permit_approvals WHERE permit_id = $1 ORDER BY sequence_order ASC`
	rows, err := s.db.QueryContext(ctx, query, permitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var approvals []*approvalDomain.Approval
	for rows.Next() {
		a := &approvalDomain.Approval{}
		err = rows.Scan(&a.ID, &a.PermitID, &a.ApproverID, &a.ApproverRole, &a.Decision, &a.Comments, &a.SignatureHash, &a.DecidedAt, &a.SequenceOrder)
		if err != nil {
			return nil, err
		}
		approvals = append(approvals, a)
	}
	return approvals, nil
}

// Update registers edits.
func (s *ApprovalStore) Update(ctx context.Context, a *approvalDomain.Approval) error {
	query := `UPDATE permit_approvals SET decision = $2, comments = $3, signature_hash = $4, decided_at = $5 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, a.ID, a.Decision, a.Comments, a.SignatureHash, a.DecidedAt)
	return err
}
