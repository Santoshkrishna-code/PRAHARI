CREATE TABLE IF NOT EXISTS visitors (
    id VARCHAR(50) PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    company VARCHAR(100) NOT NULL,
    visitor_type VARCHAR(50) NOT NULL,
    id_type VARCHAR(50) NOT NULL,
    id_number VARCHAR(100) NOT NULL,
    blacklisted BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS visits (
    id VARCHAR(50) PRIMARY KEY,
    visitor_id VARCHAR(50) REFERENCES visitors(id),
    host_id VARCHAR(50) NOT NULL,
    plant_id VARCHAR(50) NOT NULL,
    purpose TEXT NOT NULL,
    scheduled_in TIMESTAMP WITH TIME ZONE NOT NULL,
    scheduled_out TIMESTAMP WITH TIME ZONE NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS hosts (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    host_name VARCHAR(100) NOT NULL,
    department VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS gate_passes (
    id VARCHAR(50) PRIMARY KEY,
    visit_id VARCHAR(50) REFERENCES visits(id),
    pass_number VARCHAR(100) UNIQUE NOT NULL,
    qr_code_url TEXT NOT NULL,
    issued_by VARCHAR(50) NOT NULL,
    issued_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    valid_until TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS badges (
    id VARCHAR(50) PRIMARY KEY,
    badge_number VARCHAR(50) UNIQUE NOT NULL,
    rfid_code VARCHAR(100),
    status VARCHAR(50) NOT NULL DEFAULT 'AVAILABLE',
    issued_to VARCHAR(50),
    issued_at TIMESTAMP WITH TIME ZONE,
    returned_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS vehicles (
    id VARCHAR(50) PRIMARY KEY,
    visit_id VARCHAR(50) REFERENCES visits(id),
    license_plate VARCHAR(50) NOT NULL,
    vehicle_type VARCHAR(50) NOT NULL,
    driver_name VARCHAR(100) NOT NULL,
    approved BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS inductions (
    id VARCHAR(50) PRIMARY KEY,
    visitor_id VARCHAR(50) REFERENCES visitors(id),
    induction_type VARCHAR(100) NOT NULL,
    completed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    score NUMERIC(5, 2) DEFAULT 0
);

CREATE TABLE IF NOT EXISTS document_verifications (
    id VARCHAR(50) PRIMARY KEY,
    visitor_id VARCHAR(50) REFERENCES visitors(id),
    document_name VARCHAR(200) NOT NULL,
    verified BOOLEAN DEFAULT false,
    verified_by VARCHAR(50) NOT NULL,
    verified_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    document_doc_ref VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS checkins (
    id VARCHAR(50) PRIMARY KEY,
    visit_id VARCHAR(50) REFERENCES visits(id),
    security_check_point VARCHAR(100) NOT NULL,
    gate_number VARCHAR(50),
    check_in_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    checked_in_by VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS checkouts (
    id VARCHAR(50) PRIMARY KEY,
    visit_id VARCHAR(50) REFERENCES visits(id),
    check_out_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    checked_out_by VARCHAR(50) NOT NULL,
    badge_returned BOOLEAN DEFAULT true
);

CREATE TABLE IF NOT EXISTS contractor_entries (
    id VARCHAR(50) PRIMARY KEY,
    visit_id VARCHAR(50) REFERENCES visits(id),
    contractor_id VARCHAR(50) NOT NULL,
    work_order_ref VARCHAR(100),
    permit_ref VARCHAR(100),
    verified BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS security_checkpoints (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    checkpoint_name VARCHAR(100) NOT NULL,
    location VARCHAR(200) NOT NULL,
    active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS visitor_categories (
    id VARCHAR(50) PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    induction_required BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS blacklists (
    id VARCHAR(50) PRIMARY KEY,
    id_type VARCHAR(50) NOT NULL,
    id_number VARCHAR(100) UNIQUE NOT NULL,
    reason TEXT NOT NULL,
    blocked_by VARCHAR(50) NOT NULL,
    blocked_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
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

CREATE TABLE IF NOT EXISTS emergencymuster_records (
    id VARCHAR(50) PRIMARY KEY,
    muster_id VARCHAR(50) NOT NULL,
    visitor_id VARCHAR(50) NOT NULL,
    assembly_point VARCHAR(100),
    accounted_for BOOLEAN DEFAULT false,
    accounted_at TIMESTAMP WITH TIME ZONE,
    warden_id VARCHAR(50)
);
