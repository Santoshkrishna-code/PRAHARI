package postgres

import (
	"context"
	"database/sql"

	planDomain "prahari/services/audit/internal/domain/auditplan"
)

// PlanStore implements audit scheduling timelines.
type PlanStore struct {
	db *sql.DB
}

// NewPlanStore instantiates PlanStore.
func NewPlanStore(db *sql.DB) *PlanStore {
	return &PlanStore{db: db}
}

// Create persists audit plan boundaries.
func (s *PlanStore) Create(ctx context.Context, ap *planDomain.AuditPlan) error {
	query := `INSERT INTO audit_plans (id, audit_program_id, scheduled_start, scheduled_end)
		VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, ap.ID, ap.AuditProgramID, ap.ScheduledStart, ap.ScheduledEnd)
	return err
}
