package postgres

import (
	"context"
	"database/sql"

	actionDomain "prahari/services/audit/internal/domain/correctiveaction"
)

// ActionStore implements CAPA corrective action plan tasks.
type ActionStore struct {
	db *sql.DB
}

// NewActionStore instantiates ActionStore.
func NewActionStore(db *sql.DB) *ActionStore {
	return &ActionStore{db: db}
}

// Create persists corrective actions.
func (s *ActionStore) Create(ctx context.Context, ca *actionDomain.CorrectiveAction) error {
	query := `INSERT INTO corrective_actions (id, finding_id, description, target_date, is_completed)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, ca.ID, ca.FindingID, ca.Description, ca.TargetDate, ca.IsCompleted)
	return err
}

// FindByFindingID returns corrective plans checklists.
func (s *ActionStore) FindByFindingID(ctx context.Context, findingID string) ([]*actionDomain.CorrectiveAction, error) {
	query := `SELECT id, finding_id, description, target_date, is_completed FROM corrective_actions WHERE finding_id = $1`
	rows, err := s.db.QueryContext(ctx, query, findingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*actionDomain.CorrectiveAction
	for rows.Next() {
		ca := &actionDomain.CorrectiveAction{}
		err = rows.Scan(&ca.ID, &ca.FindingID, &ca.Description, &ca.TargetDate, &ca.IsCompleted)
		if err != nil {
			return nil, err
		}
		list = append(list, ca)
	}
	return list, nil
}
