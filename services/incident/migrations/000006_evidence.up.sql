-- Migration 006: Incident evidence with chain of custody

CREATE TABLE IF NOT EXISTS incident_evidence (
    id               VARCHAR(64)  PRIMARY KEY,
    incident_id      VARCHAR(64)  NOT NULL REFERENCES incidents(id),
    type             VARCHAR(32)  NOT NULL,
    description      TEXT         NOT NULL,
    storage_path     TEXT,
    collected_by     VARCHAR(64)  NOT NULL,
    collected_at     TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    chain_of_custody JSONB
);

CREATE INDEX idx_evidence_incident ON incident_evidence (incident_id);
CREATE INDEX idx_evidence_type     ON incident_evidence (type);
