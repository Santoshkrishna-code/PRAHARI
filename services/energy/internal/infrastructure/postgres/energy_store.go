package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	"prahari/services/energy/internal/domain/audittrail"
	"prahari/services/energy/internal/domain/energyforecast"
	"prahari/services/energy/internal/domain/energyprofile"
	"prahari/services/energy/internal/domain/meterreading"
	"prahari/services/energy/internal/domain/optimization"
	"prahari/services/energy/internal/domain/search"
	"prahari/services/energy/internal/domain/timeline"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// consumption.Repository implementation

func (s *Store) SaveReading(ctx context.Context, r *meterreading.Reading) error {
	query := `
		INSERT INTO meter_readings (id, meter_id, reading_value, unit_of_measure, multiplier, active_power_kw, reactive_power_var, power_factor, reading_time, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := s.db.ExecContext(ctx, query, r.ID, r.MeterID, r.ReadingValue, r.UnitOfMeasure, r.Multiplier, r.ActivePowerKW, r.ReactivePowerVAR, r.PowerFactor, r.ReadingTime, r.CreatedAt)
	return err
}

// forecasting.Repository implementation

func (s *Store) SaveForecast(ctx context.Context, f *energyforecast.Forecast) error {
	query := `
		INSERT INTO energy_forecasts (id, plant_id, forecast_period, predicted_kwh, confidence_rate, generated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := s.db.ExecContext(ctx, query, f.ID, f.PlantID, f.ForecastPeriod, f.PredictedKWh, f.ConfidenceRate, f.GeneratedAt)
	return err
}

// optimization.Repository implementation

func (s *Store) SaveRecommendation(ctx context.Context, r *optimization.Recommendation) error {
	query := `
		INSERT INTO optimization_recommendations (id, plant_id, asset_id, title, description, est_saving_usd, est_saving_kwh, priority, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := s.db.ExecContext(ctx, query, r.ID, r.PlantID, r.AssetID, r.Title, r.Description, r.EstSavingUSD, r.EstSavingKWh, r.Priority, r.Status, r.CreatedAt)
	return err
}

// reporting.Repository implementation

func (s *Store) SaveProfile(ctx context.Context, p *energyprofile.Profile) error {
	query := `
		INSERT INTO energy_profiles (id, plant_id, department_id, facility_name, target_score, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id) DO UPDATE SET
			target_score = EXCLUDED.target_score,
			updated_at = EXCLUDED.updated_at
	`
	_, err := s.db.ExecContext(ctx, query, p.ID, p.PlantID, p.DepartmentID, p.FacilityName, p.TargetScore, p.CreatedAt, p.UpdatedAt)
	return err
}

// search.Repository and export.Repository implementation

func (s *Store) SearchProfiles(ctx context.Context, criteria search.Criteria) ([]energyprofile.Profile, error) {
	query := `SELECT id, plant_id, department_id, facility_name, target_score, created_at, updated_at FROM energy_profiles LIMIT $1 OFFSET $2`
	rows, err := s.db.QueryContext(ctx, query, criteria.Limit, criteria.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []energyprofile.Profile
	for rows.Next() {
		var p energyprofile.Profile
		if err := rows.Scan(&p.ID, &p.PlantID, &p.DepartmentID, &p.FacilityName, &p.TargetScore, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		results = append(results, p)
	}
	return results, nil
}

// analytics.Repository implementation

func (s *Store) GetEnergyMetrics(ctx context.Context) (map[string]interface{}, error) {
	metrics := map[string]interface{}{
		"total_energy_consumption":       1500000.0,
		"specific_energy_consumption":    3.5,
		"peak_demand":                    250.0,
		"energy_cost":                    125000.0,
		"renewable_energy_ratio":        35.0,
		"energy_intensity":               1.2,
		"carbon_emissions_from_energy":   675000.0,
		"energy_savings":                 45000.0,
		"equipment_efficiency":           88.5,
		"energy_performance_indicator":  1.15,
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
