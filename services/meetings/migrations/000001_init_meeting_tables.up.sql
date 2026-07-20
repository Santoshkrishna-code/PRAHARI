CREATE TABLE IF NOT EXISTS meetings (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    meeting_type VARCHAR(50) NOT NULL,
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    location VARCHAR(200) NOT NULL DEFAULT '',
    scheduled_at TIMESTAMP WITH TIME ZONE NOT NULL,
    started_at TIMESTAMP WITH TIME ZONE,
    ended_at TIMESTAMP WITH TIME ZONE,
    organizer_id VARCHAR(50) NOT NULL,
    facilitator_id VARCHAR(50),
    shift_id VARCHAR(50),
    permit_id VARCHAR(50),
    status VARCHAR(50) NOT NULL DEFAULT 'PLANNED',
    duration_min INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS toolbox_talks (
    id VARCHAR(50) PRIMARY KEY,
    meeting_id VARCHAR(50) REFERENCES meetings(id),
    topic_title VARCHAR(200) NOT NULL,
    topic_body TEXT NOT NULL DEFAULT '',
    category VARCHAR(50) NOT NULL,
    mandatory BOOLEAN NOT NULL DEFAULT false,
    frequency_days INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS prejob_briefings (
    id VARCHAR(50) PRIMARY KEY,
    meeting_id VARCHAR(50) REFERENCES meetings(id),
    permit_id VARCHAR(50),
    work_order_id VARCHAR(50),
    hazards_summary TEXT NOT NULL DEFAULT '',
    ppe_required TEXT NOT NULL DEFAULT '',
    emergency_plan TEXT NOT NULL DEFAULT '',
    acknowledged BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS safety_meetings (
    id VARCHAR(50) PRIMARY KEY,
    meeting_id VARCHAR(50) REFERENCES meetings(id),
    frequency VARCHAR(50) NOT NULL,
    department VARCHAR(100) NOT NULL,
    review_period VARCHAR(50) NOT NULL DEFAULT '',
    chair_id VARCHAR(50) NOT NULL,
    secretary_id VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS committees (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    name VARCHAR(200) NOT NULL,
    charter TEXT NOT NULL DEFAULT '',
    chair_id VARCHAR(50) NOT NULL,
    secretary_id VARCHAR(50),
    meet_freq VARCHAR(50) NOT NULL DEFAULT 'MONTHLY',
    active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS agendas (
    id VARCHAR(50) PRIMARY KEY,
    meeting_id VARCHAR(50) REFERENCES meetings(id),
    seq_order INTEGER NOT NULL DEFAULT 0,
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    presenter_id VARCHAR(50),
    duration_min INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS attendances (
    id VARCHAR(50) PRIMARY KEY,
    meeting_id VARCHAR(50) REFERENCES meetings(id),
    attendee_id VARCHAR(50) NOT NULL,
    attendee_name VARCHAR(200) NOT NULL DEFAULT '',
    check_in_at TIMESTAMP WITH TIME ZONE NOT NULL,
    check_out_at TIMESTAMP WITH TIME ZONE,
    method VARCHAR(50) NOT NULL DEFAULT 'MANUAL',
    verified BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS acknowledgements (
    id VARCHAR(50) PRIMARY KEY,
    meeting_id VARCHAR(50) REFERENCES meetings(id),
    attendee_id VARCHAR(50) NOT NULL,
    signature_url TEXT,
    acknowledged_at TIMESTAMP WITH TIME ZONE NOT NULL,
    ip_address VARCHAR(50),
    device_id VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS minutes (
    id VARCHAR(50) PRIMARY KEY,
    meeting_id VARCHAR(50) REFERENCES meetings(id),
    body TEXT NOT NULL DEFAULT '',
    recorder_id VARCHAR(50) NOT NULL,
    approver_id VARCHAR(50),
    approved_at TIMESTAMP WITH TIME ZONE,
    status VARCHAR(50) NOT NULL DEFAULT 'DRAFT',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS discussion_topics (
    id VARCHAR(50) PRIMARY KEY,
    meeting_id VARCHAR(50) REFERENCES meetings(id),
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    source_type VARCHAR(50),
    source_ref_id VARCHAR(50),
    raised_by VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS generated_actions (
    id VARCHAR(50) PRIMARY KEY,
    meeting_id VARCHAR(50) REFERENCES meetings(id),
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    assigned_to VARCHAR(50) NOT NULL,
    due_date TIMESTAMP WITH TIME ZONE NOT NULL,
    priority VARCHAR(50) NOT NULL DEFAULT 'MEDIUM',
    synced_to_capa BOOLEAN NOT NULL DEFAULT false,
    capa_ref_id VARCHAR(50),
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
