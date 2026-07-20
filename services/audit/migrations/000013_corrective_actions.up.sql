-- Migration 013: Corrective actions plans (CAPA) checklists tasks

CREATE TABLE IF NOT EXISTS corrective_actions (
    id           VARCHAR(64) PRIMARY KEY,
    finding_id   VARCHAR(64) NOT NULL REFERENCES findings(id),
    description  TEXT        NOT NULL,
    target_date  TIMESTAMPTZ NOT NULL,
    is_completed BOOLEAN     NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_capa_finding ON corrective_actions (finding_id);
