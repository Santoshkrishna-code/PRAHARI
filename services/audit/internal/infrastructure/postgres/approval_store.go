package postgres

import (
	"context"
	"database/sql"

	approvalDomain "prahari/services/audit/internal/domain/approval"
)

// ApprovalStore implements workflow digital signatures approvals database.
type ApprovalStore struct {
	db *sql.DB
}

// NewApprovalStore instantiates ApprovalStore.
func NewApprovalStore(db *sql.DB) *ApprovalStore {
	return &ApprovalStore{db: db}
}

// Create persists digital verification signature metadata.
func (s *ApprovalStore) Create(ctx context.Context, a *approvalDomain.Approval) error {
	query := `INSERT INTO approvals (id, audit_id, approver_id, approved_date, signature)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, a.ID, a.AuditID, a.ApproverID, a.ApprovedDate, a.Signature)
	return err
}

// FindByAuditID returns approval signatures.
func (s *ApprovalStore) FindByAuditID(ctx context.Context, auditID string) ([]*approvalDomain.Approval, error) {
	query := `SELECT id, audit_id, approver_id, approved_date, signature FROM approvals WHERE audit_id = $1`
	rows, err := s.db.QueryContext(ctx, query, auditID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*approvalDomain.Approval
	for rows.Next() {
		a := &approvalDomain.Approval{}
		err = rows.Scan(&a.ID, &a.AuditID, &a.ApproverID, &a.ApprovedDate, &a.Signature)
		if err != nil {
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}
