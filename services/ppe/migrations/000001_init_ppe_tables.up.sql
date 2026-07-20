CREATE TABLE IF NOT EXISTS ppe_catalog (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    model_name VARCHAR(100) NOT NULL,
    category_id VARCHAR(50) NOT NULL,
    manufacturer VARCHAR(100) NOT NULL,
    part_number VARCHAR(100) NOT NULL,
    standard_ref VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS ppe_categories (
    id VARCHAR(50) PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS ppe_items (
    id VARCHAR(50) PRIMARY KEY,
    ppe_id VARCHAR(50) REFERENCES ppe_catalog(id),
    serial_number VARCHAR(100) UNIQUE NOT NULL,
    rfid_code VARCHAR(100),
    barcode VARCHAR(100),
    manufacture_date TIMESTAMP WITH TIME ZONE NOT NULL,
    expiry_date TIMESTAMP WITH TIME ZONE NOT NULL,
    status VARCHAR(50) NOT NULL,
    issued_to VARCHAR(50),
    last_inspected_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS ppe_requirements (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    hazard_type VARCHAR(100) NOT NULL,
    work_area VARCHAR(100) NOT NULL,
    ppe_id_list TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS ppe_issues (
    id VARCHAR(50) PRIMARY KEY,
    item_id VARCHAR(50) REFERENCES ppe_items(id),
    issued_to_type VARCHAR(50) NOT NULL,
    issued_to_id VARCHAR(50) NOT NULL,
    issued_by VARCHAR(50) NOT NULL,
    issued_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    expected_return TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS ppe_returns (
    id VARCHAR(50) PRIMARY KEY,
    issue_id VARCHAR(50) NOT NULL,
    item_id VARCHAR(50) REFERENCES ppe_items(id),
    returned_by VARCHAR(50) NOT NULL,
    returned_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    condition VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS ppe_inspections (
    id VARCHAR(50) PRIMARY KEY,
    item_id VARCHAR(50) REFERENCES ppe_items(id),
    inspected_by VARCHAR(50) NOT NULL,
    inspected_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    result VARCHAR(50) NOT NULL,
    findings TEXT
);

CREATE TABLE IF NOT EXISTS ppe_maintenance (
    id VARCHAR(50) PRIMARY KEY,
    item_id VARCHAR(50) REFERENCES ppe_items(id),
    maintenance_by VARCHAR(50) NOT NULL,
    maintenance_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    cost NUMERIC(10, 2) DEFAULT 0,
    actions_taken TEXT NOT NULL,
    completed_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS ppe_certifications (
    id VARCHAR(50) PRIMARY KEY,
    ppe_id VARCHAR(50) REFERENCES ppe_catalog(id),
    certifying_body VARCHAR(100) NOT NULL,
    certification_code VARCHAR(100) NOT NULL,
    issued_date TIMESTAMP WITH TIME ZONE NOT NULL,
    expiry_date TIMESTAMP WITH TIME ZONE NOT NULL,
    document_doc_ref VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS ppe_replacements (
    id VARCHAR(50) PRIMARY KEY,
    item_id VARCHAR(50) REFERENCES ppe_items(id),
    trigger_reason VARCHAR(100) NOT NULL,
    scheduled_at TIMESTAMP WITH TIME ZONE NOT NULL,
    replaced_by VARCHAR(50),
    replaced_at TIMESTAMP WITH TIME ZONE,
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING'
);

CREATE TABLE IF NOT EXISTS inventory (
    id VARCHAR(50) PRIMARY KEY,
    ppe_id VARCHAR(50) REFERENCES ppe_catalog(id),
    plant_id VARCHAR(50) NOT NULL,
    quantity_on_hand INT NOT NULL DEFAULT 0,
    buffer_level INT NOT NULL DEFAULT 0,
    location VARCHAR(100) NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS stock_movements (
    id VARCHAR(50) PRIMARY KEY,
    ppe_id VARCHAR(50) REFERENCES ppe_catalog(id),
    movement_type VARCHAR(50) NOT NULL,
    quantity_change INT NOT NULL,
    recorded_by VARCHAR(50) NOT NULL,
    recorded_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    remarks TEXT
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
