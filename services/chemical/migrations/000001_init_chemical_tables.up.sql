CREATE TABLE IF NOT EXISTS chemicals (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    name VARCHAR(200) NOT NULL,
    cas_number VARCHAR(50) NOT NULL,
    iupac_name VARCHAR(200) NOT NULL DEFAULT '',
    formula VARCHAR(100) NOT NULL DEFAULT '',
    molecular_weight NUMERIC(10, 4) NOT NULL DEFAULT 0,
    physical_state VARCHAR(50) NOT NULL,
    is_restricted BOOLEAN NOT NULL DEFAULT false,
    max_allowable_qty NUMERIC(15, 4) NOT NULL DEFAULT 0,
    status VARCHAR(50) NOT NULL DEFAULT 'REQUESTED',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS manufacturers (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    address TEXT NOT NULL DEFAULT '',
    contact_info TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS suppliers (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    address TEXT NOT NULL DEFAULT '',
    contact_info TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS chemical_batches (
    id VARCHAR(50) PRIMARY KEY,
    chemical_id VARCHAR(50) REFERENCES chemicals(id),
    batch_number VARCHAR(100) NOT NULL,
    manufacture_date TIMESTAMP WITH TIME ZONE NOT NULL,
    expiry_date TIMESTAMP WITH TIME ZONE NOT NULL,
    manufacturer_id VARCHAR(50) REFERENCES manufacturers(id),
    cert_of_analysis TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS storage_areas (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    name VARCHAR(200) NOT NULL,
    code VARCHAR(50) NOT NULL,
    ventilation_type VARCHAR(100) NOT NULL,
    max_capacity_qty NUMERIC(15, 4) NOT NULL DEFAULT 0,
    current_load_qty NUMERIC(15, 4) NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS containers (
    id VARCHAR(50) PRIMARY KEY,
    chemical_id VARCHAR(50) REFERENCES chemicals(id),
    batch_id VARCHAR(50) REFERENCES chemical_batches(id),
    barcode VARCHAR(100) UNIQUE NOT NULL,
    storage_area_id VARCHAR(50) REFERENCES storage_areas(id),
    capacity NUMERIC(15, 4) NOT NULL,
    current_volume NUMERIC(15, 4) NOT NULL,
    unit_of_measure VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'RECEIVED',
    expiry_date TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS chemical_inventory (
    id VARCHAR(50) PRIMARY KEY,
    chemical_id VARCHAR(50) REFERENCES chemicals(id),
    storage_area_id VARCHAR(50) REFERENCES storage_areas(id),
    current_qty NUMERIC(15, 4) NOT NULL DEFAULT 0,
    reserved_qty NUMERIC(15, 4) NOT NULL DEFAULT 0,
    unit_of_measure VARCHAR(50) NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS chemical_transfers (
    id VARCHAR(50) PRIMARY KEY,
    container_id VARCHAR(50) REFERENCES containers(id),
    from_area_id VARCHAR(50) REFERENCES storage_areas(id),
    to_area_id VARCHAR(50) REFERENCES storage_areas(id),
    transferred_by VARCHAR(50) NOT NULL,
    transferred_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS storage_conditions (
    id VARCHAR(50) PRIMARY KEY,
    chemical_id VARCHAR(50) REFERENCES chemicals(id),
    temp_min_c NUMERIC(5, 2) NOT NULL DEFAULT 0,
    temp_max_c NUMERIC(5, 2) NOT NULL DEFAULT 40,
    humidity_max NUMERIC(5, 2) NOT NULL DEFAULT 100,
    lighting VARCHAR(100) NOT NULL DEFAULT 'GENERAL',
    ventilation VARCHAR(100) NOT NULL DEFAULT 'STANDARD'
);

CREATE TABLE IF NOT EXISTS compatibility_rules (
    id VARCHAR(50) PRIMARY KEY,
    class_a VARCHAR(100) NOT NULL,
    class_b VARCHAR(100) NOT NULL,
    compatible BOOLEAN NOT NULL DEFAULT true,
    segregation_req TEXT
);

CREATE TABLE IF NOT EXISTS incompatibility_records (
    id VARCHAR(50) PRIMARY KEY,
    chemical_id VARCHAR(50) REFERENCES chemicals(id),
    other_chem_id VARCHAR(50) REFERENCES chemicals(id),
    alert_reason TEXT NOT NULL,
    severity_level VARCHAR(50) NOT NULL DEFAULT 'WARNING'
);

CREATE TABLE IF NOT EXISTS safety_data_sheets (
    id VARCHAR(50) PRIMARY KEY,
    chemical_id VARCHAR(50) REFERENCES chemicals(id),
    version VARCHAR(50) NOT NULL,
    manufacturer VARCHAR(200) NOT NULL DEFAULT '',
    publish_date TIMESTAMP WITH TIME ZONE NOT NULL,
    document_url TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sds_revisions (
    id VARCHAR(50) PRIMARY KEY,
    sds_id VARCHAR(50) REFERENCES safety_data_sheets(id),
    revision_num VARCHAR(50) NOT NULL,
    revised_by VARCHAR(50) NOT NULL,
    revised_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    change_log TEXT NOT NULL,
    document_url TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS ghs_classifications (
    id VARCHAR(50) PRIMARY KEY,
    chemical_id VARCHAR(50) REFERENCES chemicals(id),
    signal_word VARCHAR(50) NOT NULL DEFAULT 'WARNING',
    hazard_class VARCHAR(100) NOT NULL,
    category VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS hazard_statements (
    id VARCHAR(50) PRIMARY KEY,
    chemical_id VARCHAR(50) REFERENCES chemicals(id),
    h_code VARCHAR(50) NOT NULL,
    statement TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS precautionary_statements (
    id VARCHAR(50) PRIMARY KEY,
    chemical_id VARCHAR(50) REFERENCES chemicals(id),
    p_code VARCHAR(50) NOT NULL,
    statement TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS pictograms (
    id VARCHAR(50) PRIMARY KEY,
    chemical_id VARCHAR(50) REFERENCES chemicals(id),
    code VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    image_url TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS labels (
    id VARCHAR(50) PRIMARY KEY,
    chemical_id VARCHAR(50) REFERENCES chemicals(id),
    nfpa_health INTEGER NOT NULL DEFAULT 0,
    nfpa_fire INTEGER NOT NULL DEFAULT 0,
    nfpa_noreact INTEGER NOT NULL DEFAULT 0,
    nfpa_special VARCHAR(50),
    label_format VARCHAR(50) NOT NULL DEFAULT 'GHS',
    generated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    generated_by VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS exposure_limits (
    id VARCHAR(50) PRIMARY KEY,
    chemical_id VARCHAR(50) REFERENCES chemicals(id),
    limit_type VARCHAR(50) NOT NULL,
    value_ppm NUMERIC(10, 4),
    value_mg_m3 NUMERIC(10, 4),
    source VARCHAR(100) NOT NULL,
    effective_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS exposure_monitoring (
    id VARCHAR(50) PRIMARY KEY,
    chemical_id VARCHAR(50) REFERENCES chemicals(id),
    employee_id VARCHAR(50) NOT NULL,
    exposure_ppm NUMERIC(10, 4) NOT NULL,
    measured_at TIMESTAMP WITH TIME ZONE NOT NULL,
    measured_by VARCHAR(50) NOT NULL,
    comments TEXT
);

CREATE TABLE IF NOT EXISTS spill_responses (
    id VARCHAR(50) PRIMARY KEY,
    chemical_id VARCHAR(50) REFERENCES chemicals(id),
    spill_kit_req TEXT NOT NULL,
    containment TEXT NOT NULL,
    absorbent TEXT NOT NULL,
    ppe_required TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS fire_responses (
    id VARCHAR(50) PRIMARY KEY,
    chemical_id VARCHAR(50) REFERENCES chemicals(id),
    media_suitable TEXT NOT NULL,
    media_unsuitable TEXT NOT NULL,
    hazard_products TEXT NOT NULL,
    fire_fighter_ppe TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS first_aid (
    id VARCHAR(50) PRIMARY KEY,
    chemical_id VARCHAR(50) REFERENCES chemicals(id),
    inhalation TEXT NOT NULL,
    skin TEXT NOT NULL,
    eyes TEXT NOT NULL,
    ingestion TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS emergency_contacts (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    contact_name VARCHAR(200) NOT NULL,
    phone_number VARCHAR(50) NOT NULL,
    role VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS waste_classifications (
    id VARCHAR(50) PRIMARY KEY,
    chemical_id VARCHAR(50) REFERENCES chemicals(id),
    rcra_code VARCHAR(50),
    ewc_code VARCHAR(50),
    hazard_class VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS disposals (
    id VARCHAR(50) PRIMARY KEY,
    container_id VARCHAR(50) REFERENCES containers(id),
    qty_disposed NUMERIC(15, 4) NOT NULL,
    unit_of_measure VARCHAR(50) NOT NULL,
    disposed_by VARCHAR(50) NOT NULL,
    disposed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    facility_name VARCHAR(200) NOT NULL,
    manifest_num VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS shelf_lives (
    id VARCHAR(50) PRIMARY KEY,
    chemical_id VARCHAR(50) REFERENCES chemicals(id),
    life_days INTEGER NOT NULL,
    alert_threshold_days INTEGER NOT NULL DEFAULT 30
);

CREATE TABLE IF NOT EXISTS expiry_tracking (
    id VARCHAR(50) PRIMARY KEY,
    container_id VARCHAR(50) REFERENCES containers(id),
    expiry_date TIMESTAMP WITH TIME ZONE NOT NULL,
    alert_sent BOOLEAN NOT NULL DEFAULT false,
    alert_sent_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS chemical_approvals (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    chemical_name VARCHAR(200) NOT NULL,
    cas_number VARCHAR(50) NOT NULL,
    requested_by VARCHAR(50) NOT NULL,
    technical_reviewer_id VARCHAR(50),
    technical_approved_at TIMESTAMP WITH TIME ZONE,
    safety_reviewer_id VARCHAR(50),
    safety_approved_at TIMESTAMP WITH TIME ZONE,
    env_reviewer_id VARCHAR(50),
    env_approved_at TIMESTAMP WITH TIME ZONE,
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING_REVIEW',
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
    metric_value NUMERIC(15, 4) NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
