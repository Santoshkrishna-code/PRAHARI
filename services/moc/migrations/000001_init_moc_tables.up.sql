CREATE TABLE IF NOT EXISTS change_requests (
    id VARCHAR(50) PRIMARY KEY,
    moc_number VARCHAR(100) UNIQUE NOT NULL,
    plant_id VARCHAR(50) NOT NULL,
    department_id VARCHAR(50),
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    reason_for_change TEXT NOT NULL,
    category VARCHAR(50) NOT NULL,
    change_type VARCHAR(50) NOT NULL,
    target_asset_id VARCHAR(50),
    risk_level VARCHAR(20) NOT NULL DEFAULT 'MEDIUM',
    status VARCHAR(50) NOT NULL,
    requester_id VARCHAR(50) NOT NULL,
    target_date TIMESTAMP WITH TIME ZONE NOT NULL,
    expiry_date TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS change_types (
    id VARCHAR(50) PRIMARY KEY,
    type_code VARCHAR(50) UNIQUE NOT NULL,
    type_name VARCHAR(100) NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS impact_assessments (
    id VARCHAR(50) PRIMARY KEY,
    change_request_id VARCHAR(50) REFERENCES change_requests(id),
    safety_impact BOOLEAN DEFAULT false,
    environmental_impact BOOLEAN DEFAULT false,
    quality_impact BOOLEAN DEFAULT false,
    reliability_impact BOOLEAN DEFAULT false,
    cybersecurity_impact BOOLEAN DEFAULT false,
    regulatory_impact BOOLEAN DEFAULT false,
    p_and_id_impact BOOLEAN DEFAULT false,
    hazop_required BOOLEAN DEFAULT false,
    summary_notes TEXT,
    assessed_by VARCHAR(50) NOT NULL,
    assessed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS technical_reviews (
    id VARCHAR(50) PRIMARY KEY,
    change_request_id VARCHAR(50) REFERENCES change_requests(id),
    discipline VARCHAR(50) NOT NULL,
    reviewer_id VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    findings TEXT NOT NULL,
    conditions TEXT,
    reviewed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS risk_reviews (
    id VARCHAR(50) PRIMARY KEY,
    change_request_id VARCHAR(50) REFERENCES change_requests(id),
    risk_assessment_id VARCHAR(50),
    pre_change_risk VARCHAR(20) NOT NULL,
    post_change_risk VARCHAR(20) NOT NULL,
    mitigations_reqd TEXT,
    risk_manager_id VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    reviewed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS safety_reviews (
    id VARCHAR(50) PRIMARY KEY,
    change_request_id VARCHAR(50) REFERENCES change_requests(id),
    psm_impact_verified BOOLEAN DEFAULT false,
    occupational_health BOOLEAN DEFAULT false,
    emergency_response BOOLEAN DEFAULT false,
    safety_officer_id VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    comments TEXT,
    reviewed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS environmental_reviews (
    id VARCHAR(50) PRIMARY KEY,
    change_request_id VARCHAR(50) REFERENCES change_requests(id),
    emissions_impact BOOLEAN DEFAULT false,
    waste_impact BOOLEAN DEFAULT false,
    permit_revision_required BOOLEAN DEFAULT false,
    reviewer_id VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    reviewed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS compliance_reviews (
    id VARCHAR(50) PRIMARY KEY,
    change_request_id VARCHAR(50) REFERENCES change_requests(id),
    statutory_approval_required BOOLEAN DEFAULT false,
    standard_clause VARCHAR(100),
    reviewer_id VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    reviewed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS approvals (
    id VARCHAR(50) PRIMARY KEY,
    change_request_id VARCHAR(50) REFERENCES change_requests(id),
    approver_id VARCHAR(50) NOT NULL,
    role VARCHAR(50) NOT NULL,
    decision VARCHAR(50) NOT NULL,
    comments TEXT,
    approved_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS implementations (
    id VARCHAR(50) PRIMARY KEY,
    change_request_id VARCHAR(50) REFERENCES change_requests(id),
    work_order_id VARCHAR(50),
    permit_id VARCHAR(50),
    implemented_by VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    completed_date TIMESTAMP WITH TIME ZONE,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS verifications (
    id VARCHAR(50) PRIMARY KEY,
    change_request_id VARCHAR(50) REFERENCES change_requests(id),
    pssr_completed BOOLEAN DEFAULT false,
    training_verified BOOLEAN DEFAULT false,
    docs_updated BOOLEAN DEFAULT false,
    verified_by VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    comments TEXT,
    verified_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS rollback_plans (
    id VARCHAR(50) PRIMARY KEY,
    change_request_id VARCHAR(50) REFERENCES change_requests(id),
    trigger_reason TEXT NOT NULL,
    reversion_steps TEXT NOT NULL,
    executed_by VARCHAR(50),
    status VARCHAR(50) NOT NULL,
    executed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS training_verifications (
    id VARCHAR(50) PRIMARY KEY,
    change_request_id VARCHAR(50) REFERENCES change_requests(id),
    course_id VARCHAR(50),
    trained_count INT NOT NULL DEFAULT 0,
    target_count INT NOT NULL DEFAULT 0,
    is_complete BOOLEAN DEFAULT false,
    verified_by VARCHAR(50) NOT NULL,
    verified_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS document_updates (
    id VARCHAR(50) PRIMARY KEY,
    change_request_id VARCHAR(50) REFERENCES change_requests(id),
    document_number VARCHAR(100) NOT NULL,
    document_title VARCHAR(200) NOT NULL,
    new_revision VARCHAR(20) NOT NULL,
    is_updated BOOLEAN DEFAULT false,
    updated_by VARCHAR(50) NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS communications (
    id VARCHAR(50) PRIMARY KEY,
    change_request_id VARCHAR(50) REFERENCES change_requests(id),
    recipient_group VARCHAR(100) NOT NULL,
    subject VARCHAR(200) NOT NULL,
    sent_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
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
