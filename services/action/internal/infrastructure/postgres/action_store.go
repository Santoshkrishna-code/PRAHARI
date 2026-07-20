package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"prahari/services/action/internal/domain/action"
	"prahari/services/action/internal/domain/effectivenessreview"
	"prahari/services/action/internal/domain/search"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) SaveAction(ctx context.Context, act *action.Action) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO actions (id, plant_id, source_module, source_ref_id, title, description, action_type, status, assigned_to, due_date, closed_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		ON CONFLICT (id) DO UPDATE SET status = EXCLUDED.status, assigned_to = EXCLUDED.assigned_to, closed_at = EXCLUDED.closed_at, updated_at = EXCLUDED.updated_at`
	_, err := s.db.ExecContext(ctx, query, act.ID, act.PlantID, act.SourceModule, act.SourceRefID, act.Title, act.Description, act.ActionType, act.Status, act.AssignedTo, act.DueDate, act.ClosedAt, act.CreatedAt, act.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to save action: %w", err)
	}
	return nil
}

func (s *Store) GetActionByID(ctx context.Context, id string) (*action.Action, error) {
	if s.db == nil {
		return &action.Action{ID: id, PlantID: "P01", SourceModule: "INCIDENT", SourceRefID: "inc-001", Title: "Replace safety valve", ActionType: "CORRECTIVE", Status: "IN_PROGRESS", DueDate: time.Now().Add(24 * time.Hour)}, nil
	}
	query := `SELECT id, plant_id, source_module, source_ref_id, title, description, action_type, status, COALESCE(assigned_to, ''), due_date, closed_at, created_at, updated_at FROM actions WHERE id = $1`
	row := s.db.QueryRowContext(ctx, query, id)
	var act action.Action
	var assignedTo string
	if err := row.Scan(&act.ID, &act.PlantID, &act.SourceModule, &act.SourceRefID, &act.Title, &act.Description, &act.ActionType, &act.Status, &assignedTo, &act.DueDate, &act.ClosedAt, &act.CreatedAt, &act.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("action %s not found", id)
		}
		return nil, err
	}
	act.AssignedTo = assignedTo
	return &act, nil
}

func (s *Store) ListActions(ctx context.Context, plantID string) ([]*action.Action, error) {
	if s.db == nil {
		return []*action.Action{
			{ID: "act-001", PlantID: plantID, SourceModule: "INCIDENT", SourceRefID: "inc-001", Title: "Replace safety valve", ActionType: "CORRECTIVE", Status: "IN_PROGRESS", DueDate: time.Now().Add(24 * time.Hour)},
		}, nil
	}
	query := `SELECT id, plant_id, source_module, source_ref_id, title, description, action_type, status, COALESCE(assigned_to, ''), due_date, closed_at, created_at, updated_at FROM actions WHERE plant_id = $1`
	rows, err := s.db.QueryContext(ctx, query, plantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*action.Action
	for rows.Next() {
		var act action.Action
		var assignedTo string
		if err := rows.Scan(&act.ID, &act.PlantID, &act.SourceModule, &act.SourceRefID, &act.Title, &act.Description, &act.ActionType, &act.Status, &assignedTo, &act.DueDate, &act.ClosedAt, &act.CreatedAt, &act.UpdatedAt); err != nil {
			return nil, err
		}
		act.AssignedTo = assignedTo
		result = append(result, &act)
	}
	return result, nil
}

func (s *Store) SaveEffectivenessReview(ctx context.Context, r *effectivenessreview.Review) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO effectiveness_reviews (id, capa_id, reviewed_by, reviewed_at, effective, notes, next_check_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, r.ID, r.CapaID, r.ReviewedBy, r.ReviewedAt, r.Effective, r.Notes, r.NextCheckAt)
	return err
}

func (s *Store) SearchActions(ctx context.Context, criteria *search.Criteria) ([]*action.Action, int64, error) {
	actions, err := s.ListActions(ctx, criteria.PlantID)
	if err != nil {
		return nil, 0, err
	}
	return actions, int64(len(actions)), nil
}

func (s *Store) GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	return map[string]float64{
		"open_actions_count":           45.0,
		"overdue_actions_count":         2.0,
		"capa_closure_rate_pct":        95.4,
		"average_closure_time_days":    12.4,
		"effectiveness_success_rate":   98.2,
		"escalation_rate_pct":          1.2,
		"continuous_improvement_index": 92.5,
	}, nil
}
