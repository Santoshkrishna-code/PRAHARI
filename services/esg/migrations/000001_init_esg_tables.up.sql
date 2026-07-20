CREATE TABLE IF NOT EXISTS sustainability_objectives (
    id VARCHAR(50) PRIMARY KEY,
    business_unit_id VARCHAR(50) NOT NULL,
    title VARCHAR(200) NOT NULL,
    category VARCHAR(50) NOT NULL,
    target_value NUMERIC(15, 4) NOT NULL,
    current_value NUMERIC(15, 4) NOT NULL,
    unit_of_measure VARCHAR(20) NOT NULL,
    deadline TIMESTAMP WITH TIME ZONE NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sustainability_initiatives (
    id VARCHAR(50) PRIMARY KEY,
    objective_id VARCHAR(50) REFERENCES sustainability_objectives(id),
    title VARCHAR(200) NOT NULL,
    description TEXT,
    budget_usd NUMERIC(15, 2) NOT NULL,
    actual_cost_usd NUMERIC(15, 2) NOT NULL,
    co2_saved_kg NUMERIC(15, 2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS carbon_inventory (
    id VARCHAR(50) PRIMARY KEY,
    business_unit_id VARCHAR(50) NOT NULL,
    period_start TIMESTAMP WITH TIME ZONE NOT NULL,
    period_end TIMESTAMP WITH TIME ZONE NOT NULL,
    scope_1_co2_kg NUMERIC(15, 2) NOT NULL,
    scope_2_co2_kg NUMERIC(15, 2) NOT NULL,
    scope_3_co2_kg NUMERIC(15, 2) NOT NULL,
    total_co2_kg NUMERIC(15, 2) NOT NULL,
    is_calculated BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS frameworks (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS disclosures (
    id VARCHAR(50) PRIMARY KEY,
    framework_id VARCHAR(50) REFERENCES frameworks(id),
    reference_code VARCHAR(100) NOT NULL,
    disclosure_text TEXT NOT NULL,
    status VARCHAR(50) NOT NULL,
    approved_by_id VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS ghg_inventory (
    id VARCHAR(50) PRIMARY KEY,
    inventory_id VARCHAR(50) REFERENCES carbon_inventory(id),
    gas_type VARCHAR(50) NOT NULL,
    volume_tonnes NUMERIC(15, 4) NOT NULL,
    gwp_factor NUMERIC(15, 4) NOT NULL,
    co2_eq_tonnes NUMERIC(15, 4) NOT NULL,
    recorded_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS emission_factors (
    id VARCHAR(50) PRIMARY KEY,
    source_name VARCHAR(200) NOT NULL,
    factor_value NUMERIC(15, 4) NOT NULL,
    unit_of_measure VARCHAR(50) NOT NULL,
    standard_name VARCHAR(100) NOT NULL,
    year INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS carbon_footprints (
    id VARCHAR(50) PRIMARY KEY,
    business_unit_id VARCHAR(50) NOT NULL,
    reporting_year INT NOT NULL,
    total_emissions NUMERIC(15, 4) NOT NULL,
    carbon_intensity NUMERIC(15, 4) NOT NULL,
    calculated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS renewable_energy (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    generation_type VARCHAR(100) NOT NULL,
    amount_kwh NUMERIC(15, 2) NOT NULL,
    carbon_offset_kg NUMERIC(15, 2) NOT NULL,
    recorded_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS energy_efficiency (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    baseline_kwh NUMERIC(15, 2) NOT NULL,
    actual_kwh NUMERIC(15, 2) NOT NULL,
    efficiency_rate NUMERIC(10, 4) NOT NULL,
    recorded_period VARCHAR(50) NOT NULL,
    recorded_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS water_stewardship (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    withdrawal_m3 NUMERIC(15, 2) NOT NULL,
    consumption_m3 NUMERIC(15, 2) NOT NULL,
    recycled_m3 NUMERIC(15, 2) NOT NULL,
    recycle_ratio NUMERIC(10, 4) NOT NULL,
    recorded_period VARCHAR(50) NOT NULL,
    recorded_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS waste_performance (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    hazardous_kg NUMERIC(15, 2) NOT NULL,
    non_hazardous_kg NUMERIC(15, 2) NOT NULL,
    total_waste_kg NUMERIC(15, 2) NOT NULL,
    recorded_period VARCHAR(50) NOT NULL,
    recorded_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS climate_indicators (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    risk_category VARCHAR(100) NOT NULL,
    impact_score INT NOT NULL,
    probability INT NOT NULL,
    financial_risk NUMERIC(15, 2) NOT NULL,
    recorded_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS esg_scorecards (
    id VARCHAR(50) PRIMARY KEY,
    business_unit_id VARCHAR(50) NOT NULL,
    period VARCHAR(50) NOT NULL,
    env_score NUMERIC(5, 2) NOT NULL,
    social_score NUMERIC(5, 2) NOT NULL,
    gov_score NUMERIC(5, 2) NOT NULL,
    overall_score NUMERIC(5, 2) NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sustainability_reports (
    id VARCHAR(50) PRIMARY KEY,
    business_unit_id VARCHAR(50) NOT NULL,
    title VARCHAR(200) NOT NULL,
    reporting_year INT NOT NULL,
    frameworks_used VARCHAR(200) NOT NULL,
    status VARCHAR(50) NOT NULL,
    approved_by VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS kpis (
    id VARCHAR(50) PRIMARY KEY,
    code VARCHAR(100) UNIQUE NOT NULL,
    title VARCHAR(200) NOT NULL,
    target_value NUMERIC(15, 4) NOT NULL,
    current_value NUMERIC(15, 4) NOT NULL,
    unit_of_measure VARCHAR(50) NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
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
    metric_value NUMERIC(10, 4) NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
