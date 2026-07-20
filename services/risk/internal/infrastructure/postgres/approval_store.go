package postgres

import (
	"context"
	"database/sql"

	approvalDomain "prahari/services/risk/internal/domain/approval"
)

// ApprovalStore implements digital workflow signatures database.
type ApprovalStore struct {
	db *sql.DB
}

// NewApprovalStore instantiates ApprovalStore.
func NewApprovalStore(db *sql.DB) *ApprovalStore {
	return &ApprovalStore{db: db}
}

// Create persists signature.
func (s *ApprovalStore) Create(ctx context.Context, a *approvalDomain.Approval) error {
	query := `INSERT INTO risk_approvals (id, risk_id, approver_id, approved_date, signature)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, a.ID, a.RiskID, a.ApproverID, a.ApprovedDate, a.Signature)
	return err
}

// FindByRiskID returns approvals.
func (s *ApprovalStore) FindByRiskID(ctx context.Context, riskID string) ([]*approvalDomain.Approval, error) {
	query := `SELECT id, risk_id, approver_id, approved_date, signature FROM risk_approvals WHERE risk_id = $1`
	rows, err := s.db.QueryContext(ctx, query, riskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*approvalDomain.Approval
	for rows.Next() {
		a := &approvalDomain.Approval{}
		err = rows.Scan(&a.ID, &a.RiskID, &a.ApproverID, &a.ApprovedDate, &a.Signature)
		if err != nil {
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}
