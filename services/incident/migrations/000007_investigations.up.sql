-- Migration 007: Incident investigations

CREATE TABLE IF NOT EXISTS incident_investigations (
    id              VARCHAR(64)  PRIMARY KEY,
    incident_id     VARCHAR(64)  NOT NULL REFERENCES incidents(id),
    investigator_id VARCHAR(64)  NOT NULL,
    methodology     VARCHAR(64)  NOT NULL,
    findings        TEXT,
    recommendations TEXT,
    status          VARCHAR(32)  NOT NULL DEFAULT 'OPEN',
    started_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    completed_at    TIMESTAMPTZ,
    created_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_investigations_incident ON incident_investigations (incident_id);
CREATE INDEX idx_investigations_status   ON incident_investigations (status);
