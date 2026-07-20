CREATE TABLE IF NOT EXISTS water_profiles (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    department_id VARCHAR(50),
    facility_name VARCHAR(200) NOT NULL,
    water_basin_region VARCHAR(200) NOT NULL,
    annual_budget_kl NUMERIC(15, 4) NOT NULL DEFAULT 0,
    target_recycle_pct NUMERIC(5, 2) NOT NULL DEFAULT 0,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS water_sources (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    source_name VARCHAR(100) NOT NULL,
    source_type VARCHAR(50) NOT NULL,
    location_code VARCHAR(100),
    capacity_kld NUMERIC(15, 4) NOT NULL DEFAULT 0,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS reservoirs (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    reservoir_name VARCHAR(200) NOT NULL,
    max_capacity_kl NUMERIC(15, 4) NOT NULL,
    current_level_kl NUMERIC(15, 4) NOT NULL DEFAULT 0,
    min_operating_kl NUMERIC(15, 4) NOT NULL DEFAULT 0,
    location_code VARCHAR(100),
    last_inspected_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS water_tanks (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    tank_name VARCHAR(200) NOT NULL,
    tank_type VARCHAR(50) NOT NULL,
    max_capacity_kl NUMERIC(15, 4) NOT NULL,
    current_level_kl NUMERIC(15, 4) NOT NULL DEFAULT 0,
    asset_id VARCHAR(50),
    location_code VARCHAR(100),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS pipelines (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    pipeline_name VARCHAR(200) NOT NULL,
    diameter_mm NUMERIC(10, 2) NOT NULL,
    length_meters NUMERIC(10, 2) NOT NULL,
    material VARCHAR(100),
    pressure_bar_max NUMERIC(10, 2) NOT NULL DEFAULT 0,
    from_node VARCHAR(100),
    to_node VARCHAR(100),
    asset_id VARCHAR(50),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS distribution_networks (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    zone_name VARCHAR(200) NOT NULL,
    zone_code VARCHAR(50) NOT NULL,
    supply_source_id VARCHAR(50),
    design_flow_kld NUMERIC(15, 4) NOT NULL DEFAULT 0,
    actual_flow_kld NUMERIC(15, 4) NOT NULL DEFAULT 0,
    loss_percent NUMERIC(5, 2) NOT NULL DEFAULT 0,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS flow_meters (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    meter_number VARCHAR(100) UNIQUE NOT NULL,
    meter_type VARCHAR(50) NOT NULL,
    source_id VARCHAR(50),
    pipeline_id VARCHAR(50),
    asset_id VARCHAR(50),
    location_code VARCHAR(100),
    unit_of_measure VARCHAR(20) NOT NULL DEFAULT 'KL',
    status VARCHAR(50) NOT NULL,
    last_calibrated_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS meter_readings (
    id VARCHAR(50) PRIMARY KEY,
    meter_id VARCHAR(50) REFERENCES flow_meters(id),
    reading_value_kl NUMERIC(15, 4) NOT NULL,
    flow_rate_klh NUMERIC(15, 4) NOT NULL DEFAULT 0,
    pressure_bar NUMERIC(10, 2) NOT NULL DEFAULT 0,
    temperature_c NUMERIC(5, 2) NOT NULL DEFAULT 0,
    reading_time TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS water_consumption (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    meter_id VARCHAR(50) REFERENCES flow_meters(id),
    period_start TIMESTAMP WITH TIME ZONE NOT NULL,
    period_end TIMESTAMP WITH TIME ZONE NOT NULL,
    consumption_kl NUMERIC(15, 4) NOT NULL,
    peak_flow_klh NUMERIC(15, 4) NOT NULL DEFAULT 0,
    water_intensity NUMERIC(10, 4) NOT NULL DEFAULT 0,
    production_units NUMERIC(15, 4) NOT NULL DEFAULT 0,
    cost_usd NUMERIC(15, 2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS water_balances (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    period_name VARCHAR(100) NOT NULL,
    total_intake_kl NUMERIC(15, 4) NOT NULL,
    total_consumed_kl NUMERIC(15, 4) NOT NULL,
    total_recycled_kl NUMERIC(15, 4) NOT NULL DEFAULT 0,
    water_loss_kl NUMERIC(15, 4) NOT NULL DEFAULT 0,
    loss_percent NUMERIC(5, 2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS recycling_programs (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    program_name VARCHAR(200) NOT NULL,
    source_type VARCHAR(100) NOT NULL,
    treatment_method VARCHAR(100) NOT NULL,
    input_kld NUMERIC(15, 4) NOT NULL,
    output_kld NUMERIC(15, 4) NOT NULL,
    recycle_rate_pct NUMERIC(5, 2) NOT NULL DEFAULT 0,
    status VARCHAR(50) NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS reuse_programs (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    program_name VARCHAR(200) NOT NULL,
    reuse_application VARCHAR(200) NOT NULL,
    source_stream VARCHAR(100) NOT NULL,
    volume_kld NUMERIC(15, 4) NOT NULL,
    quality_grade VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS conservation_programs (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    program_name VARCHAR(200) NOT NULL,
    baseline_kld NUMERIC(15, 4) NOT NULL,
    target_saved_kld NUMERIC(15, 4) NOT NULL,
    actual_saved_kld NUMERIC(15, 4) NOT NULL DEFAULT 0,
    investment_usd NUMERIC(15, 2) NOT NULL DEFAULT 0,
    roi_percent NUMERIC(5, 2) NOT NULL DEFAULT 0,
    status VARCHAR(50) NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS treatment_plants (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    facility_name VARCHAR(200) NOT NULL,
    treatment_type VARCHAR(100) NOT NULL,
    design_capacity_kld NUMERIC(15, 4) NOT NULL,
    operating_kld NUMERIC(15, 4) NOT NULL DEFAULT 0,
    efficiency_pct NUMERIC(5, 2) NOT NULL DEFAULT 0,
    asset_id VARCHAR(50),
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS stp (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    facility_name VARCHAR(200) NOT NULL,
    technology_type VARCHAR(100) NOT NULL,
    design_capacity_kld NUMERIC(15, 4) NOT NULL,
    influent_kld NUMERIC(15, 4) NOT NULL DEFAULT 0,
    effluent_kld NUMERIC(15, 4) NOT NULL DEFAULT 0,
    bod_influent_mg_l NUMERIC(10, 2) NOT NULL DEFAULT 0,
    bod_effluent_mg_l NUMERIC(10, 2) NOT NULL DEFAULT 0,
    cod_influent_mg_l NUMERIC(10, 2) NOT NULL DEFAULT 0,
    cod_effluent_mg_l NUMERIC(10, 2) NOT NULL DEFAULT 0,
    tss_effluent_mg_l NUMERIC(10, 2) NOT NULL DEFAULT 0,
    asset_id VARCHAR(50),
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS etp (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    facility_name VARCHAR(200) NOT NULL,
    technology_type VARCHAR(100) NOT NULL,
    design_capacity_kld NUMERIC(15, 4) NOT NULL,
    influent_kld NUMERIC(15, 4) NOT NULL DEFAULT 0,
    effluent_kld NUMERIC(15, 4) NOT NULL DEFAULT 0,
    ph_influent NUMERIC(4, 2) NOT NULL DEFAULT 7.0,
    ph_effluent NUMERIC(4, 2) NOT NULL DEFAULT 7.0,
    cod_influent_mg_l NUMERIC(10, 2) NOT NULL DEFAULT 0,
    cod_effluent_mg_l NUMERIC(10, 2) NOT NULL DEFAULT 0,
    tds_effluent_mg_l NUMERIC(10, 2) NOT NULL DEFAULT 0,
    oil_grease_effluent_mg_l NUMERIC(10, 2) NOT NULL DEFAULT 0,
    asset_id VARCHAR(50),
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS desalination_plants (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    facility_name VARCHAR(200) NOT NULL,
    technology VARCHAR(100) NOT NULL,
    design_capacity_kld NUMERIC(15, 4) NOT NULL,
    operating_kld NUMERIC(15, 4) NOT NULL DEFAULT 0,
    recovery_rate_pct NUMERIC(5, 2) NOT NULL DEFAULT 0,
    energy_per_kl_kwh NUMERIC(10, 4) NOT NULL DEFAULT 0,
    brine_discharge_kld NUMERIC(15, 4) NOT NULL DEFAULT 0,
    asset_id VARCHAR(50),
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS leak_detection (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    pipeline_id VARCHAR(50),
    zone_code VARCHAR(50) NOT NULL,
    detection_method VARCHAR(100) NOT NULL,
    estimated_loss_kld NUMERIC(15, 4) NOT NULL,
    severity VARCHAR(20) NOT NULL DEFAULT 'MEDIUM',
    location_desc TEXT,
    work_order_id VARCHAR(50),
    is_resolved BOOLEAN DEFAULT false,
    detected_at TIMESTAMP WITH TIME ZONE NOT NULL,
    resolved_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS forecasts (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    forecast_period VARCHAR(50) NOT NULL,
    predicted_kl NUMERIC(15, 4) NOT NULL,
    confidence_rate NUMERIC(5, 2) NOT NULL DEFAULT 0,
    seasonal_factor NUMERIC(5, 2) NOT NULL DEFAULT 1.0,
    generated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS optimization_recommendations (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    asset_id VARCHAR(50),
    title VARCHAR(200) NOT NULL,
    description TEXT,
    est_saving_kld NUMERIC(15, 4) NOT NULL DEFAULT 0,
    est_saving_usd NUMERIC(15, 2) NOT NULL DEFAULT 0,
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
