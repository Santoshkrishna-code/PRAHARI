-- Migration 008: Root cause analysis

CREATE TABLE IF NOT EXISTS incident_root_causes (
    id                  VARCHAR(64)  PRIMARY KEY,
    investigation_id    VARCHAR(64)  NOT NULL REFERENCES incident_investigations(id),
    incident_id         VARCHAR(64)  NOT NULL REFERENCES incidents(id),
    category            VARCHAR(32)  NOT NULL,
    description         TEXT         NOT NULL,
    contributing_factors JSONB,
    identified_by       VARCHAR(64)  NOT NULL,
    identified_at       TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_root_causes_investigation ON incident_root_causes (investigation_id);
CREATE INDEX idx_root_causes_incident      ON incident_root_causes (incident_id);
