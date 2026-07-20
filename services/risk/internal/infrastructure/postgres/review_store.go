package postgres

import (
	"context"
	"database/sql"

	reviewDomain "prahari/services/risk/internal/domain/review"
)

// ReviewStore implements periodic reassessment reviewer logs database.
type ReviewStore struct {
	db *sql.DB
}

// NewReviewStore instantiates ReviewStore.
func NewReviewStore(db *sql.DB) *ReviewStore {
	return &ReviewStore{db: db}
}

// Create persists evaluation.
func (s *ReviewStore) Create(ctx context.Context, r *reviewDomain.RiskReview) error {
	query := `INSERT INTO risk_reviews (id, risk_id, reviewer_id, review_date, next_review_date, notes)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, r.ID, r.RiskID, r.ReviewerID, r.ReviewDate, r.NextReviewDate, r.Notes)
	return err
}

// FindByRiskID returns reviews list.
func (s *ReviewStore) FindByRiskID(ctx context.Context, riskID string) ([]*reviewDomain.RiskReview, error) {
	query := `SELECT id, risk_id, reviewer_id, review_date, next_review_date, notes FROM risk_reviews WHERE risk_id = $1`
	rows, err := s.db.QueryContext(ctx, query, riskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*reviewDomain.RiskReview
	for rows.Next() {
		r := &reviewDomain.RiskReview{}
		err = rows.Scan(&r.ID, &r.RiskID, &r.ReviewerID, &r.ReviewDate, &r.NextReviewDate, &r.Notes)
		if err != nil {
			return nil, err
		}
		list = append(list, r)
	}
	return list, nil
}
