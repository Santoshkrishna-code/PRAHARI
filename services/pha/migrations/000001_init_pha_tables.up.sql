CREATE TABLE IF NOT EXISTS pha_studies (
    id VARCHAR(50) PRIMARY KEY,
    study_number VARCHAR(100) UNIQUE NOT NULL,
    plant_id VARCHAR(50) NOT NULL,
    unit_id VARCHAR(50) NOT NULL,
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    method VARCHAR(50) NOT NULL,
    moc_id VARCHAR(50),
    status VARCHAR(50) NOT NULL,
    leader_id VARCHAR(50) NOT NULL,
    scribe_id VARCHAR(50) NOT NULL,
    target_date TIMESTAMP WITH TIME ZONE NOT NULL,
    revalidation_due_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS process_nodes (
    id VARCHAR(50) PRIMARY KEY,
    study_id VARCHAR(50) REFERENCES pha_studies(id),
    node_number INT NOT NULL,
    node_name VARCHAR(200) NOT NULL,
    design_intent TEXT NOT NULL,
    p_and_id_number VARCHAR(100) NOT NULL,
    operating_temp_c NUMERIC(10, 2) NOT NULL DEFAULT 0,
    operating_press_bar NUMERIC(10, 2) NOT NULL DEFAULT 0,
    location_code VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS guide_words (
    id VARCHAR(50) PRIMARY KEY,
    word VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS deviations (
    id VARCHAR(50) PRIMARY KEY,
    node_id VARCHAR(50) REFERENCES process_nodes(id),
    guide_word VARCHAR(50) NOT NULL,
    parameter VARCHAR(50) NOT NULL,
    deviation_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS hazard_scenarios (
    id VARCHAR(50) PRIMARY KEY,
    node_id VARCHAR(50) REFERENCES process_nodes(id),
    deviation_id VARCHAR(50) REFERENCES deviations(id),
    cause_description TEXT NOT NULL,
    severity INT NOT NULL DEFAULT 1,
    likelihood INT NOT NULL DEFAULT 1,
    risk_rank INT NOT NULL DEFAULT 1,
    risk_category VARCHAR(50) NOT NULL,
    is_sil_relevant BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS consequences (
    id VARCHAR(50) PRIMARY KEY,
    scenario_id VARCHAR(50) REFERENCES hazard_scenarios(id),
    category VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    severity_score INT NOT NULL DEFAULT 1,
    unmitigated_risk_rank INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS safeguards (
    id VARCHAR(50) PRIMARY KEY,
    scenario_id VARCHAR(50) REFERENCES hazard_scenarios(id),
    title VARCHAR(200) NOT NULL,
    safeguard_type VARCHAR(50) NOT NULL,
    is_ipl BOOLEAN DEFAULT false,
    pfd NUMERIC(10, 6) NOT NULL DEFAULT 1.0,
    asset_id VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS recommendations (
    id VARCHAR(50) PRIMARY KEY,
    study_id VARCHAR(50) REFERENCES pha_studies(id),
    scenario_id VARCHAR(50) REFERENCES hazard_scenarios(id),
    rec_number VARCHAR(50) UNIQUE NOT NULL,
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    priority VARCHAR(20) NOT NULL DEFAULT 'MEDIUM',
    target_sil VARCHAR(20),
    status VARCHAR(50) NOT NULL DEFAULT 'OPEN',
    assigned_to VARCHAR(50) NOT NULL,
    target_date TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS action_items (
    id VARCHAR(50) PRIMARY KEY,
    recommendation_id VARCHAR(50) REFERENCES recommendations(id),
    action_title VARCHAR(200) NOT NULL,
    assignee_id VARCHAR(50) NOT NULL,
    work_order_id VARCHAR(50),
    status VARCHAR(50) NOT NULL DEFAULT 'OPEN',
    due_date TIMESTAMP WITH TIME ZONE NOT NULL,
    completed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS hazop_sessions (
    id VARCHAR(50) PRIMARY KEY,
    study_id VARCHAR(50) REFERENCES pha_studies(id),
    session_date TIMESTAMP WITH TIME ZONE NOT NULL,
    duration_hrs NUMERIC(5, 2) NOT NULL DEFAULT 0,
    attendees TEXT NOT NULL,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS lopa_analyses (
    id VARCHAR(50) PRIMARY KEY,
    study_id VARCHAR(50) REFERENCES pha_studies(id),
    scenario_id VARCHAR(50) REFERENCES hazard_scenarios(id),
    initiating_event_freq NUMERIC(15, 6) NOT NULL,
    tolerable_target_freq NUMERIC(15, 6) NOT NULL,
    total_ipl_mitigation NUMERIC(15, 6) NOT NULL,
    mitigated_event_freq NUMERIC(15, 6) NOT NULL,
    required_rrf NUMERIC(15, 2) NOT NULL,
    target_sil VARCHAR(20) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS bowtie_analyses (
    id VARCHAR(50) PRIMARY KEY,
    study_id VARCHAR(50) REFERENCES pha_studies(id),
    top_event VARCHAR(200) NOT NULL,
    hazard_description TEXT NOT NULL,
    preventive_barriers TEXT NOT NULL,
    mitigation_barriers TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS fmea_analyses (
    id VARCHAR(50) PRIMARY KEY,
    study_id VARCHAR(50) REFERENCES pha_studies(id),
    equipment_item VARCHAR(200) NOT NULL,
    failure_mode VARCHAR(200) NOT NULL,
    failure_effect TEXT NOT NULL,
    severity INT NOT NULL DEFAULT 1,
    occurrence INT NOT NULL DEFAULT 1,
    detection INT NOT NULL DEFAULT 1,
    rpn INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS whatif_analyses (
    id VARCHAR(50) PRIMARY KEY,
    study_id VARCHAR(50) REFERENCES pha_studies(id),
    question TEXT NOT NULL,
    consequence TEXT NOT NULL,
    safeguards TEXT NOT NULL,
    recommendation TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS checklist_analyses (
    id VARCHAR(50) PRIMARY KEY,
    study_id VARCHAR(50) REFERENCES pha_studies(id),
    checklist_group VARCHAR(100) NOT NULL,
    item_question TEXT NOT NULL,
    is_compliant BOOLEAN DEFAULT true,
    comments TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS pssr_reviews (
    id VARCHAR(50) PRIMARY KEY,
    study_id VARCHAR(50) REFERENCES pha_studies(id),
    moc_id VARCHAR(50),
    pssr_title VARCHAR(200) NOT NULL,
    construction_ok BOOLEAN DEFAULT false,
    procedures_ok BOOLEAN DEFAULT false,
    training_ok BOOLEAN DEFAULT false,
    pha_action_items_ok BOOLEAN DEFAULT false,
    status VARCHAR(50) NOT NULL,
    verified_by VARCHAR(50) NOT NULL,
    verified_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
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
