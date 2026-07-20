CREATE TABLE IF NOT EXISTS documents (
    id VARCHAR(50) PRIMARY KEY,
    document_number VARCHAR(100) UNIQUE NOT NULL,
    plant_id VARCHAR(50) NOT NULL,
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    category_id VARCHAR(50) NOT NULL,
    document_type VARCHAR(50) NOT NULL,
    current_version VARCHAR(20) NOT NULL DEFAULT '1.0',
    status VARCHAR(50) NOT NULL,
    checked_out_by VARCHAR(50),
    checked_out_at TIMESTAMP WITH TIME ZONE,
    owner_id VARCHAR(50) NOT NULL,
    review_cycle_months INT NOT NULL DEFAULT 12,
    next_review_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS document_versions (
    id VARCHAR(50) PRIMARY KEY,
    document_id VARCHAR(50) REFERENCES documents(id),
    version_number VARCHAR(20) NOT NULL,
    file_url TEXT NOT NULL,
    file_hash VARCHAR(100) NOT NULL,
    change_summary TEXT NOT NULL,
    created_by VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS document_categories (
    id VARCHAR(50) PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    parent_id VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS document_templates (
    id VARCHAR(50) PRIMARY KEY,
    template_name VARCHAR(100) NOT NULL,
    document_type VARCHAR(50) NOT NULL,
    file_url TEXT NOT NULL,
    active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS document_metadata (
    id VARCHAR(50) PRIMARY KEY,
    document_id VARCHAR(50) REFERENCES documents(id),
    meta_key VARCHAR(100) NOT NULL,
    meta_value TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS document_relationships (
    id VARCHAR(50) PRIMARY KEY,
    source_doc_id VARCHAR(50) REFERENCES documents(id),
    target_doc_id VARCHAR(50) REFERENCES documents(id),
    relation_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS document_reviews (
    id VARCHAR(50) PRIMARY KEY,
    document_id VARCHAR(50) REFERENCES documents(id),
    reviewer_id VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    comments TEXT,
    reviewed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS document_approvals (
    id VARCHAR(50) PRIMARY KEY,
    document_id VARCHAR(50) REFERENCES documents(id),
    version_id VARCHAR(50) NOT NULL,
    approver_id VARCHAR(50) NOT NULL,
    approved BOOLEAN DEFAULT true,
    approved_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    comments TEXT
);

CREATE TABLE IF NOT EXISTS document_distributions (
    id VARCHAR(50) PRIMARY KEY,
    document_id VARCHAR(50) REFERENCES documents(id),
    version_id VARCHAR(50) NOT NULL,
    recipient_id VARCHAR(50) NOT NULL,
    distributed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    acknowledged_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS controlled_copies (
    id VARCHAR(50) PRIMARY KEY,
    document_id VARCHAR(50) REFERENCES documents(id),
    version_id VARCHAR(50) NOT NULL,
    copy_number VARCHAR(50) UNIQUE NOT NULL,
    issued_to VARCHAR(50) NOT NULL,
    location VARCHAR(200) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'ACTIVE',
    issued_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS electronic_signatures (
    id VARCHAR(50) PRIMARY KEY,
    document_id VARCHAR(50) REFERENCES documents(id),
    version_id VARCHAR(50) NOT NULL,
    signer_id VARCHAR(50) NOT NULL,
    signer_role VARCHAR(100) NOT NULL,
    signature_hash TEXT NOT NULL,
    signed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS retention_policies (
    id VARCHAR(50) PRIMARY KEY,
    category_id VARCHAR(50) REFERENCES document_categories(id),
    document_type VARCHAR(50) NOT NULL,
    retention_years INT NOT NULL DEFAULT 5,
    action_on_expiry VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS records (
    id VARCHAR(50) PRIMARY KEY,
    document_id VARCHAR(50) REFERENCES documents(id),
    archived_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    archived_by VARCHAR(50) NOT NULL,
    storage_class VARCHAR(50) NOT NULL DEFAULT 'GLACIER'
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
