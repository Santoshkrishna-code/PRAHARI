CREATE TABLE IF NOT EXISTS shifts (
    id VARCHAR(50) PRIMARY KEY,
    shift_name VARCHAR(100) NOT NULL,
    plant_id VARCHAR(50) NOT NULL,
    unit_id VARCHAR(50) NOT NULL,
    supervisor_id VARCHAR(50) NOT NULL,
    scheduled_start TIMESTAMP WITH TIME ZONE NOT NULL,
    scheduled_end TIMESTAMP WITH TIME ZONE NOT NULL,
    actual_start TIMESTAMP WITH TIME ZONE,
    actual_end TIMESTAMP WITH TIME ZONE,
    status VARCHAR(50) NOT NULL,
    handover_id VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS shift_calendars (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    pattern_name VARCHAR(100) NOT NULL,
    cycle_days INT NOT NULL DEFAULT 7,
    description TEXT,
    active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS shift_rosters (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    calendar_id VARCHAR(50) REFERENCES shift_calendars(id),
    is_published BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS crews (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    crew_name VARCHAR(100) NOT NULL,
    lead_id VARCHAR(50) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS crew_assignments (
    id VARCHAR(50) PRIMARY KEY,
    crew_id VARCHAR(50) REFERENCES crews(id),
    user_id VARCHAR(50) NOT NULL,
    role VARCHAR(100) NOT NULL,
    assigned_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS shift_logs (
    id VARCHAR(50) PRIMARY KEY,
    shift_id VARCHAR(50) REFERENCES shifts(id),
    logged_by_id VARCHAR(50) NOT NULL,
    category VARCHAR(50) NOT NULL,
    log_entry TEXT NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    is_critical BOOLEAN DEFAULT false,
    asset_id VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS operator_journals (
    id VARCHAR(50) PRIMARY KEY,
    shift_id VARCHAR(50) REFERENCES shifts(id),
    operator_id VARCHAR(50) NOT NULL,
    subject VARCHAR(200) NOT NULL,
    detail TEXT NOT NULL,
    logged_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS handovers (
    id VARCHAR(50) PRIMARY KEY,
    outgoing_shift_id VARCHAR(50) NOT NULL,
    incoming_shift_id VARCHAR(50) NOT NULL,
    outgoing_lead_id VARCHAR(50) NOT NULL,
    incoming_lead_id VARCHAR(50) NOT NULL,
    open_permit_ids TEXT,
    active_maintenance TEXT,
    safety_incidents TEXT,
    operational_continuity_notes TEXT,
    initiated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    accepted_at TIMESTAMP WITH TIME ZONE,
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING'
);

CREATE TABLE IF NOT EXISTS outstanding_actions (
    id VARCHAR(50) PRIMARY KEY,
    shift_id VARCHAR(50) REFERENCES shifts(id),
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    assigned_to VARCHAR(50) NOT NULL,
    due_date TIMESTAMP WITH TIME ZONE NOT NULL,
    completed_at TIMESTAMP WITH TIME ZONE,
    status VARCHAR(50) NOT NULL DEFAULT 'OPEN'
);

CREATE TABLE IF NOT EXISTS communications (
    id VARCHAR(50) PRIMARY KEY,
    shift_id VARCHAR(50) REFERENCES shifts(id),
    sender_id VARCHAR(50) NOT NULL,
    channel VARCHAR(50) NOT NULL,
    content TEXT NOT NULL,
    sent_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS production_summaries (
    id VARCHAR(50) PRIMARY KEY,
    shift_id VARCHAR(50) REFERENCES shifts(id),
    output_quantity NUMERIC(15, 4) DEFAULT 0,
    unit_of_measure VARCHAR(50) NOT NULL,
    downtime_hours NUMERIC(10, 2) DEFAULT 0,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS attendance (
    id VARCHAR(50) PRIMARY KEY,
    shift_id VARCHAR(50) REFERENCES shifts(id),
    user_id VARCHAR(50) NOT NULL,
    present BOOLEAN DEFAULT true,
    check_in_at TIMESTAMP WITH TIME ZONE,
    check_out_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS overtime (
    id VARCHAR(50) PRIMARY KEY,
    shift_id VARCHAR(50) REFERENCES shifts(id),
    user_id VARCHAR(50) NOT NULL,
    hours_worked NUMERIC(10, 2) DEFAULT 0,
    reason TEXT NOT NULL,
    approved_by VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS checklists (
    id VARCHAR(50) PRIMARY KEY,
    shift_id VARCHAR(50) REFERENCES shifts(id),
    check_name VARCHAR(200) NOT NULL,
    completed BOOLEAN DEFAULT false,
    completed_by VARCHAR(50),
    completed_at TIMESTAMP WITH TIME ZONE
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
