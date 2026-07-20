-- Migration 009: CAPA (Corrective and Preventive Actions)

CREATE TABLE IF NOT EXISTS incident_capa (
    id                   VARCHAR(64)  PRIMARY KEY,
    incident_id          VARCHAR(64)  NOT NULL REFERENCES incidents(id),
    type                 VARCHAR(32)  NOT NULL,
    description          TEXT         NOT NULL,
    assignee_id          VARCHAR(64)  NOT NULL,
    due_date             TIMESTAMPTZ  NOT NULL,
    completed_at         TIMESTAMPTZ,
    status               VARCHAR(32)  NOT NULL DEFAULT 'OPEN',
    verified_by          VARCHAR(64),
    verified_at          TIMESTAMPTZ,
    effectiveness_review TEXT,
    created_at           TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at           TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_capa_incident ON incident_capa (incident_id);
CREATE INDEX idx_capa_status   ON incident_capa (status);
CREATE INDEX idx_capa_due      ON incident_capa (due_date) WHERE status NOT IN ('COMPLETED', 'VERIFIED');
