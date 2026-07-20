package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"prahari/services/pha/internal/domain/actionitem"
	"prahari/services/pha/internal/domain/hazardscenario"
	"prahari/services/pha/internal/domain/hazop"
	"prahari/services/pha/internal/domain/lopa"
	"prahari/services/pha/internal/domain/phastudy"
	"prahari/services/pha/internal/domain/processnode"
	"prahari/services/pha/internal/domain/recommendation"
	"prahari/services/pha/internal/domain/search"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) SaveStudy(ctx context.Context, st *phastudy.Study) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO pha_studies (id, study_number, plant_id, unit_id, title, description, method, moc_id, status, leader_id, scribe_id, target_date, revalidation_due_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
		ON CONFLICT (id) DO UPDATE SET title = EXCLUDED.title, status = EXCLUDED.status, updated_at = EXCLUDED.updated_at`
	_, err := s.db.ExecContext(ctx, query, st.ID, st.StudyNumber, st.PlantID, st.UnitID, st.Title, st.Description, st.Method, st.MOCID, st.Status, st.LeaderID, st.ScribeID, st.TargetDate, st.RevalidationDueAt, st.CreatedAt, st.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to save PHA study: %w", err)
	}
	return nil
}

func (s *Store) GetStudyByID(ctx context.Context, id string) (*phastudy.Study, error) {
	if s.db == nil {
		return &phastudy.Study{ID: id, StudyNumber: "PHA-P01-2001", Title: "Hydrocracker Reaction Unit HAZOP", Method: phastudy.MethodHAZOP, Status: "STUDY", LeaderID: "usr-lead-01"}, nil
	}
	query := `SELECT id, study_number, plant_id, unit_id, title, description, method, COALESCE(moc_id, ''), status, leader_id, scribe_id, target_date, revalidation_due_at, created_at, updated_at FROM pha_studies WHERE id = $1`
	row := s.db.QueryRowContext(ctx, query, id)
	var st phastudy.Study
	if err := row.Scan(&st.ID, &st.StudyNumber, &st.PlantID, &st.UnitID, &st.Title, &st.Description, &st.Method, &st.MOCID, &st.Status, &st.LeaderID, &st.ScribeID, &st.TargetDate, &st.RevalidationDueAt, &st.CreatedAt, &st.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("PHA study %s not found", id)
		}
		return nil, err
	}
	return &st, nil
}

func (s *Store) ListStudies(ctx context.Context, plantID string) ([]*phastudy.Study, error) {
	if s.db == nil {
		return []*phastudy.Study{
			{ID: "pha-001", StudyNumber: "PHA-P01-2001", PlantID: plantID, Title: "Distillation Column Overpressure HAZOP", Method: phastudy.MethodHAZOP, Status: "STUDY"},
		}, nil
	}
	query := `SELECT id, study_number, plant_id, unit_id, title, description, method, COALESCE(moc_id, ''), status, leader_id, scribe_id, target_date, revalidation_due_at, created_at, updated_at FROM pha_studies WHERE plant_id = $1`
	rows, err := s.db.QueryContext(ctx, query, plantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*phastudy.Study
	for rows.Next() {
		var st phastudy.Study
		if err := rows.Scan(&st.ID, &st.StudyNumber, &st.PlantID, &st.UnitID, &st.Title, &st.Description, &st.Method, &st.MOCID, &st.Status, &st.LeaderID, &st.ScribeID, &st.TargetDate, &st.RevalidationDueAt, &st.CreatedAt, &st.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, &st)
	}
	return result, nil
}

func (s *Store) SaveNode(ctx context.Context, node *processnode.Node) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO process_nodes (id, study_id, node_number, node_name, design_intent, p_and_id_number, operating_temp_c, operating_press_bar, location_code, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := s.db.ExecContext(ctx, query, node.ID, node.StudyID, node.NodeNumber, node.NodeName, node.DesignIntent, node.PAndIDNumber, node.OperatingTemp, node.OperatingPress, node.LocationCode, node.CreatedAt)
	return err
}

func (s *Store) SaveScenario(ctx context.Context, sc *hazardscenario.Scenario) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO hazard_scenarios (id, node_id, deviation_id, cause_description, severity, likelihood, risk_rank, risk_category, is_sil_relevant, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := s.db.ExecContext(ctx, query, sc.ID, sc.NodeID, sc.DeviationID, sc.CauseDescription, sc.Severity, sc.Likelihood, sc.RiskRank, sc.RiskCategory, sc.IsSILRelevant, sc.CreatedAt)
	return err
}

func (s *Store) SaveSession(ctx context.Context, sess *hazop.Session) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO hazop_sessions (id, study_id, session_date, duration_hrs, attendees, notes, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, sess.ID, sess.StudyID, sess.SessionDate, sess.DurationHrs, sess.Attendees, sess.Notes, sess.CreatedAt)
	return err
}

func (s *Store) SaveLOPA(ctx context.Context, l *lopa.Analysis) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO lopa_analyses (id, study_id, scenario_id, initiating_event_freq, tolerable_target_freq, total_ipl_mitigation, mitigated_event_freq, required_rrf, target_sil, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := s.db.ExecContext(ctx, query, l.ID, l.StudyID, l.ScenarioID, l.InitiatingEventFreq, l.TolerableTargetFreq, l.TotalIPLmitigation, l.MitigatedEventFreq, l.RequiredRRF, l.TargetSIL, l.CreatedAt)
	return err
}

func (s *Store) SaveRecommendation(ctx context.Context, rec *recommendation.Recommendation) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO recommendations (id, study_id, scenario_id, rec_number, title, description, priority, target_sil, status, assigned_to, target_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	_, err := s.db.ExecContext(ctx, query, rec.ID, rec.StudyID, rec.ScenarioID, rec.RecNumber, rec.Title, rec.Description, rec.Priority, rec.TargetSIL, rec.Status, rec.AssignedTo, rec.TargetDate, rec.CreatedAt, rec.UpdatedAt)
	return err
}

func (s *Store) SaveActionItem(ctx context.Context, item *actionitem.ActionItem) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO action_items (id, recommendation_id, action_title, assignee_id, work_order_id, status, due_date, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := s.db.ExecContext(ctx, query, item.ID, item.RecommendationID, item.ActionTitle, item.AssigneeID, item.WorkOrderID, item.Status, item.DueDate, item.CreatedAt)
	return err
}

func (s *Store) SearchStudies(ctx context.Context, criteria *search.Criteria) ([]*phastudy.Study, int64, error) {
	studies, err := s.ListStudies(ctx, criteria.PlantID)
	if err != nil {
		return nil, 0, err
	}
	return studies, int64(len(studies)), nil
}

func (s *Store) GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	return map[string]float64{
		"active_pha_studies":        8.0,
		"hazop_completion_rate_pct": 95.0,
		"lopa_completion_rate_pct":  92.0,
		"recommendations_open":      24.0,
		"action_closure_rate_pct":   94.5,
		"high_risk_scenario_count":  12.0,
	}, nil
}
