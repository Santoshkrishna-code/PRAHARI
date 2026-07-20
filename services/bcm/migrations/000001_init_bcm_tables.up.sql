CREATE TABLE IF NOT EXISTS continuity_plans (
    id VARCHAR(50) PRIMARY KEY,
    plan_number VARCHAR(100) UNIQUE NOT NULL,
    plant_id VARCHAR(50) NOT NULL,
    business_unit VARCHAR(100) NOT NULL,
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    scope TEXT NOT NULL,
    version VARCHAR(20) NOT NULL,
    status VARCHAR(50) NOT NULL,
    approved_by VARCHAR(50) NOT NULL,
    approved_at TIMESTAMP WITH TIME ZONE,
    next_review_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS critical_processes (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    process_code VARCHAR(100) UNIQUE NOT NULL,
    process_name VARCHAR(200) NOT NULL,
    owner_id VARCHAR(50) NOT NULL,
    priority_tier VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS business_impact_analyses (
    id VARCHAR(50) PRIMARY KEY,
    plan_id VARCHAR(50) REFERENCES continuity_plans(id),
    process_id VARCHAR(50) REFERENCES critical_processes(id),
    financial_loss_per_day NUMERIC(15, 2) DEFAULT 0,
    operational_impact VARCHAR(50) NOT NULL,
    regulatory_impact VARCHAR(100) NOT NULL,
    maximum_tolerable_downtime_hrs NUMERIC(10, 2) NOT NULL,
    rto_hrs NUMERIC(10, 2) NOT NULL,
    rpo_hrs NUMERIC(10, 2) NOT NULL,
    evaluated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS recovery_objectives (
    id VARCHAR(50) PRIMARY KEY,
    process_id VARCHAR(50) REFERENCES critical_processes(id),
    target_rto_hrs NUMERIC(10, 2) NOT NULL,
    target_rpo_hrs NUMERIC(10, 2) NOT NULL,
    target_mbco_pct NUMERIC(5, 2) NOT NULL DEFAULT 50.0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS continuity_strategies (
    id VARCHAR(50) PRIMARY KEY,
    plan_id VARCHAR(50) REFERENCES continuity_plans(id),
    strategy_type VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    cost_estimate NUMERIC(15, 2) DEFAULT 0,
    approved BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS disaster_recovery_plans (
    id VARCHAR(50) PRIMARY KEY,
    plan_id VARCHAR(50) REFERENCES continuity_plans(id),
    system_name VARCHAR(200) NOT NULL,
    dr_site_location VARCHAR(200) NOT NULL,
    failover_type VARCHAR(50) NOT NULL,
    target_rto_hrs NUMERIC(10, 2) NOT NULL,
    target_rpo_hrs NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS recovery_teams (
    id VARCHAR(50) PRIMARY KEY,
    plan_id VARCHAR(50) REFERENCES continuity_plans(id),
    team_name VARCHAR(100) NOT NULL,
    leader_id VARCHAR(50) NOT NULL,
    role VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS alternate_facilities (
    id VARCHAR(50) PRIMARY KEY,
    facility_code VARCHAR(100) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,
    facility_type VARCHAR(50) NOT NULL,
    location VARCHAR(200) NOT NULL,
    seating_cap INT NOT NULL DEFAULT 50,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS dependencies (
    id VARCHAR(50) PRIMARY KEY,
    process_id VARCHAR(50) REFERENCES critical_processes(id),
    depends_on_type VARCHAR(50) NOT NULL,
    depends_on_id VARCHAR(50) NOT NULL,
    criticality VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS supplier_continuity (
    id VARCHAR(50) PRIMARY KEY,
    supplier_id VARCHAR(50) NOT NULL,
    supplier_name VARCHAR(200) NOT NULL,
    resilience_score NUMERIC(5, 2) DEFAULT 100.0,
    sla_rto_hrs NUMERIC(10, 2) NOT NULL,
    last_audited_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS communication_plans (
    id VARCHAR(50) PRIMARY KEY,
    plan_id VARCHAR(50) REFERENCES continuity_plans(id),
    spokesperson_id VARCHAR(50) NOT NULL,
    channels TEXT NOT NULL,
    templates TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS continuity_exercises (
    id VARCHAR(50) PRIMARY KEY,
    plan_id VARCHAR(50) REFERENCES continuity_plans(id),
    title VARCHAR(200) NOT NULL,
    type VARCHAR(50) NOT NULL,
    scheduled_at TIMESTAMP WITH TIME ZONE NOT NULL,
    executed_at TIMESTAMP WITH TIME ZONE,
    passed BOOLEAN DEFAULT true,
    rto_achieved_hrs NUMERIC(10, 2) DEFAULT 0,
    status VARCHAR(50) NOT NULL DEFAULT 'SCHEDULED',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tabletop_exercises (
    id VARCHAR(50) PRIMARY KEY,
    exercise_id VARCHAR(50) REFERENCES continuity_exercises(id),
    scenario TEXT NOT NULL,
    moderator_id VARCHAR(50) NOT NULL,
    conducted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS failover_tests (
    id VARCHAR(50) PRIMARY KEY,
    dr_plan_id VARCHAR(50) REFERENCES disaster_recovery_plans(id),
    test_name VARCHAR(200) NOT NULL,
    passed BOOLEAN DEFAULT true,
    rto_achieved_hrs NUMERIC(10, 2) NOT NULL,
    rpo_achieved_hrs NUMERIC(10, 2) NOT NULL,
    executed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS resilience_assessments (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    business_unit VARCHAR(100) NOT NULL,
    resilience_index_pct NUMERIC(5, 2) DEFAULT 100.0,
    iso22301_status VARCHAR(50) NOT NULL,
    assessed_by VARCHAR(50) NOT NULL,
    assessed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS after_action_reviews (
    id VARCHAR(50) PRIMARY KEY,
    plan_id VARCHAR(50) REFERENCES continuity_plans(id),
    exercise_id VARCHAR(50) REFERENCES continuity_exercises(id),
    title VARCHAR(200) NOT NULL,
    summary TEXT NOT NULL,
    what_went_well TEXT,
    improvements TEXT,
    facilitator_id VARCHAR(50) NOT NULL,
    reviewed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS lessons_learned (
    id VARCHAR(50) PRIMARY KEY,
    review_id VARCHAR(50) REFERENCES after_action_reviews(id),
    category VARCHAR(50) NOT NULL,
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    action_needed TEXT NOT NULL,
    assigned_to VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'OPEN',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
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
