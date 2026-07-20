package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"prahari/services/moc/internal/domain/approval"
	"prahari/services/moc/internal/domain/changerequest"
	"prahari/services/moc/internal/domain/impactassessment"
	"prahari/services/moc/internal/domain/implementation"
	"prahari/services/moc/internal/domain/riskreview"
	"prahari/services/moc/internal/domain/rollback"
	"prahari/services/moc/internal/domain/safetyreview"
	"prahari/services/moc/internal/domain/search"
	"prahari/services/moc/internal/domain/technicalreview"
	"prahari/services/moc/internal/domain/verification"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) SaveRequest(ctx context.Context, req *changerequest.Request) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO change_requests (id, moc_number, plant_id, department_id, title, description, reason_for_change, category, change_type, target_asset_id, risk_level, status, requester_id, target_date, expiry_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
		ON CONFLICT (id) DO UPDATE SET title = EXCLUDED.title, status = EXCLUDED.status, updated_at = EXCLUDED.updated_at`
	_, err := s.db.ExecContext(ctx, query, req.ID, req.MOCNumber, req.PlantID, req.DepartmentID, req.Title, req.Description, req.ReasonForChange, req.Category, req.ChangeType, req.TargetAssetID, req.RiskLevel, req.Status, req.RequesterID, req.TargetDate, req.ExpiryDate, req.CreatedAt, req.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to save change request: %w", err)
	}
	return nil
}

func (s *Store) GetRequestByID(ctx context.Context, id string) (*changerequest.Request, error) {
	if s.db == nil {
		return &changerequest.Request{ID: id, MOCNumber: "MOC-P01-1001", Title: "Mock Pump impeller Modification", Category: changerequest.CategoryPermanent, ChangeType: "MECHANICAL", RiskLevel: "HIGH", Status: "DRAFT", RequesterID: "usr-01"}, nil
	}
	query := `SELECT id, moc_number, plant_id, department_id, title, description, reason_for_change, category, change_type, COALESCE(target_asset_id, ''), risk_level, status, requester_id, target_date, expiry_date, created_at, updated_at FROM change_requests WHERE id = $1`
	row := s.db.QueryRowContext(ctx, query, id)
	var req changerequest.Request
	if err := row.Scan(&req.ID, &req.MOCNumber, &req.PlantID, &req.DepartmentID, &req.Title, &req.Description, &req.ReasonForChange, &req.Category, &req.ChangeType, &req.TargetAssetID, &req.RiskLevel, &req.Status, &req.RequesterID, &req.TargetDate, &req.ExpiryDate, &req.CreatedAt, &req.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("change request %s not found", id)
		}
		return nil, err
	}
	return &req, nil
}

func (s *Store) ListRequests(ctx context.Context, plantID string) ([]*changerequest.Request, error) {
	if s.db == nil {
		return []*changerequest.Request{
			{ID: "moc-001", MOCNumber: "MOC-P01-1001", PlantID: plantID, Title: "Safety Valve Setpoint Adjustment", Category: changerequest.CategoryPermanent, ChangeType: "INSTRUMENTATION", RiskLevel: "HIGH", Status: "APPROVAL"},
		}, nil
	}
	query := `SELECT id, moc_number, plant_id, department_id, title, description, reason_for_change, category, change_type, COALESCE(target_asset_id, ''), risk_level, status, requester_id, target_date, expiry_date, created_at, updated_at FROM change_requests WHERE plant_id = $1`
	rows, err := s.db.QueryContext(ctx, query, plantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*changerequest.Request
	for rows.Next() {
		var req changerequest.Request
		if err := rows.Scan(&req.ID, &req.MOCNumber, &req.PlantID, &req.DepartmentID, &req.Title, &req.Description, &req.ReasonForChange, &req.Category, &req.ChangeType, &req.TargetAssetID, &req.RiskLevel, &req.Status, &req.RequesterID, &req.TargetDate, &req.ExpiryDate, &req.CreatedAt, &req.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, &req)
	}
	return result, nil
}

func (s *Store) SaveImpactAssessment(ctx context.Context, ia *impactassessment.Assessment) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO impact_assessments (id, change_request_id, safety_impact, environmental_impact, quality_impact, reliability_impact, cybersecurity_impact, regulatory_impact, p_and_id_impact, hazop_required, summary_notes, assessed_by, assessed_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	_, err := s.db.ExecContext(ctx, query, ia.ID, ia.ChangeRequestID, ia.SafetyImpact, ia.EnvironmentalImpact, ia.QualityImpact, ia.ReliabilityImpact, ia.CybersecurityImpact, ia.RegulatoryImpact, ia.PAndIDImpact, ia.HAZOPRequired, ia.SummaryNotes, ia.AssessedBy, ia.AssessedAt)
	return err
}

func (s *Store) SaveTechnicalReview(ctx context.Context, tr *technicalreview.Review) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO technical_reviews (id, change_request_id, discipline, reviewer_id, status, findings, conditions, reviewed_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := s.db.ExecContext(ctx, query, tr.ID, tr.ChangeRequestID, tr.Discipline, tr.ReviewerID, tr.Status, tr.Findings, tr.Conditions, tr.ReviewedAt)
	return err
}

func (s *Store) SaveRiskReview(ctx context.Context, rr *riskreview.Review) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO risk_reviews (id, change_request_id, risk_assessment_id, pre_change_risk, post_change_risk, mitigations_reqd, risk_manager_id, status, reviewed_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := s.db.ExecContext(ctx, query, rr.ID, rr.ChangeRequestID, rr.RiskAssessmentID, rr.PreChangeRisk, rr.PostChangeRisk, rr.MitigationsReqd, rr.RiskManagerID, rr.Status, rr.ReviewedAt)
	return err
}

func (s *Store) SaveSafetyReview(ctx context.Context, sr *safetyreview.Review) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO safety_reviews (id, change_request_id, psm_impact_verified, occupational_health, emergency_response, safety_officer_id, status, comments, reviewed_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := s.db.ExecContext(ctx, query, sr.ID, sr.ChangeRequestID, sr.PSMImpactVerified, sr.OccupationalHealth, sr.EmergencyResponse, sr.SafetyOfficerID, sr.Status, sr.Comments, sr.ReviewedAt)
	return err
}

func (s *Store) SaveApproval(ctx context.Context, app *approval.Record) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO approvals (id, change_request_id, approver_id, role, decision, comments, approved_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, app.ID, app.ChangeRequestID, app.ApproverID, app.Role, app.Decision, app.Comments, app.ApprovedAt)
	return err
}

func (s *Store) SaveImplementation(ctx context.Context, plan *implementation.Plan) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO implementations (id, change_request_id, work_order_id, permit_id, implemented_by, status, start_date, notes, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := s.db.ExecContext(ctx, query, plan.ID, plan.ChangeRequestID, plan.WorkOrderID, plan.PermitID, plan.ImplementedBy, plan.Status, plan.StartDate, plan.Notes, plan.CreatedAt)
	return err
}

func (s *Store) SaveVerification(ctx context.Context, v *verification.Record) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO verifications (id, change_request_id, pssr_completed, training_verified, docs_updated, verified_by, status, comments, verified_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := s.db.ExecContext(ctx, query, v.ID, v.ChangeRequestID, v.PSSRCompleted, v.TrainingVerified, v.DocsUpdated, v.VerifiedBy, v.Status, v.Comments, v.VerifiedAt)
	return err
}

func (s *Store) SaveRollback(ctx context.Context, plan *rollback.Plan) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO rollback_plans (id, change_request_id, trigger_reason, reversion_steps, executed_by, status, executed_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := s.db.ExecContext(ctx, query, plan.ID, plan.ChangeRequestID, plan.TriggerReason, plan.ReversionSteps, plan.ExecutedBy, plan.Status, plan.ExecutedAt, plan.CreatedAt)
	return err
}

func (s *Store) SearchRequests(ctx context.Context, criteria *search.Criteria) ([]*changerequest.Request, int64, error) {
	requests, err := s.ListRequests(ctx, criteria.PlantID)
	if err != nil {
		return nil, 0, err
	}
	return requests, int64(len(requests)), nil
}

func (s *Store) GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	return map[string]float64{
		"open_changes":               14.0,
		"avg_approval_cycle_hours":   36.5,
		"emergency_changes":          2.0,
		"rollback_rate_pct":          1.5,
		"pssr_completion_rate_pct":   98.2,
		"training_verif_rate_pct":    96.0,
	}, nil
}
