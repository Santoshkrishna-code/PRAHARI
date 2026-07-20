package postgres

import (
	"context"
	"database/sql"

	actionPlanDomain "prahari/services/compliance/internal/domain/actionplan"
)

// ActionStore implements corrective task schedulers database.
type ActionStore struct {
	db *sql.DB
}

// NewActionStore instantiates ActionStore.
func NewActionStore(db *sql.DB) *ActionStore {
	return &ActionStore{db: db}
}

// Create persists corrective actions.
func (s *ActionStore) Create(ctx context.Context, a *actionPlanDomain.ActionPlan) error {
	query := `INSERT INTO action_plans (id, finding_id, description, target_date, is_completed)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, a.ID, a.FindingID, a.Description, a.TargetDate, a.IsCompleted)
	return err
}

// FindByFindingID returns action checklists.
func (s *ActionStore) FindByFindingID(ctx context.Context, findingID string) ([]*actionPlanDomain.ActionPlan, error) {
	query := `SELECT id, finding_id, description, target_date, is_completed FROM action_plans WHERE finding_id = $1`
	rows, err := s.db.QueryContext(ctx, query, findingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*actionPlanDomain.ActionPlan
	for rows.Next() {
		a := &actionPlanDomain.ActionPlan{}
		err = rows.Scan(&a.ID, &a.FindingID, &a.Description, &a.TargetDate, &a.IsCompleted)
		if err != nil {
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}
