package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"prahari/services/visitor/internal/domain/checkin"
	"prahari/services/visitor/internal/domain/checkout"
	"prahari/services/visitor/internal/domain/emergencymuster"
	"prahari/services/visitor/internal/domain/search"
	"prahari/services/visitor/internal/domain/visit"
	"prahari/services/visitor/internal/domain/visitor"
)


type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) SaveVisitor(ctx context.Context, vis *visitor.Visitor) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO visitors (id, first_name, last_name, email, phone, company, visitor_type, id_type, id_number, blacklisted, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		ON CONFLICT (id) DO UPDATE SET first_name = EXCLUDED.first_name, last_name = EXCLUDED.last_name, blacklisted = EXCLUDED.blacklisted, updated_at = EXCLUDED.updated_at`
	_, err := s.db.ExecContext(ctx, query, vis.ID, vis.FirstName, vis.LastName, vis.Email, vis.Phone, vis.Company, vis.VisitorType, vis.IDType, vis.IDNumber, vis.Blacklisted, vis.CreatedAt, vis.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to save visitor: %w", err)
	}
	return nil
}

func (s *Store) GetVisitorByID(ctx context.Context, id string) (*visitor.Visitor, error) {
	if s.db == nil {
		return &visitor.Visitor{ID: id, FirstName: "John", LastName: "Doe", Email: "john.doe@example.com", Company: "Shell Inc", VisitorType: "CONTRACTOR", IDType: "PASSPORT", IDNumber: "PP-100234"}, nil
	}
	query := `SELECT id, first_name, last_name, email, phone, company, visitor_type, id_type, id_number, blacklisted, created_at, updated_at FROM visitors WHERE id = $1`
	row := s.db.QueryRowContext(ctx, query, id)
	var vis visitor.Visitor
	if err := row.Scan(&vis.ID, &vis.FirstName, &vis.LastName, &vis.Email, &vis.Phone, &vis.Company, &vis.VisitorType, &vis.IDType, &vis.IDNumber, &vis.Blacklisted, &vis.CreatedAt, &vis.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("visitor %s not found", id)
		}
		return nil, err
	}
	return &vis, nil
}

func (s *Store) SaveVisit(ctx context.Context, v *visit.Visit) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO visits (id, visitor_id, host_id, plant_id, purpose, scheduled_in, scheduled_out, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		ON CONFLICT (id) DO UPDATE SET status = EXCLUDED.status, updated_at = EXCLUDED.updated_at`
	_, err := s.db.ExecContext(ctx, query, v.ID, v.VisitorID, v.HostID, v.PlantID, v.Purpose, v.ScheduledIn, v.ScheduledOut, v.Status, v.CreatedAt, v.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to save visit: %w", err)
	}
	return nil
}

func (s *Store) GetVisitByID(ctx context.Context, id string) (*visit.Visit, error) {
	if s.db == nil {
		return &visit.Visit{ID: id, VisitorID: "vis-01", HostID: "usr-host-01", PlantID: "P01", Purpose: "Asset maintenance audit", Status: "SCHEDULED"}, nil
	}
	query := `SELECT id, visitor_id, host_id, plant_id, purpose, scheduled_in, scheduled_out, status, created_at, updated_at FROM visits WHERE id = $1`
	row := s.db.QueryRowContext(ctx, query, id)
	var v visit.Visit
	if err := row.Scan(&v.ID, &v.VisitorID, &v.HostID, &v.PlantID, &v.Purpose, &v.ScheduledIn, &v.ScheduledOut, &v.Status, &v.CreatedAt, &v.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("visit %s not found", id)
		}
		return nil, err
	}
	return &v, nil
}

func (s *Store) ListVisits(ctx context.Context, plantID string) ([]*visit.Visit, error) {
	if s.db == nil {
		return []*visit.Visit{
			{ID: "vst-001", VisitorID: "vis-01", HostID: "usr-host-01", PlantID: plantID, Purpose: "Asset audit", Status: "SCHEDULED"},
		}, nil
	}
	query := `SELECT id, visitor_id, host_id, plant_id, purpose, scheduled_in, scheduled_out, status, created_at, updated_at FROM visits WHERE plant_id = $1`
	rows, err := s.db.QueryContext(ctx, query, plantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*visit.Visit
	for rows.Next() {
		var v visit.Visit
		if err := rows.Scan(&v.ID, &v.VisitorID, &v.HostID, &v.PlantID, &v.Purpose, &v.ScheduledIn, &v.ScheduledOut, &v.Status, &v.CreatedAt, &v.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, &v)
	}
	return result, nil
}

func (s *Store) SaveCheckin(ctx context.Context, rec *checkin.Record) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO checkins (id, visit_id, security_check_point, gate_number, check_in_at, checked_in_by)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, rec.ID, rec.VisitID, rec.SecurityCheckPoint, rec.GateNumber, rec.CheckInAt, rec.CheckedInBy)
	return err
}

func (s *Store) SaveCheckout(ctx context.Context, rec *checkout.Record) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO checkouts (id, visit_id, check_out_at, checked_out_by, badge_returned)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, rec.ID, rec.VisitID, rec.CheckOutAt, rec.CheckedOutBy, rec.BadgeReturned)
	return err
}

func (s *Store) SaveMusterRecord(ctx context.Context, rec *emergencymuster.Record) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO emergencymuster_records (id, muster_id, visitor_id, assembly_point, accounted_for, accounted_at, warden_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, rec.ID, rec.MusterID, rec.VisitorID, rec.AssemblyPoint, rec.AccountedFor, rec.AccountedAt, rec.WardenID)
	return err
}

func (s *Store) SearchVisits(ctx context.Context, criteria *search.Criteria) ([]*visit.Visit, int64, error) {
	visits, err := s.ListVisits(ctx, criteria.PlantID)
	if err != nil {
		return nil, 0, err
	}
	return visits, int64(len(visits)), nil
}

func (s *Store) GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	return map[string]float64{
		"active_visitors":              42.0,
		"average_visit_duration":       4.5,
		"checkin_compliance_pct":       99.2,
		"badge_return_compliance_pct":  99.8,
		"induction_compliance_pct":     100.0,
	}, nil
}
