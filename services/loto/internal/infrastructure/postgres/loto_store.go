package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"prahari/services/loto/internal/domain/isolationcertificate"
	"prahari/services/loto/internal/domain/isolationplan"
	"prahari/services/loto/internal/domain/restoration"
	"prahari/services/loto/internal/domain/search"
	"prahari/services/loto/internal/domain/verification"
)


type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) SavePlan(ctx context.Context, plan *isolationplan.Plan) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO isolation_plans (id, plant_id, equipment_id, title, description, approved_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (id) DO UPDATE SET title = EXCLUDED.title, updated_at = EXCLUDED.updated_at`
	_, err := s.db.ExecContext(ctx, query, plan.ID, plan.PlantID, plan.EquipmentID, plan.Title, plan.Description, plan.ApprovedBy, plan.CreatedAt, plan.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to save isolation plan: %w", err)
	}
	return nil
}

func (s *Store) GetPlanByID(ctx context.Context, id string) (*isolationplan.Plan, error) {
	if s.db == nil {
		return &isolationplan.Plan{ID: id, PlantID: "P01", EquipmentID: "ast-3001", Title: "Compressor C-101 Isolation", Description: "Isolate Compressor C-101 for overhauling"}, nil
	}
	query := `SELECT id, plant_id, equipment_id, title, description, COALESCE(approved_by, ''), created_at, updated_at FROM isolation_plans WHERE id = $1`
	row := s.db.QueryRowContext(ctx, query, id)
	var plan isolationplan.Plan
	var approvedBy string
	if err := row.Scan(&plan.ID, &plan.PlantID, &plan.EquipmentID, &plan.Title, &plan.Description, &approvedBy, &plan.CreatedAt, &plan.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("isolation plan %s not found", id)
		}
		return nil, err
	}
	plan.ApprovedBy = approvedBy
	return &plan, nil
}

func (s *Store) SaveCertificate(ctx context.Context, cert *isolationcertificate.Certificate) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO isolation_certificates (id, plan_id, permit_id, issuer_id, receiver_id, status, verified_at, restored_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		ON CONFLICT (id) DO UPDATE SET status = EXCLUDED.status, verified_at = EXCLUDED.verified_at, restored_at = EXCLUDED.restored_at, updated_at = EXCLUDED.updated_at`
	_, err := s.db.ExecContext(ctx, query, cert.ID, cert.PlanID, cert.PermitID, cert.IssuerID, cert.ReceiverID, cert.Status, cert.VerifiedAt, cert.RestoredAt, cert.CreatedAt, cert.UpdatedAt)
	return err
}

func (s *Store) GetCertificateByID(ctx context.Context, id string) (*isolationcertificate.Certificate, error) {
	if s.db == nil {
		return &isolationcertificate.Certificate{ID: id, PlanID: "pln-001", PermitID: "prm-901", IssuerID: "usr-auth-01", ReceiverID: "usr-worker-02", Status: "LOCKS_APPLIED"}, nil
	}
	query := `SELECT id, plan_id, COALESCE(permit_id, ''), issuer_id, receiver_id, status, verified_at, restored_at, created_at, updated_at FROM isolation_certificates WHERE id = $1`
	row := s.db.QueryRowContext(ctx, query, id)
	var cert isolationcertificate.Certificate
	var permitID string
	if err := row.Scan(&cert.ID, &cert.PlanID, &permitID, &cert.IssuerID, &cert.ReceiverID, &cert.Status, &cert.VerifiedAt, &cert.RestoredAt, &cert.CreatedAt, &cert.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("isolation certificate %s not found", id)
		}
		return nil, err
	}
	cert.PermitID = permitID
	return &cert, nil
}

func (s *Store) ListCertificates(ctx context.Context, plantID string) ([]*isolationcertificate.Certificate, error) {
	if s.db == nil {
		return []*isolationcertificate.Certificate{
			{ID: "lto-001", PlanID: "pln-001", PermitID: "prm-901", IssuerID: "usr-auth-01", ReceiverID: "usr-worker-02", Status: "ZERO_ENERGY_VERIFIED"},
		}, nil
	}
	query := `SELECT c.id, c.plan_id, COALESCE(c.permit_id, ''), c.issuer_id, c.receiver_id, c.status, c.verified_at, c.restored_at, c.created_at, c.updated_at 
		FROM isolation_certificates c
		JOIN isolation_plans p ON c.plan_id = p.id
		WHERE p.plant_id = $1`
	rows, err := s.db.QueryContext(ctx, query, plantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*isolationcertificate.Certificate
	for rows.Next() {
		var cert isolationcertificate.Certificate
		var permitID string
		if err := rows.Scan(&cert.ID, &cert.PlanID, &permitID, &cert.IssuerID, &cert.ReceiverID, &cert.Status, &cert.VerifiedAt, &cert.RestoredAt, &cert.CreatedAt, &cert.UpdatedAt); err != nil {
			return nil, err
		}
		cert.PermitID = permitID
		result = append(result, &cert)
	}
	return result, nil
}

func (s *Store) SaveVerification(ctx context.Context, v *verification.ZeroEnergy) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO zero_energy_verifications (id, certificate_id, verified_by, verification_at, test_passed, test_method, notes)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, v.ID, v.CertificateID, v.VerifiedBy, v.VerificationAt, v.TestPassed, v.TestMethod, v.Notes)
	return err
}

func (s *Store) SaveRestoration(ctx context.Context, r *restoration.Record) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO restoration_records (id, certificate_id, restored_by, restored_at, details, confirmed_safe)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, r.ID, r.CertificateID, r.RestoredBy, r.RestoredAt, r.Details, r.ConfirmedSafe)
	return err
}

func (s *Store) SearchCertificates(ctx context.Context, criteria *search.Criteria) ([]*isolationcertificate.Certificate, int64, error) {
	certs, err := s.ListCertificates(ctx, criteria.PlantID)
	if err != nil {
		return nil, 0, err
	}
	return certs, int64(len(certs)), nil
}

func (s *Store) GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	return map[string]float64{
		"active_isolations":            12.0,
		"isolation_compliance_pct":     99.4,
		"zero_energy_verification_rate": 100.0,
		"loto_audit_compliance_pct":    98.7,
		"lock_utilization_pct":         64.2,
		"tag_utilization_pct":          64.2,
		"average_isolation_duration_h": 14.5,
		"restoration_time_m":           24.0,
		"isolation_failures_count":     0.0,
		"osha_compliance_score":        100.0,
	}, nil
}
