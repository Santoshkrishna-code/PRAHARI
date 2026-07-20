package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"prahari/services/water/internal/domain/consumption"
	"prahari/services/water/internal/domain/distributionnetwork"
	"prahari/services/water/internal/domain/flowmeter"
	"prahari/services/water/internal/domain/forecasting"
	"prahari/services/water/internal/domain/leakdetection"
	"prahari/services/water/internal/domain/meterreading"
	"prahari/services/water/internal/domain/optimization"
	"prahari/services/water/internal/domain/pipeline"
	"prahari/services/water/internal/domain/recycling"
	"prahari/services/water/internal/domain/reuse"
	"prahari/services/water/internal/domain/search"
	"prahari/services/water/internal/domain/waterprofile"
	"prahari/services/water/internal/domain/watersource"
)


type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) SaveProfile(ctx context.Context, p *waterprofile.Profile) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO water_profiles (id, plant_id, department_id, facility_name, water_basin_region, annual_budget_kl, target_recycle_pct, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		ON CONFLICT (id) DO UPDATE SET facility_name = EXCLUDED.facility_name, updated_at = EXCLUDED.updated_at`
	_, err := s.db.ExecContext(ctx, query, p.ID, p.PlantID, p.DepartmentID, p.FacilityName, p.WaterBasinRegion, p.AnnualBudgetKL, p.TargetRecyclePct, p.Status, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to save water profile: %w", err)
	}
	return nil
}

func (s *Store) GetProfileByID(ctx context.Context, id string) (*waterprofile.Profile, error) {
	if s.db == nil {
		return &waterprofile.Profile{ID: id, FacilityName: "Mock Facility", WaterBasinRegion: "Basin A", AnnualBudgetKL: 50000.0, TargetRecyclePct: 40.0, Status: "ACTIVE"}, nil
	}
	query := `SELECT id, plant_id, COALESCE(department_id, ''), facility_name, water_basin_region, annual_budget_kl, target_recycle_pct, status, created_at, updated_at FROM water_profiles WHERE id = $1`
	row := s.db.QueryRowContext(ctx, query, id)
	var p waterprofile.Profile
	if err := row.Scan(&p.ID, &p.PlantID, &p.DepartmentID, &p.FacilityName, &p.WaterBasinRegion, &p.AnnualBudgetKL, &p.TargetRecyclePct, &p.Status, &p.CreatedAt, &p.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("water profile %s not found", id)
		}
		return nil, err
	}

	return &p, nil
}

func (s *Store) ListProfiles(ctx context.Context, plantID string) ([]*waterprofile.Profile, error) {
	if s.db == nil {
		return []*waterprofile.Profile{
			{ID: "wp-001", PlantID: plantID, FacilityName: "Plant Main Water Network", WaterBasinRegion: "Northern Basin", AnnualBudgetKL: 120000.0, TargetRecyclePct: 50.0, Status: "ACTIVE"},
		}, nil
	}
	query := `SELECT id, plant_id, COALESCE(department_id, ''), facility_name, water_basin_region, annual_budget_kl, target_recycle_pct, status, created_at, updated_at FROM water_profiles WHERE plant_id = $1`
	rows, err := s.db.QueryContext(ctx, query, plantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*waterprofile.Profile
	for rows.Next() {
		var p waterprofile.Profile
		if err := rows.Scan(&p.ID, &p.PlantID, &p.DepartmentID, &p.FacilityName, &p.WaterBasinRegion, &p.AnnualBudgetKL, &p.TargetRecyclePct, &p.Status, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, &p)
	}
	return result, nil
}

func (s *Store) SaveSource(ctx context.Context, src *watersource.Source) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO water_sources (id, plant_id, source_name, source_type, location_code, capacity_kld, is_active, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := s.db.ExecContext(ctx, query, src.ID, src.PlantID, src.SourceName, src.SourceType, src.LocationCode, src.CapacityKLD, src.IsActive, src.CreatedAt)
	return err
}

func (s *Store) SaveReading(ctx context.Context, r *meterreading.Reading) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO meter_readings (id, meter_id, reading_value_kl, flow_rate_klh, pressure_bar, temperature_c, reading_time, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := s.db.ExecContext(ctx, query, r.ID, r.MeterID, r.ReadingValueKL, r.FlowRateKLH, r.PressureBar, r.TemperatureC, r.ReadingTime, r.CreatedAt)
	return err
}

func (s *Store) SaveConsumption(ctx context.Context, c *consumption.Consumption) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO water_consumption (id, plant_id, meter_id, period_start, period_end, consumption_kl, peak_flow_klh, water_intensity, production_units, cost_usd, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.PlantID, c.MeterID, c.PeriodStart, c.PeriodEnd, c.ConsumptionKL, c.PeakFlowKLH, c.WaterIntensity, c.ProductionUnits, c.CostUSD, c.CreatedAt)
	return err
}

func (s *Store) GetLatestReading(ctx context.Context, meterID string) (*meterreading.Reading, error) {
	if s.db == nil {
		return &meterreading.Reading{MeterID: meterID, ReadingValueKL: 12500.0}, nil
	}
	query := `SELECT id, meter_id, reading_value_kl, flow_rate_klh, pressure_bar, temperature_c, reading_time, created_at FROM meter_readings WHERE meter_id = $1 ORDER BY reading_time DESC LIMIT 1`
	var r meterreading.Reading
	err := s.db.QueryRowContext(ctx, query, meterID).Scan(&r.ID, &r.MeterID, &r.ReadingValueKL, &r.FlowRateKLH, &r.PressureBar, &r.TemperatureC, &r.ReadingTime, &r.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *Store) SaveRecyclingProgram(ctx context.Context, prog *recycling.Program) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO recycling_programs (id, plant_id, program_name, source_type, treatment_method, input_kld, output_kld, recycle_rate_pct, status, start_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err := s.db.ExecContext(ctx, query, prog.ID, prog.PlantID, prog.ProgramName, prog.SourceType, prog.TreatmentMethod, prog.InputKLD, prog.OutputKLD, prog.RecycleRatePct, prog.Status, prog.StartDate, prog.CreatedAt, prog.UpdatedAt)
	return err
}

func (s *Store) SaveReuseProgram(ctx context.Context, prog *reuse.Program) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO reuse_programs (id, plant_id, program_name, reuse_application, source_stream, volume_kld, quality_grade, status, start_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := s.db.ExecContext(ctx, query, prog.ID, prog.PlantID, prog.ProgramName, prog.ReuseApplication, prog.SourceStream, prog.VolumeKLD, prog.QualityGrade, prog.Status, prog.StartDate, prog.CreatedAt, prog.UpdatedAt)
	return err
}

func (s *Store) SaveNetwork(ctx context.Context, net *distributionnetwork.Network) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO distribution_networks (id, plant_id, zone_name, zone_code, supply_source_id, design_flow_kld, actual_flow_kld, loss_percent, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := s.db.ExecContext(ctx, query, net.ID, net.PlantID, net.ZoneName, net.ZoneCode, net.SupplySourceID, net.DesignFlowKLD, net.ActualFlowKLD, net.LossPercent, net.IsActive, net.CreatedAt, net.UpdatedAt)
	return err
}

func (s *Store) SavePipeline(ctx context.Context, pipe *pipeline.Pipeline) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO pipelines (id, plant_id, pipeline_name, diameter_mm, length_meters, material, pressure_bar_max, from_node, to_node, is_active, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := s.db.ExecContext(ctx, query, pipe.ID, pipe.PlantID, pipe.PipelineName, pipe.DiameterMM, pipe.LengthMeters, pipe.Material, pipe.PressureBarMax, pipe.FromNode, pipe.ToNode, pipe.IsActive, pipe.CreatedAt)
	return err
}

func (s *Store) SaveLeak(ctx context.Context, leak *leakdetection.Leak) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO leak_detection (id, plant_id, zone_code, detection_method, estimated_loss_kld, severity, location_desc, work_order_id, is_resolved, detected_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := s.db.ExecContext(ctx, query, leak.ID, leak.PlantID, leak.ZoneCode, leak.DetectionMethod, leak.EstimatedLossKLD, leak.Severity, leak.LocationDesc, leak.WorkOrderID, leak.IsResolved, leak.DetectedAt, leak.CreatedAt)
	return err
}

func (s *Store) SaveForecast(ctx context.Context, fc *forecasting.Forecast) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO forecasts (id, plant_id, forecast_period, predicted_kl, confidence_rate, seasonal_factor, generated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, fc.ID, fc.PlantID, fc.ForecastPeriod, fc.PredictedKL, fc.ConfidenceRate, fc.SeasonalFactor, fc.GeneratedAt)
	return err
}

func (s *Store) SaveRecommendation(ctx context.Context, rec *optimization.Recommendation) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO optimization_recommendations (id, plant_id, title, description, est_saving_kld, est_saving_usd, priority, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := s.db.ExecContext(ctx, query, rec.ID, rec.PlantID, rec.Title, rec.Description, rec.EstSavingKLD, rec.EstSavingUSD, rec.Priority, rec.Status, rec.CreatedAt)
	return err
}

func (s *Store) SearchProfiles(ctx context.Context, criteria *search.Criteria) ([]*waterprofile.Profile, int64, error) {
	profiles, err := s.ListProfiles(ctx, criteria.PlantID)
	if err != nil {
		return nil, 0, err
	}
	return profiles, int64(len(profiles)), nil
}

func (s *Store) GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	return map[string]float64{
		"total_consumption_kl":     48500.0,
		"recycling_ratio_pct":      42.5,
		"water_loss_pct":           4.2,
		"water_stewardship_index":  94.8,
		"treatment_efficiency_pct": 98.1,
	}, nil
}

func (s *Store) SaveFlowMeter(ctx context.Context, m *flowmeter.FlowMeter) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO flow_meters (id, plant_id, meter_number, meter_type, location_code, unit_of_measure, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := s.db.ExecContext(ctx, query, m.ID, m.PlantID, m.MeterNumber, m.MeterType, m.LocationCode, m.UnitOfMeasure, m.Status, m.CreatedAt, m.UpdatedAt)
	return err
}
