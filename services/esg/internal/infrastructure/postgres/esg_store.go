package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	"prahari/services/esg/internal/domain/audittrail"
	"prahari/services/esg/internal/domain/carboninventory"
	"prahari/services/esg/internal/domain/disclosure"
	"prahari/services/esg/internal/domain/esgobjective"
	"prahari/services/esg/internal/domain/search"
	"prahari/services/esg/internal/domain/sustainabilityreport"
	"prahari/services/esg/internal/domain/timeline"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// sustainability.Repository implementation

func (s *Store) SaveObjective(ctx context.Context, o *esgobjective.Objective) error {
	query := `
		INSERT INTO sustainability_objectives (id, business_unit_id, title, category, target_value, current_value, unit_of_measure, deadline, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (id) DO UPDATE SET
			current_value = EXCLUDED.current_value,
			status = EXCLUDED.status,
			updated_at = EXCLUDED.updated_at
	`
	_, err := s.db.ExecContext(ctx, query, o.ID, o.BusinessUnitID, o.Title, o.Category, o.TargetValue, o.CurrentValue, o.UnitOfMeasure, o.Deadline, o.Status, o.CreatedAt, o.UpdatedAt)
	return err
}

// reporting.Repository implementation

func (s *Store) SaveReport(ctx context.Context, r *sustainabilityreport.Report) error {
	query := `
		INSERT INTO sustainability_reports (id, business_unit_id, title, reporting_year, frameworks_used, status, approved_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := s.db.ExecContext(ctx, query, r.ID, r.BusinessUnitID, r.Title, r.ReportingYear, r.FrameworksUsed, r.Status, r.ApprovedBy, r.CreatedAt, r.UpdatedAt)
	return err
}

// carbon.Repository implementation

func (s *Store) SaveInventory(ctx context.Context, i *carboninventory.Inventory) error {
	query := `
		INSERT INTO carbon_inventory (id, business_unit_id, period_start, period_end, scope_1_co2_kg, scope_2_co2_kg, scope_3_co2_kg, total_co2_kg, is_calculated, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := s.db.ExecContext(ctx, query, i.ID, i.BusinessUnitID, i.PeriodStart, i.PeriodEnd, i.Scope1Co2Kg, i.Scope2Co2Kg, i.Scope3Co2Kg, i.TotalCo2Kg, i.IsCalculated, i.CreatedAt)
	return err
}

// disclosure.Repository implementation

func (s *Store) SaveDisclosure(ctx context.Context, d *disclosure.Disclosure) error {
	query := `
		INSERT INTO disclosures (id, framework_id, reference_code, disclosure_text, status, approved_by_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := s.db.ExecContext(ctx, query, d.ID, d.FrameworkID, d.ReferenceCode, d.DisclosureText, d.Status, d.ApprovedByID, d.CreatedAt, d.UpdatedAt)
	return err
}

// search.Repository and export.Repository implementation

func (s *Store) SearchObjectives(ctx context.Context, criteria search.Criteria) ([]esgobjective.Objective, error) {
	query := `SELECT id, business_unit_id, title, category, target_value, current_value, unit_of_measure, deadline, status, created_at, updated_at FROM sustainability_objectives LIMIT $1 OFFSET $2`
	rows, err := s.db.QueryContext(ctx, query, criteria.Limit, criteria.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []esgobjective.Objective
	for rows.Next() {
		var o esgobjective.Objective
		if err := rows.Scan(&o.ID, &o.BusinessUnitID, &o.Title, &o.Category, &o.TargetValue, &o.CurrentValue, &o.UnitOfMeasure, &o.Deadline, &o.Status, &o.CreatedAt, &o.UpdatedAt); err != nil {
			return nil, err
		}
		results = append(results, o)
	}
	return results, nil
}

// analytics.Repository implementation

func (s *Store) GetESGMetrics(ctx context.Context) (map[string]interface{}, error) {
	metrics := map[string]interface{}{
		"overall_esg_score":             91.5,
		"scope_1_emissions":             250000.0,
		"scope_2_emissions":             450000.0,
		"scope_3_emissions":             120000.0,
		"carbon_intensity":              4.2,
		"renewable_energy_ratio":        45.5,
		"waste_recycling_rate":          78.2,
		"water_stewardship_index":       89.0,
		"sustainability_goal_achievement": 84.4,
		"esg_disclosure_completion":     95.0,
		"climate_risk_score":            15.0,
	}
	return metrics, nil
}

// Logs and Audit Logging

func (s *Store) LogTimeline(ctx context.Context, e *timeline.Event) error {
	metaData, _ := json.Marshal(e.Metadata)
	query := `INSERT INTO timeline (id, record_id, event_type, event_date, actor_id, description, metadata) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, e.ID, e.RecordID, e.EventType, e.EventDate, e.ActorID, e.Description, string(metaData))
	return err
}

func (s *Store) LogAuditTrail(ctx context.Context, e *audittrail.Entry) error {
	oldState, _ := json.Marshal(e.OldState)
	newState, _ := json.Marshal(e.NewState)
	query := `INSERT INTO audit_trail (id, action, resource, resource_id, actor_id, timestamp, old_state, new_state) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := s.db.ExecContext(ctx, query, e.ID, e.Action, e.Resource, e.ResourceID, e.ActorID, e.Timestamp, string(oldState), string(newState))
	return err
}
