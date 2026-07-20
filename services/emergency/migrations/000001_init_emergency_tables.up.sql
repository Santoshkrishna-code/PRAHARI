CREATE TABLE IF NOT EXISTS emergencies (
    id VARCHAR(50) PRIMARY KEY,
    emergency_number VARCHAR(100) UNIQUE NOT NULL,
    plant_id VARCHAR(50) NOT NULL,
    unit_id VARCHAR(50) NOT NULL,
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    category VARCHAR(50) NOT NULL,
    severity VARCHAR(50) NOT NULL,
    incident_id VARCHAR(50),
    status VARCHAR(50) NOT NULL,
    commander_id VARCHAR(50) NOT NULL,
    declared_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    command_established_at TIMESTAMP WITH TIME ZONE,
    stabilized_at TIMESTAMP WITH TIME ZONE,
    closed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS emergency_types (
    id VARCHAR(50) PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,
    description TEXT,
    severity VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS response_plans (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    plan_number VARCHAR(100) UNIQUE NOT NULL,
    title VARCHAR(200) NOT NULL,
    category VARCHAR(50) NOT NULL,
    procedures TEXT NOT NULL,
    version VARCHAR(20) NOT NULL,
    approved_by VARCHAR(50) NOT NULL,
    approved_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS incident_commands (
    id VARCHAR(50) PRIMARY KEY,
    emergency_id VARCHAR(50) REFERENCES emergencies(id),
    commander_id VARCHAR(50) NOT NULL,
    command_post VARCHAR(200) NOT NULL,
    established_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS command_roles (
    id VARCHAR(50) PRIMARY KEY,
    command_id VARCHAR(50) REFERENCES incident_commands(id),
    role_name VARCHAR(100) NOT NULL,
    assignee_id VARCHAR(50) NOT NULL,
    assigned_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS emergency_teams (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    team_name VARCHAR(100) NOT NULL,
    team_type VARCHAR(50) NOT NULL,
    leader_id VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS responders (
    id VARCHAR(50) PRIMARY KEY,
    team_id VARCHAR(50) REFERENCES emergency_teams(id),
    user_id VARCHAR(50) NOT NULL,
    role VARCHAR(100) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'STANDBY',
    deployed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS evacuations (
    id VARCHAR(50) PRIMARY KEY,
    emergency_id VARCHAR(50) REFERENCES emergencies(id),
    zone_id VARCHAR(50) NOT NULL,
    initiated_by VARCHAR(50) NOT NULL,
    total_personnel INT NOT NULL DEFAULT 0,
    accounted_for INT NOT NULL DEFAULT 0,
    missing_count INT NOT NULL DEFAULT 0,
    status VARCHAR(50) NOT NULL DEFAULT 'IN_PROGRESS',
    evacuation_time_sec NUMERIC(10, 2) DEFAULT 0,
    initiated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS assembly_points (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    point_code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    location VARCHAR(200) NOT NULL,
    capacity INT NOT NULL DEFAULT 100,
    current_count INT NOT NULL DEFAULT 0,
    warden_user_id VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS emergency_resources (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    resource_code VARCHAR(100) UNIQUE NOT NULL,
    resource_type VARCHAR(100) NOT NULL,
    quantity INT NOT NULL DEFAULT 1,
    available_qty INT NOT NULL DEFAULT 1,
    location_code VARCHAR(100) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'READY',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS mutual_aid_agreements (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    partner_name VARCHAR(200) NOT NULL,
    contact_phone VARCHAR(50) NOT NULL,
    aid_provided TEXT NOT NULL,
    active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS communications (
    id VARCHAR(50) PRIMARY KEY,
    emergency_id VARCHAR(50) REFERENCES emergencies(id),
    sender_id VARCHAR(50) NOT NULL,
    channel VARCHAR(50) NOT NULL,
    message TEXT NOT NULL,
    sent_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS notifications (
    id VARCHAR(50) PRIMARY KEY,
    emergency_id VARCHAR(50) REFERENCES emergencies(id),
    recipient VARCHAR(100) NOT NULL,
    alert_type VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    sent_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS drills (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    title VARCHAR(200) NOT NULL,
    drill_type VARCHAR(50) NOT NULL,
    scheduled_at TIMESTAMP WITH TIME ZONE NOT NULL,
    executed_at TIMESTAMP WITH TIME ZONE,
    duration_min NUMERIC(10, 2) DEFAULT 0,
    passed BOOLEAN DEFAULT true,
    evaluator_id VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'SCHEDULED',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS exercises (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    title VARCHAR(200) NOT NULL,
    exercise_type VARCHAR(50) NOT NULL,
    objectives TEXT NOT NULL,
    conducted_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS recoveries (
    id VARCHAR(50) PRIMARY KEY,
    emergency_id VARCHAR(50) REFERENCES emergencies(id),
    title VARCHAR(200) NOT NULL,
    damage_summary TEXT NOT NULL,
    estimated_cost NUMERIC(15, 2) DEFAULT 0,
    status VARCHAR(50) NOT NULL DEFAULT 'IN_PROGRESS',
    target_complete TIMESTAMP WITH TIME ZONE NOT NULL,
    completed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS after_action_reviews (
    id VARCHAR(50) PRIMARY KEY,
    emergency_id VARCHAR(50) REFERENCES emergencies(id),
    drill_id VARCHAR(50) REFERENCES drills(id),
    title VARCHAR(200) NOT NULL,
    summary TEXT NOT NULL,
    what_went_well TEXT,
    improvements TEXT,
    facilitator VARCHAR(50) NOT NULL,
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
