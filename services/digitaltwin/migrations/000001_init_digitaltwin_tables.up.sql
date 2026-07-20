CREATE TABLE IF NOT EXISTS digital_twins (
    id VARCHAR(50) PRIMARY KEY,
    plant_id VARCHAR(50) NOT NULL,
    name VARCHAR(200) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'DRAFT',
    version INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS facilities (
    id VARCHAR(50) PRIMARY KEY,
    twin_id VARCHAR(50) REFERENCES digital_twins(id),
    name VARCHAR(200) NOT NULL,
    type VARCHAR(50) NOT NULL,
    parent_id VARCHAR(50),
    latitude NUMERIC(10, 8),
    longitude NUMERIC(11, 8),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS topologies (
    id VARCHAR(50) PRIMARY KEY,
    twin_id VARCHAR(50) REFERENCES digital_twins(id),
    facility_id VARCHAR(50) REFERENCES facilities(id),
    label VARCHAR(200) NOT NULL,
    node_type VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS equipment_graph (
    id VARCHAR(50) PRIMARY KEY,
    twin_id VARCHAR(50) REFERENCES digital_twins(id),
    external_id VARCHAR(50) NOT NULL,
    tag VARCHAR(100) NOT NULL,
    asset_type VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS zones (
    id VARCHAR(50) PRIMARY KEY,
    twin_id VARCHAR(50) REFERENCES digital_twins(id),
    facility_id VARCHAR(50) REFERENCES facilities(id),
    name VARCHAR(200) NOT NULL,
    zone_type VARCHAR(50) NOT NULL,
    coordinates TEXT NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS telemetry_bindings (
    id VARCHAR(50) PRIMARY KEY,
    twin_id VARCHAR(50) REFERENCES digital_twins(id),
    equipment_id VARCHAR(50) REFERENCES equipment_graph(id),
    sensor_tag VARCHAR(100) NOT NULL,
    unit VARCHAR(50) NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS live_states (
    id VARCHAR(50) PRIMARY KEY,
    twin_id VARCHAR(50) REFERENCES digital_twins(id),
    equipment_id VARCHAR(50) REFERENCES equipment_graph(id),
    value NUMERIC(15, 4) NOT NULL,
    quality VARCHAR(50) NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS simulations (
    id VARCHAR(50) PRIMARY KEY,
    twin_id VARCHAR(50) REFERENCES digital_twins(id),
    name VARCHAR(200) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
    parameters TEXT NOT NULL,
    result_data TEXT,
    started_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS playback_sessions (
    id VARCHAR(50) PRIMARY KEY,
    twin_id VARCHAR(50) REFERENCES digital_twins(id),
    start_time TIMESTAMP WITH TIME ZONE NOT NULL,
    end_time TIMESTAMP WITH TIME ZONE NOT NULL,
    speed NUMERIC(3, 2) NOT NULL DEFAULT 1.00,
    status VARCHAR(50) NOT NULL DEFAULT 'PLAYING',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS overlays (
    id VARCHAR(50) PRIMARY KEY,
    twin_id VARCHAR(50) REFERENCES digital_twins(id),
    layer_type VARCHAR(50) NOT NULL,
    source_id VARCHAR(50) NOT NULL,
    label VARCHAR(200) NOT NULL,
    metadata TEXT NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS alarms (
    id VARCHAR(50) PRIMARY KEY,
    twin_id VARCHAR(50) REFERENCES digital_twins(id),
    equipment_id VARCHAR(50) REFERENCES equipment_graph(id),
    severity VARCHAR(50) NOT NULL,
    message TEXT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT true,
    triggered_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS versions (
    id VARCHAR(50) PRIMARY KEY,
    twin_id VARCHAR(50) REFERENCES digital_twins(id),
    version INTEGER NOT NULL,
    label VARCHAR(100) NOT NULL,
    state_data TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
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
