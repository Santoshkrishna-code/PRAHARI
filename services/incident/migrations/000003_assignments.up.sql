-- Migration 003: Incident assignments

CREATE TABLE IF NOT EXISTS incident_assignments (
    id          VARCHAR(64)  PRIMARY KEY,
    incident_id VARCHAR(64)  NOT NULL REFERENCES incidents(id),
    assignee_id VARCHAR(64)  NOT NULL,
    assigner_id VARCHAR(64)  NOT NULL,
    role        VARCHAR(32)  NOT NULL,
    assigned_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    accepted_at TIMESTAMPTZ,
    is_active   BOOLEAN      NOT NULL DEFAULT TRUE,
    note        TEXT
);

CREATE INDEX idx_assignments_incident ON incident_assignments (incident_id);
CREATE INDEX idx_assignments_assignee ON incident_assignments (assignee_id);
