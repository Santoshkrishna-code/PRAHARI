-- Migration 011: Audit log and timeline events

-- Append-only audit log
CREATE TABLE IF NOT EXISTS incident_audit (
    id          VARCHAR(64)  PRIMARY KEY,
    entity_type VARCHAR(64)  NOT NULL,
    entity_id   VARCHAR(64)  NOT NULL,
    action      VARCHAR(32)  NOT NULL,
    actor_id    VARCHAR(64)  NOT NULL,
    old_value   JSONB,
    new_value   JSONB,
    ip_address  VARCHAR(45),
    user_agent  TEXT,
    timestamp   TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_audit_entity   ON incident_audit (entity_type, entity_id);
CREATE INDEX idx_audit_actor    ON incident_audit (actor_id);
CREATE INDEX idx_audit_time     ON incident_audit (timestamp);

-- Immutable timeline events
CREATE TABLE IF NOT EXISTS incident_timeline (
    id          VARCHAR(64)  PRIMARY KEY,
    incident_id VARCHAR(64)  NOT NULL REFERENCES incidents(id),
    event_type  VARCHAR(64)  NOT NULL,
    actor_id    VARCHAR(64)  NOT NULL,
    description TEXT         NOT NULL,
    metadata    JSONB,
    occurred_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_timeline_incident ON incident_timeline (incident_id);
CREATE INDEX idx_timeline_time     ON incident_timeline (occurred_at);
