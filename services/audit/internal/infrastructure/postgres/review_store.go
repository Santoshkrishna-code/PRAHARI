package postgres

import (
	"context"
	"database/sql"

	reviewDomain "prahari/services/audit/internal/domain/review"
)

// ReviewStore implements verification checklists reviews database.
type ReviewStore struct {
	db *sql.DB
}

// NewReviewStore instantiates ReviewStore.
func NewReviewStore(db *sql.DB) *ReviewStore {
	return &ReviewStore{db: db}
}

// Create persists review metrics.
func (s *ReviewStore) Create(ctx context.Context, ar *reviewDomain.AuditReview) error {
	query := `INSERT INTO reviews (id, audit_id, reviewer_id, review_date, next_review_date, notes)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, ar.ID, ar.AuditID, ar.ReviewerID, ar.ReviewDate, ar.NextReviewDate, ar.Notes)
	return err
}

// FindByAuditID returns reviews lists.
func (s *ReviewStore) FindByAuditID(ctx context.Context, auditID string) ([]*reviewDomain.AuditReview, error) {
	query := `SELECT id, audit_id, reviewer_id, review_date, next_review_date, notes FROM reviews WHERE audit_id = $1`
	rows, err := s.db.QueryContext(ctx, query, auditID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*reviewDomain.AuditReview
	for rows.Next() {
		ar := &reviewDomain.AuditReview{}
		err = rows.Scan(&ar.ID, &ar.AuditID, &ar.ReviewerID, &ar.ReviewDate, &ar.NextReviewDate, &ar.Notes)
		if err != nil {
			return nil, err
		}
		list = append(list, ar)
	}
	return list, nil
}
