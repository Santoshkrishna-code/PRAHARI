package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"prahari/services/ppe/internal/domain/ppe"
	"prahari/services/ppe/internal/domain/ppeinspection"
	"prahari/services/ppe/internal/domain/ppeitem"
	"prahari/services/ppe/internal/domain/ppeissue"
	"prahari/services/ppe/internal/domain/ppemaintenance"
	"prahari/services/ppe/internal/domain/search"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) SavePPE(ctx context.Context, p *ppe.PPE) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO ppe_catalog (id, plant_id, model_name, category_id, manufacturer, part_number, standard_ref, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		ON CONFLICT (id) DO UPDATE SET model_name = EXCLUDED.model_name, updated_at = EXCLUDED.updated_at`
	_, err := s.db.ExecContext(ctx, query, p.ID, p.PlantID, p.ModelName, p.CategoryID, p.Manufacturer, p.PartNumber, p.StandardRef, p.Description, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to save ppe catalog: %w", err)
	}
	return nil
}

func (s *Store) GetPPEByID(ctx context.Context, id string) (*ppe.PPE, error) {
	if s.db == nil {
		return &ppe.PPE{ID: id, PlantID: "P01", ModelName: "Standard Hard Hat", CategoryID: "HEAD", Manufacturer: "3M", PartNumber: "H701R", StandardRef: "ANSI Z89.1"}, nil
	}
	query := `SELECT id, plant_id, model_name, category_id, manufacturer, part_number, standard_ref, description, created_at, updated_at FROM ppe_catalog WHERE id = $1`
	row := s.db.QueryRowContext(ctx, query, id)
	var p ppe.PPE
	if err := row.Scan(&p.ID, &p.PlantID, &p.ModelName, &p.CategoryID, &p.Manufacturer, &p.PartNumber, &p.StandardRef, &p.Description, &p.CreatedAt, &p.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("ppe catalog %s not found", id)
		}
		return nil, err
	}
	return &p, nil
}

func (s *Store) ListPPEs(ctx context.Context, plantID string) ([]*ppe.PPE, error) {
	if s.db == nil {
		return []*ppe.PPE{
			{ID: "ppe-001", PlantID: plantID, ModelName: "Standard Hard Hat", CategoryID: "HEAD", Manufacturer: "3M", PartNumber: "H701R", StandardRef: "ANSI Z89.1"},
		}, nil
	}
	query := `SELECT id, plant_id, model_name, category_id, manufacturer, part_number, standard_ref, description, created_at, updated_at FROM ppe_catalog WHERE plant_id = $1`
	rows, err := s.db.QueryContext(ctx, query, plantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*ppe.PPE
	for rows.Next() {
		var p ppe.PPE
		if err := rows.Scan(&p.ID, &p.PlantID, &p.ModelName, &p.CategoryID, &p.Manufacturer, &p.PartNumber, &p.StandardRef, &p.Description, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, &p)
	}
	return result, nil
}

func (s *Store) SavePPEItem(ctx context.Context, item *ppeitem.Item) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO ppe_items (id, ppe_id, serial_number, rfid_code, barcode, manufacture_date, expiry_date, status, issued_to, last_inspected_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (id) DO UPDATE SET status = EXCLUDED.status, issued_to = EXCLUDED.issued_to, last_inspected_at = EXCLUDED.last_inspected_at`
	_, err := s.db.ExecContext(ctx, query, item.ID, item.PPEID, item.SerialNumber, item.RFIDCode, item.Barcode, item.ManufactureDate, item.ExpiryDate, item.Status, item.IssuedTo, item.LastInspectedAt, item.CreatedAt)
	return err
}

func (s *Store) GetPPEItemByID(ctx context.Context, id string) (*ppeitem.Item, error) {
	if s.db == nil {
		now := time.Now()
		return &ppeitem.Item{ID: id, PPEID: "ppe-001", SerialNumber: "SN-900231", RFIDCode: "RF-80123", Barcode: "BAR-10023", ManufactureDate: now.Add(-365 * 24 * time.Hour), ExpiryDate: now.Add(365 * 24 * time.Hour), Status: "AVAILABLE"}, nil
	}
	query := `SELECT id, ppe_id, serial_number, rfid_code, barcode, manufacture_date, expiry_date, status, COALESCE(issued_to, ''), last_inspected_at, created_at FROM ppe_items WHERE id = $1`
	row := s.db.QueryRowContext(ctx, query, id)
	var item ppeitem.Item
	var issuedTo string
	if err := row.Scan(&item.ID, &item.PPEID, &item.SerialNumber, &item.RFIDCode, &item.Barcode, &item.ManufactureDate, &item.ExpiryDate, &item.Status, &issuedTo, &item.LastInspectedAt, &item.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("ppe item %s not found", id)
		}
		return nil, err
	}
	if issuedTo != "" {
		item.IssuedTo = issuedTo
	}
	return &item, nil
}

func (s *Store) SavePPEIssue(ctx context.Context, issue *ppeissue.Record) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO ppe_issues (id, item_id, issued_to_type, issued_to_id, issued_by, issued_at, expected_return)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, issue.ID, issue.ItemID, issue.IssuedToType, issue.IssuedToID, issue.IssuedBy, issue.IssuedAt, issue.ExpectedReturn)
	return err
}

func (s *Store) SavePPEInspection(ctx context.Context, rec *ppeinspection.Record) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO ppe_inspections (id, item_id, inspected_by, inspected_at, result, findings)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, rec.ID, rec.ItemID, rec.InspectedBy, rec.InspectedAt, rec.Result, rec.Findings)
	return err
}

func (s *Store) SavePPEMaintenance(ctx context.Context, rec *ppemaintenance.Record) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO ppe_maintenance (id, item_id, maintenance_by, maintenance_at, cost, actions_taken, completed_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, rec.ID, rec.ItemID, rec.MaintenanceBy, rec.MaintenanceAt, rec.Cost, rec.ActionsTaken, rec.CompletedAt)
	return err
}

func (s *Store) SearchPPEs(ctx context.Context, criteria *search.Criteria) ([]*ppe.PPE, int64, error) {
	ppeList, err := s.ListPPEs(ctx, criteria.PlantID)
	if err != nil {
		return nil, 0, err
	}
	return ppeList, int64(len(ppeList)), nil
}

func (s *Store) GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	return map[string]float64{
		"active_ppe_items":             1450.0,
		"issued_ppe_items":             940.0,
		"compliance_rate_pct":          99.4,
		"low_stock_alerts_count":       2.0,
		"inspections_due_count":        12.0,
	}, nil
}
