CREATE TABLE IF NOT EXISTS isolation_plans (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    equipment_id VARCHAR(50) NOT NULL,
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    approved_by VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS isolation_certificates (
    id VARCHAR(50) PRIMARY KEY,
    plan_id VARCHAR(50) REFERENCES isolation_plans(id),
    permit_id VARCHAR(50),
    issuer_id VARCHAR(50) NOT NULL,
    receiver_id VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    verified_at TIMESTAMP WITH TIME ZONE,
    restored_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS energy_sources (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    energy_type VARCHAR(50) NOT NULL,
    magnitude VARCHAR(100) NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS isolation_points (
    id VARCHAR(50) PRIMARY KEY,
    equipment_id VARCHAR(50) NOT NULL,
    source_id VARCHAR(50) REFERENCES energy_sources(id),
    point_name VARCHAR(100) NOT NULL,
    location VARCHAR(200) NOT NULL,
    isolation_method VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS locks (
    id VARCHAR(50) PRIMARY KEY,
    lock_number VARCHAR(50) UNIQUE NOT NULL,
    color VARCHAR(30) NOT NULL,
    status VARCHAR(50) NOT NULL,
    assigned_to VARCHAR(50),
    applied_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS tags (
    id VARCHAR(50) PRIMARY KEY,
    tag_number VARCHAR(50) UNIQUE NOT NULL,
    status VARCHAR(50) NOT NULL,
    details TEXT,
    assigned_to VARCHAR(50),
    applied_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS group_lockboxes (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    box_number VARCHAR(50) UNIQUE NOT NULL,
    location VARCHAR(200) NOT NULL,
    master_key_id VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS zero_energy_verifications (
    id VARCHAR(50) PRIMARY KEY,
    certificate_id VARCHAR(50) REFERENCES isolation_certificates(id),
    verified_by VARCHAR(50) NOT NULL,
    verification_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    test_passed BOOLEAN NOT NULL DEFAULT false,
    test_method VARCHAR(100) NOT NULL,
    notes TEXT
);

CREATE TABLE IF NOT EXISTS restoration_records (
    id VARCHAR(50) PRIMARY KEY,
    certificate_id VARCHAR(50) REFERENCES isolation_certificates(id),
    restored_by VARCHAR(50) NOT NULL,
    restored_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    details TEXT NOT NULL,
    confirmed_safe BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS loto_audits (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    auditor_id VARCHAR(50) NOT NULL,
    audited_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    result_score NUMERIC(5, 2) NOT NULL,
    compliance BOOLEAN NOT NULL DEFAULT false,
    notes TEXT
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
