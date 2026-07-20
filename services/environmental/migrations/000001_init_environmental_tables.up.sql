CREATE TABLE IF NOT EXISTS environmental_aspects (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    department_id VARCHAR(50) NOT NULL,
    name VARCHAR(200) NOT NULL,
    description TEXT,
    aspect_category VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS environmental_impacts (
    id VARCHAR(50) PRIMARY KEY,
    aspect_id VARCHAR(50) REFERENCES environmental_aspects(id),
    description TEXT,
    severity INT NOT NULL,
    probability INT NOT NULL,
    risk_score INT NOT NULL
);

CREATE TABLE IF NOT EXISTS environmental_permits (
    id VARCHAR(50) PRIMARY KEY,
    permit_number VARCHAR(100) UNIQUE NOT NULL,
    title VARCHAR(200) NOT NULL,
    agency VARCHAR(100) NOT NULL,
    issue_date TIMESTAMP WITH TIME ZONE NOT NULL,
    expiry_date TIMESTAMP WITH TIME ZONE NOT NULL,
    status VARCHAR(50) NOT NULL,
    conditions_text TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS emissions (
    id VARCHAR(50) PRIMARY KEY,
    source_id VARCHAR(50) NOT NULL,
    gas_type VARCHAR(50) NOT NULL,
    release_rate NUMERIC(10, 4) NOT NULL,
    unit_of_measure VARCHAR(20) NOT NULL,
    limit_threshold NUMERIC(10, 4) NOT NULL,
    is_exceeded BOOLEAN DEFAULT false,
    reading_time TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS air_quality (
    id VARCHAR(50) PRIMARY KEY,
    station_id VARCHAR(50) NOT NULL,
    aqi INT NOT NULL,
    pm10 NUMERIC(10, 4) NOT NULL,
    pm25 NUMERIC(10, 4) NOT NULL,
    no2 NUMERIC(10, 4) NOT NULL,
    so2 NUMERIC(10, 4) NOT NULL,
    o3 NUMERIC(10, 4) NOT NULL,
    recorded_at TIMESTAMP WITH TIME ZONE NOT NULL,
    limit_exceeded BOOLEAN DEFAULT false
);

CREATE TABLE IF NOT EXISTS water_quality (
    id VARCHAR(50) PRIMARY KEY,
    location_id VARCHAR(50) NOT NULL,
    ph NUMERIC(4, 2) NOT NULL,
    turbidity_ntu NUMERIC(10, 4) NOT NULL,
    dissolved_oxygen NUMERIC(10, 4) NOT NULL,
    tds NUMERIC(10, 4) NOT NULL,
    conductivity NUMERIC(10, 4) NOT NULL,
    temperature_c NUMERIC(5, 2) NOT NULL,
    is_compliant BOOLEAN DEFAULT true,
    recorded_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS wastewater (
    id VARCHAR(50) PRIMARY KEY,
    outfall_id VARCHAR(50) NOT NULL,
    bod NUMERIC(10, 4) NOT NULL,
    cod NUMERIC(10, 4) NOT NULL,
    tss NUMERIC(10, 4) NOT NULL,
    oil_and_grease NUMERIC(10, 4) NOT NULL,
    flow_rate_m3 NUMERIC(10, 4) NOT NULL,
    is_compliant BOOLEAN DEFAULT true,
    recorded_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS solid_waste (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    waste_category VARCHAR(50) NOT NULL,
    weight_kg NUMERIC(10, 2) NOT NULL,
    disposal_method VARCHAR(50) NOT NULL,
    disposal_date TIMESTAMP WITH TIME ZONE NOT NULL,
    vendor_name VARCHAR(200),
    is_recycled BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS hazardous_waste (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    chemical_name VARCHAR(200) NOT NULL,
    un_number VARCHAR(20) NOT NULL,
    weight_kg NUMERIC(10, 2) NOT NULL,
    storage_location VARCHAR(200) NOT NULL,
    manifest_document_id VARCHAR(50),
    disposal_method VARCHAR(50) NOT NULL,
    authorized_vendor_id VARCHAR(50) NOT NULL,
    disposal_date TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS chemical_spills (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    incident_id VARCHAR(50),
    chemical_name VARCHAR(200) NOT NULL,
    volume_spilled NUMERIC(10, 2) NOT NULL,
    unit_of_measure VARCHAR(20) NOT NULL,
    contained BOOLEAN DEFAULT false,
    reached_water BOOLEAN DEFAULT false,
    response_action TEXT,
    spill_time TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS soil_contamination (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    agent_name VARCHAR(100) NOT NULL,
    area_sq_meters NUMERIC(10, 2) NOT NULL,
    depth_cm NUMERIC(10, 2) NOT NULL,
    remediation_status VARCHAR(50) NOT NULL,
    notes TEXT,
    detected_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS noise_monitoring (
    id VARCHAR(50) PRIMARY KEY,
    location_id VARCHAR(50) NOT NULL,
    decibels_dba NUMERIC(5, 2) NOT NULL,
    duration_mins INT NOT NULL,
    limit_threshold NUMERIC(5, 2) NOT NULL,
    is_over_limit BOOLEAN DEFAULT false,
    recorded_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS vibration_monitoring (
    id VARCHAR(50) PRIMARY KEY,
    source_asset_id VARCHAR(50) NOT NULL,
    frequency_hz NUMERIC(10, 2) NOT NULL,
    velocity_mms NUMERIC(10, 2) NOT NULL,
    limit_threshold NUMERIC(10, 2) NOT NULL,
    is_over_limit BOOLEAN DEFAULT false,
    recorded_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS monitoring_programs (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    program_type VARCHAR(100) NOT NULL,
    title VARCHAR(200) NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    next_schedule TIMESTAMP WITH TIME ZONE NOT NULL,
    frequency_days INT NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sampling (
    id VARCHAR(50) PRIMARY KEY,
    program_id VARCHAR(50) REFERENCES monitoring_programs(id),
    sample_number VARCHAR(100) UNIQUE NOT NULL,
    sampled_by VARCHAR(100) NOT NULL,
    location_code VARCHAR(100) NOT NULL,
    sample_date TIMESTAMP WITH TIME ZONE NOT NULL,
    sampled_medium VARCHAR(50) NOT NULL,
    outcome_status VARCHAR(50) NOT NULL,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS laboratories (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    license_number VARCHAR(100) UNIQUE NOT NULL,
    contact_email VARCHAR(100),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS laboratory_results (
    id VARCHAR(50) PRIMARY KEY,
    sample_id VARCHAR(50) REFERENCES sampling(id),
    laboratory_id VARCHAR(50) REFERENCES laboratories(id),
    test_name VARCHAR(100) NOT NULL,
    test_value VARCHAR(100) NOT NULL,
    reference_range VARCHAR(100),
    is_abnormal BOOLEAN DEFAULT false,
    test_date TIMESTAMP WITH TIME ZONE NOT NULL,
    physician_notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS corrective_actions (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    source_type VARCHAR(50) NOT NULL,
    source_id VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    assigned_to VARCHAR(100),
    target_date TIMESTAMP WITH TIME ZONE NOT NULL,
    actual_date TIMESTAMP WITH TIME ZONE,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
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
