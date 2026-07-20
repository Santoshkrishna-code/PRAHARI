package postgres

import (
	"context"
	"database/sql"

	reviewDomain "prahari/services/compliance/internal/domain/review"
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
func (s *ReviewStore) Create(ctx context.Context, cr *reviewDomain.ComplianceReview) error {
	query := `INSERT INTO reviews (id, compliance_id, reviewer_id, review_date, next_review_date, notes)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, cr.ID, cr.ComplianceID, cr.ReviewerID, cr.ReviewDate, cr.NextReviewDate, cr.Notes)
	return err
}

// FindByComplianceID returns reviews lists.
func (s *ReviewStore) FindByComplianceID(ctx context.Context, complianceID string) ([]*reviewDomain.ComplianceReview, error) {
	query := `SELECT id, compliance_id, reviewer_id, review_date, next_review_date, notes FROM reviews WHERE compliance_id = $1`
	rows, err := s.db.QueryContext(ctx, query, complianceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*reviewDomain.ComplianceReview
	for rows.Next() {
		cr := &reviewDomain.ComplianceReview{}
		err = rows.Scan(&cr.ID, &cr.ComplianceID, &cr.ReviewerID, &cr.ReviewDate, &cr.NextReviewDate, &cr.Notes)
		if err != nil {
			return nil, err
		}
		list = append(list, cr)
	}
	return list, nil
}
