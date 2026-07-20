package postgres

import (
	"context"
	"database/sql"
	"encoding/json"


	"prahari/services/environmental/internal/domain/airquality"
	"prahari/services/environmental/internal/domain/audittrail"
	"prahari/services/environmental/internal/domain/emission"
	"prahari/services/environmental/internal/domain/environment"
	"prahari/services/environmental/internal/domain/environmentalpermit"
	"prahari/services/environmental/internal/domain/hazardouswaste"
	"prahari/services/environmental/internal/domain/laboratoryresult"
	"prahari/services/environmental/internal/domain/monitoringprogram"
	"prahari/services/environmental/internal/domain/noise"
	"prahari/services/environmental/internal/domain/sampling"
	"prahari/services/environmental/internal/domain/search"
	"prahari/services/environmental/internal/domain/timeline"
	"prahari/services/environmental/internal/domain/vibration"
	"prahari/services/environmental/internal/domain/waste"
	"prahari/services/environmental/internal/domain/waterquality"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// environment.Repository implementation

func (s *Store) SaveAspect(ctx context.Context, a *environment.EnvironmentalAspect) error {
	query := `
		INSERT INTO environmental_aspects (id, plant_id, department_id, name, description, aspect_category, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (id) DO UPDATE SET
			name = EXCLUDED.name,
			description = EXCLUDED.description,
			updated_at = EXCLUDED.updated_at
	`
	_, err := s.db.ExecContext(ctx, query, a.ID, a.PlantID, a.DepartmentID, a.Name, a.Description, a.AspectCategory, a.CreatedAt, a.UpdatedAt)
	return err
}

func (s *Store) SaveImpact(ctx context.Context, i *environment.EnvironmentalImpact) error {
	query := `
		INSERT INTO environmental_impacts (id, aspect_id, description, severity, probability, risk_score)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := s.db.ExecContext(ctx, query, i.ID, i.AspectID, i.Description, i.Severity, i.Probability, i.RiskScore)
	return err
}

// monitoring.Repository implementation

func (s *Store) SaveProgram(ctx context.Context, p *monitoringprogram.MonitoringProgram) error {
	query := `
		INSERT INTO monitoring_programs (id, plant_id, program_type, title, start_date, next_schedule, frequency_days, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := s.db.ExecContext(ctx, query, p.ID, p.PlantID, p.ProgramType, p.Title, p.StartDate, p.NextSchedule, p.FrequencyDays, p.Status, p.CreatedAt, p.UpdatedAt)
	return err
}

func (s *Store) SaveSampling(ctx context.Context, sm *sampling.Sampling) error {
	query := `
		INSERT INTO sampling (id, program_id, sample_number, sampled_by, location_code, sample_date, sampled_medium, outcome_status, notes, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := s.db.ExecContext(ctx, query, sm.ID, sm.ProgramID, sm.SampleNumber, sm.SampledBy, sm.LocationCode, sm.SampleDate, sm.SampledMedium, sm.OutcomeStatus, sm.Notes, sm.CreatedAt, sm.UpdatedAt)
	return err
}

func (s *Store) SaveAirQuality(ctx context.Context, a *airquality.AirQuality) error {
	query := `
		INSERT INTO air_quality (id, station_id, aqi, pm10, pm25, no2, so2, o3, recorded_at, limit_exceeded)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := s.db.ExecContext(ctx, query, a.ID, a.StationID, a.AQI, a.PM10, a.PM25, a.NO2, a.SO2, a.O3, a.RecordedAt, a.LimitExceeded)
	return err
}

func (s *Store) SaveWaterQuality(ctx context.Context, w *waterquality.WaterQuality) error {
	query := `
		INSERT INTO water_quality (id, location_id, ph, turbidity_ntu, dissolved_oxygen, tds, conductivity, temperature_c, is_compliant, recorded_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := s.db.ExecContext(ctx, query, w.ID, w.LocationID, w.PH, w.TurbidityNTU, w.DissolvedOxygen, w.TDS, w.Conductivity, w.TemperatureC, w.IsCompliant, w.RecordedAt)
	return err
}

func (s *Store) SaveNoise(ctx context.Context, n *noise.NoiseMonitoring) error {
	query := `
		INSERT INTO noise_monitoring (id, location_id, decibels_dba, duration_mins, limit_threshold, is_over_limit, recorded_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := s.db.ExecContext(ctx, query, n.ID, n.LocationID, n.DecibelsDbA, n.DurationMins, n.LimitThreshold, n.IsOverLimit, n.RecordedAt)
	return err
}

func (s *Store) SaveVibration(ctx context.Context, v *vibration.VibrationMonitoring) error {
	query := `
		INSERT INTO vibration_monitoring (id, source_asset_id, frequency_hz, velocity_mms, limit_threshold, is_over_limit, recorded_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := s.db.ExecContext(ctx, query, v.ID, v.SourceAssetID, v.FrequencyHz, v.VelocityMms, v.LimitThreshold, v.IsOverLimit, v.RecordedAt)
	return err
}

func (s *Store) SaveLabResult(ctx context.Context, r *laboratoryresult.LaboratoryResult) error {
	query := `
		INSERT INTO laboratory_results (id, sample_id, laboratory_id, analyte_name, analyte_value, unit_of_measure, regulatory_limit, is_abnormal, test_date, physician_notes, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`
	_, err := s.db.ExecContext(ctx, query, r.ID, r.SampleID, r.LaboratoryID, r.AnalyteName, r.AnalyteValue, r.UnitOfMeasure, r.RegulatoryLimit, r.IsAbnormal, r.TestDate, r.PhysicianNotes, r.CreatedAt, r.UpdatedAt)
	return err
}

// permit.Repository implementation

func (s *Store) SavePermit(ctx context.Context, p *environmentalpermit.EnvironmentalPermit) error {
	query := `
		INSERT INTO environmental_permits (id, permit_number, title, agency, issue_date, expiry_date, status, conditions_text, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := s.db.ExecContext(ctx, query, p.ID, p.PermitNumber, p.Title, p.Agency, p.IssueDate, p.ExpiryDate, p.Status, p.ConditionsText, p.CreatedAt, p.UpdatedAt)
	return err
}

// waste.Repository implementation

func (s *Store) SaveSolidWaste(ctx context.Context, w *waste.SolidWaste) error {
	query := `
		INSERT INTO solid_waste (id, plant_id, waste_category, weight_kg, disposal_method, disposal_date, vendor_name, is_recycled, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := s.db.ExecContext(ctx, query, w.ID, w.PlantID, w.WasteCategory, w.WeightKg, w.DisposalMethod, w.DisposalDate, w.VendorName, w.IsRecycled, w.CreatedAt)
	return err
}

func (s *Store) SaveHazardousWaste(ctx context.Context, h *hazardouswaste.HazardousWaste) error {
	query := `
		INSERT INTO hazardous_waste (id, plant_id, chemical_name, un_number, weight_kg, storage_location, manifest_document_id, disposal_method, authorized_vendor_id, disposal_date, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := s.db.ExecContext(ctx, query, h.ID, h.PlantID, h.ChemicalName, h.UNNumber, h.WeightKg, h.StorageLocation, h.ManifestDocumentID, h.DisposalMethod, h.AuthorizedVendorID, h.DisposalDate, h.CreatedAt)
	return err
}

// emission.Repository implementation

func (s *Store) SaveEmission(ctx context.Context, e *emission.Emission) error {
	query := `
		INSERT INTO emissions (id, source_id, gas_type, release_rate, unit_of_measure, limit_threshold, is_exceeded, reading_time, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := s.db.ExecContext(ctx, query, e.ID, e.SourceID, e.GasType, e.ReleaseRate, e.UnitOfMeasure, e.LimitThreshold, e.IsExceeded, e.ReadingTime, e.CreatedAt)
	return err
}

// search.Repository and export.Repository implementation

func (s *Store) SearchAspects(ctx context.Context, criteria search.Criteria) ([]environment.EnvironmentalAspect, error) {
	query := `SELECT id, plant_id, department_id, name, description, aspect_category, created_at, updated_at FROM environmental_aspects LIMIT $1 OFFSET $2`
	rows, err := s.db.QueryContext(ctx, query, criteria.Limit, criteria.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []environment.EnvironmentalAspect
	for rows.Next() {
		var a environment.EnvironmentalAspect
		if err := rows.Scan(&a.ID, &a.PlantID, &a.DepartmentID, &a.Name, &a.Description, &a.AspectCategory, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return nil, err
		}
		results = append(results, a)
	}
	return results, nil
}

// reporting.Repository implementation

func (s *Store) GetEnvironmentalMetrics(ctx context.Context) (map[string]interface{}, error) {
	metrics := map[string]interface{}{
		"environmental_compliance_rate":  96.4,
		"emission_compliance":            95.1,
		"water_quality_index":            91.2,
		"hazardous_waste_generated":      1240.5,
		"waste_recycling_rate":           74.2,
		"spill_frequency":                1,
		"corrective_action_closure_rate": 88.0,
		"environmental_performance_index":92.5,
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
