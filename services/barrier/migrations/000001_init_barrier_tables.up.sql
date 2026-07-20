CREATE TABLE IF NOT EXISTS barriers (
    id VARCHAR(50) PRIMARY KEY,
    barrier_code VARCHAR(100) UNIQUE NOT NULL,
    plant_id VARCHAR(50) NOT NULL,
    unit_id VARCHAR(50) NOT NULL,
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    type VARCHAR(50) NOT NULL,
    asset_id VARCHAR(50),
    sil_level VARCHAR(20),
    is_ipl BOOLEAN DEFAULT false,
    pfd_target NUMERIC(15, 6) DEFAULT 1.0,
    health_score NUMERIC(5, 2) DEFAULT 100.0,
    status VARCHAR(50) NOT NULL,
    last_proof_tested_at TIMESTAMP WITH TIME ZONE,
    next_proof_test_due TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS barrier_groups (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    group_name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS barrier_functions (
    id VARCHAR(50) PRIMARY KEY,
    barrier_id VARCHAR(50) REFERENCES barriers(id),
    function_name VARCHAR(200) NOT NULL,
    required_action TEXT NOT NULL,
    response_time_sec NUMERIC(10, 2) DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS safeguards (
    id VARCHAR(50) PRIMARY KEY,
    barrier_id VARCHAR(50) REFERENCES barriers(id),
    title VARCHAR(200) NOT NULL,
    description TEXT,
    safeguard_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS independent_protection_layers (
    id VARCHAR(50) PRIMARY KEY,
    barrier_id VARCHAR(50) REFERENCES barriers(id),
    ipl_name VARCHAR(200) NOT NULL,
    pfd_claimed NUMERIC(15, 6) NOT NULL,
    is_independent BOOLEAN DEFAULT true,
    is_auditable BOOLEAN DEFAULT true,
    is_specific BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS safety_instrumented_functions (
    id VARCHAR(50) PRIMARY KEY,
    barrier_id VARCHAR(50) REFERENCES barriers(id),
    sif_number VARCHAR(100) UNIQUE NOT NULL,
    target_sil VARCHAR(20) NOT NULL,
    achieved_sil VARCHAR(20) NOT NULL,
    spurious_trip_rate NUMERIC(10, 6) DEFAULT 0,
    proof_interval VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS safety_instrumented_systems (
    id VARCHAR(50) PRIMARY KEY,
    barrier_id VARCHAR(50) REFERENCES barriers(id),
    sis_name VARCHAR(200) NOT NULL,
    architecture VARCHAR(20) NOT NULL,
    vendor VARCHAR(100),
    logic_solver VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS pressure_relief_devices (
    id VARCHAR(50) PRIMARY KEY,
    barrier_id VARCHAR(50) REFERENCES barriers(id),
    tag_number VARCHAR(100) UNIQUE NOT NULL,
    set_pressure_bar NUMERIC(10, 2) NOT NULL,
    capacity_kg_h NUMERIC(12, 2) NOT NULL,
    last_tested_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS gas_detectors (
    id VARCHAR(50) PRIMARY KEY,
    barrier_id VARCHAR(50) REFERENCES barriers(id),
    tag_number VARCHAR(100) UNIQUE NOT NULL,
    gas_type VARCHAR(50) NOT NULL,
    alarm_threshold NUMERIC(10, 2) NOT NULL,
    location_code VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS fire_detectors (
    id VARCHAR(50) PRIMARY KEY,
    barrier_id VARCHAR(50) REFERENCES barriers(id),
    tag_number VARCHAR(100) UNIQUE NOT NULL,
    detector_type VARCHAR(50) NOT NULL,
    location_code VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS emergency_shutdown_systems (
    id VARCHAR(50) PRIMARY KEY,
    barrier_id VARCHAR(50) REFERENCES barriers(id),
    system_name VARCHAR(200) NOT NULL,
    valve_tag_number VARCHAR(100) NOT NULL,
    stroke_time_sec NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS alarms (
    id VARCHAR(50) PRIMARY KEY,
    barrier_id VARCHAR(50) REFERENCES barriers(id),
    tag_number VARCHAR(100) UNIQUE NOT NULL,
    priority VARCHAR(50) NOT NULL,
    setpoint_val NUMERIC(10, 2) NOT NULL,
    response_time NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS proof_tests (
    id VARCHAR(50) PRIMARY KEY,
    barrier_id VARCHAR(50) REFERENCES barriers(id),
    test_number VARCHAR(100) UNIQUE NOT NULL,
    executed_by VARCHAR(50) NOT NULL,
    passed BOOLEAN NOT NULL DEFAULT true,
    as_found_pfd NUMERIC(15, 6) DEFAULT 0,
    as_left_pfd NUMERIC(15, 6) DEFAULT 0,
    work_order_id VARCHAR(50),
    notes TEXT,
    executed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS barrier_inspections (
    id VARCHAR(50) PRIMARY KEY,
    barrier_id VARCHAR(50) REFERENCES barriers(id),
    inspector_id VARCHAR(50) NOT NULL,
    passes BOOLEAN DEFAULT true,
    defects_found TEXT,
    inspected_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS integrity_assessments (
    id VARCHAR(50) PRIMARY KEY,
    barrier_id VARCHAR(50) REFERENCES barriers(id),
    evaluator_id VARCHAR(50) NOT NULL,
    health_score NUMERIC(5, 2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    action_needed TEXT,
    assessed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS barrier_impairments (
    id VARCHAR(50) PRIMARY KEY,
    barrier_id VARCHAR(50) REFERENCES barriers(id),
    reason TEXT NOT NULL,
    compensating_ctrl TEXT NOT NULL,
    impaired_by VARCHAR(50) NOT NULL,
    is_active BOOLEAN DEFAULT true,
    impaired_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    restored_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS barrier_bypasses (
    id VARCHAR(50) PRIMARY KEY,
    barrier_id VARCHAR(50) REFERENCES barriers(id),
    permit_id VARCHAR(50),
    bypass_reason TEXT NOT NULL,
    approved_by VARCHAR(50) NOT NULL,
    authorized_period VARCHAR(100) NOT NULL,
    is_active BOOLEAN DEFAULT true,
    bypassed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    restored_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS performance_metrics (
    barrier_id VARCHAR(50) PRIMARY KEY REFERENCES barriers(id),
    availability_pct NUMERIC(5, 2) DEFAULT 100.0,
    proof_test_compliance_pct NUMERIC(5, 2) DEFAULT 100.0,
    bypass_hours NUMERIC(10, 2) DEFAULT 0,
    impairment_hours NUMERIC(10, 2) DEFAULT 0,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
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
