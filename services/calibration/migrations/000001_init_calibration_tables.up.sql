CREATE TABLE IF NOT EXISTS instruments (
    id VARCHAR(50) PRIMARY KEY,
    asset_id VARCHAR(50) NOT NULL,
    plant_id VARCHAR(50) NOT NULL,
    tag_number VARCHAR(100) NOT NULL,
    model_number VARCHAR(100) NOT NULL,
    manufacturer VARCHAR(100) NOT NULL,
    instrument_type VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS calibration_plans (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    procedure_name VARCHAR(100) NOT NULL,
    interval_months INT NOT NULL DEFAULT 12,
    instructions TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS calibration_schedules (
    id VARCHAR(50) PRIMARY KEY,
    instrument_id VARCHAR(50) REFERENCES instruments(id),
    plan_id VARCHAR(50) REFERENCES calibration_plans(id),
    scheduled_for TIMESTAMP WITH TIME ZONE NOT NULL,
    completed_at TIMESTAMP WITH TIME ZONE,
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING'
);

CREATE TABLE IF NOT EXISTS calibration_records (
    id VARCHAR(50) PRIMARY KEY,
    instrument_id VARCHAR(50) REFERENCES instruments(id),
    calibrated_by VARCHAR(50) NOT NULL,
    calibrated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(50) NOT NULL,
    result VARCHAR(50) NOT NULL,
    certificate_id VARCHAR(50),
    approved_by VARCHAR(50),
    approved_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS calibration_certificates (
    id VARCHAR(50) PRIMARY KEY,
    calibration_id VARCHAR(50) REFERENCES calibration_records(id),
    certificate_no VARCHAR(100) UNIQUE NOT NULL,
    issued_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    expiry_date TIMESTAMP WITH TIME ZONE NOT NULL,
    document_doc_ref VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS reference_standards (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    model_name VARCHAR(100) NOT NULL,
    serial_number VARCHAR(100) UNIQUE NOT NULL,
    last_calibrated TIMESTAMP WITH TIME ZONE NOT NULL,
    expiry_date TIMESTAMP WITH TIME ZONE NOT NULL,
    accuracy_class VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS measurement_results (
    id VARCHAR(50) PRIMARY KEY,
    calibration_id VARCHAR(50) REFERENCES calibration_records(id),
    test_point NUMERIC(15, 4) NOT NULL,
    nominal_value NUMERIC(15, 4) NOT NULL,
    standard_value NUMERIC(15, 4) NOT NULL,
    as_found_value NUMERIC(15, 4) NOT NULL,
    as_left_value NUMERIC(15, 4) NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tolerance_rules (
    id VARCHAR(50) PRIMARY KEY,
    instrument_type VARCHAR(50) NOT NULL,
    min_limit NUMERIC(10, 4) NOT NULL,
    max_limit NUMERIC(10, 4) NOT NULL,
    unit_of_measure VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS out_of_tolerance_cases (
    id VARCHAR(50) PRIMARY KEY,
    calibration_id VARCHAR(50) REFERENCES calibration_records(id),
    reported_by VARCHAR(50) NOT NULL,
    reported_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    impact_analysis TEXT NOT NULL,
    root_cause TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'INVESTIGATING'
);

CREATE TABLE IF NOT EXISTS traceability_records (
    id VARCHAR(50) PRIMARY KEY,
    reference_standard_id VARCHAR(50) REFERENCES reference_standards(id),
    primary_cert_no VARCHAR(100) NOT NULL,
    calibrated_by_body VARCHAR(100) NOT NULL,
    verified_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS laboratories (
    id VARCHAR(50) PRIMARY KEY,
    lab_name VARCHAR(100) NOT NULL,
    accreditation VARCHAR(100) NOT NULL,
    contact_person VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS calibration_history (
    id VARCHAR(50) PRIMARY KEY,
    instrument_id VARCHAR(50) REFERENCES instruments(id),
    calibration_id VARCHAR(50) REFERENCES calibration_records(id),
    event_description TEXT NOT NULL,
    recorded_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
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
