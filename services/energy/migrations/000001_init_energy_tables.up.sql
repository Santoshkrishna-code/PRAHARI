CREATE TABLE IF NOT EXISTS energy_profiles (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    department_id VARCHAR(50),
    facility_name VARCHAR(200) NOT NULL,
    target_score NUMERIC(10, 4) NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS energy_sources (
    id VARCHAR(50) PRIMARY KEY,
    source_name VARCHAR(100) UNIQUE NOT NULL,
    energy_type VARCHAR(50) NOT NULL,
    grid_region VARCHAR(100),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS utility_meters (
    id VARCHAR(50) PRIMARY KEY,
    meter_number VARCHAR(100) UNIQUE NOT NULL,
    source_id VARCHAR(50) REFERENCES energy_sources(id),
    asset_id VARCHAR(50),
    location_code VARCHAR(100),
    status VARCHAR(50) NOT NULL,
    last_calibrated TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS meter_readings (
    id VARCHAR(50) PRIMARY KEY,
    meter_id VARCHAR(50) REFERENCES utility_meters(id),
    reading_value NUMERIC(15, 4) NOT NULL,
    unit_of_measure VARCHAR(20) NOT NULL,
    multiplier NUMERIC(10, 4) NOT NULL DEFAULT 1.0,
    active_power_kw NUMERIC(15, 4) NOT NULL DEFAULT 0,
    reactive_power_var NUMERIC(15, 4) NOT NULL DEFAULT 0,
    power_factor NUMERIC(5, 4) NOT NULL DEFAULT 0,
    reading_time TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS energy_consumption (
    id VARCHAR(50) PRIMARY KEY,
    meter_id VARCHAR(50) REFERENCES utility_meters(id),
    period_start TIMESTAMP WITH TIME ZONE NOT NULL,
    period_end TIMESTAMP WITH TIME ZONE NOT NULL,
    consumption_kwh NUMERIC(15, 4) NOT NULL,
    peak_demand_kw NUMERIC(15, 4) NOT NULL DEFAULT 0,
    carbon_emitted_kg NUMERIC(15, 4) NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS energy_baselines (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    baseline_year INT NOT NULL,
    total_kwh NUMERIC(15, 4) NOT NULL,
    intensity_score NUMERIC(10, 4) NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS energy_targets (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    title VARCHAR(200) NOT NULL,
    target_value NUMERIC(15, 4) NOT NULL,
    current_value NUMERIC(15, 4) NOT NULL DEFAULT 0,
    unit_of_measure VARCHAR(50) NOT NULL,
    deadline TIMESTAMP WITH TIME ZONE NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS energy_forecasts (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    forecast_period VARCHAR(50) NOT NULL,
    predicted_kwh NUMERIC(15, 4) NOT NULL,
    confidence_rate NUMERIC(5, 2) NOT NULL DEFAULT 0,
    generated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS load_profiles (
    id VARCHAR(50) PRIMARY KEY,
    meter_id VARCHAR(50) REFERENCES utility_meters(id),
    demand_kw NUMERIC(15, 4) NOT NULL,
    time_interval TIMESTAMP WITH TIME ZONE NOT NULL,
    is_peak_period BOOLEAN DEFAULT false
);

CREATE TABLE IF NOT EXISTS demand_response (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    target_shed_kw NUMERIC(15, 4) NOT NULL,
    actual_shed_kw NUMERIC(15, 4) NOT NULL DEFAULT 0,
    start_time TIMESTAMP WITH TIME ZONE NOT NULL,
    end_time TIMESTAMP WITH TIME ZONE NOT NULL,
    is_successful BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tariffs (
    id VARCHAR(50) PRIMARY KEY,
    provider_name VARCHAR(200) NOT NULL,
    tariff_name VARCHAR(200) NOT NULL,
    peak_rate_kwh NUMERIC(10, 6) NOT NULL,
    off_peak_rate_kwh NUMERIC(10, 6) NOT NULL,
    demand_charge_kw NUMERIC(10, 6) NOT NULL DEFAULT 0,
    effective_from TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS renewable_generation (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    source_type VARCHAR(50) NOT NULL,
    kwh_generated NUMERIC(15, 4) NOT NULL,
    co2_offset_kg NUMERIC(15, 4) NOT NULL DEFAULT 0,
    recorded_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS energy_costs (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    billing_period VARCHAR(50) NOT NULL,
    total_cost_usd NUMERIC(15, 2) NOT NULL,
    consumption_kwh NUMERIC(15, 4) NOT NULL,
    average_rate_kwh NUMERIC(10, 6) NOT NULL DEFAULT 0,
    recorded_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS efficiency_programs (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    title VARCHAR(200) NOT NULL,
    baseline_kwh NUMERIC(15, 4) NOT NULL,
    target_saved_kwh NUMERIC(15, 4) NOT NULL,
    actual_saved_kwh NUMERIC(15, 4) NOT NULL DEFAULT 0,
    status VARCHAR(50) NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS optimization_recommendations (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    asset_id VARCHAR(50),
    title VARCHAR(200) NOT NULL,
    description TEXT,
    est_saving_usd NUMERIC(15, 2) NOT NULL DEFAULT 0,
    est_saving_kwh NUMERIC(15, 4) NOT NULL DEFAULT 0,
    priority VARCHAR(20) NOT NULL DEFAULT 'MEDIUM',
    status VARCHAR(50) NOT NULL DEFAULT 'RECOMMENDED',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS alerts (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    meter_id VARCHAR(50),
    alert_level VARCHAR(20) NOT NULL,
    trigger_message TEXT NOT NULL,
    is_resolved BOOLEAN DEFAULT false,
    triggered_at TIMESTAMP WITH TIME ZONE NOT NULL,
    resolved_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS comments (
    id VARCHAR(50) PRIMARY KEY,
    target_type VARCHAR(50) NOT NULL,
    target_id VARCHAR(50) NOT NULL,
    author_id VARCHAR(50) NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS attachments (
    id VARCHAR(50) PRIMARY KEY,
    target_type VARCHAR(50) NOT NULL,
    target_id VARCHAR(50) NOT NULL,
    file_name VARCHAR(200) NOT NULL,
    file_url TEXT NOT NULL,
    uploaded_by VARCHAR(50) NOT NULL,
    uploaded_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS timeline (
    id VARCHAR(50) PRIMARY KEY,
    record_id VARCHAR(50) NOT NULL,
    event_type VARCHAR(100) NOT NULL,
    event_date TIMESTAMP WITH TIME ZONE NOT NULL,
    actor_id VARCHAR(50) NOT NULL,
    description TEXT,
    metadata TEXT
);

CREATE TABLE IF NOT EXISTS audit_trail (
    id VARCHAR(50) PRIMARY KEY,
    action VARCHAR(50) NOT NULL,
    resource VARCHAR(100) NOT NULL,
    resource_id VARCHAR(50) NOT NULL,
    actor_id VARCHAR(50) NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    old_state TEXT,
    new_state TEXT
);

CREATE TABLE IF NOT EXISTS metrics (
    metric_key VARCHAR(100) PRIMARY KEY,
    metric_value NUMERIC(15, 4) NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
