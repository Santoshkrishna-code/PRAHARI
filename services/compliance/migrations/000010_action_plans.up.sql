-- Migration 010: Corrective Action plans tasks

CREATE TABLE IF NOT EXISTS action_plans (
    id           VARCHAR(64) PRIMARY KEY,
    finding_id   VARCHAR(64) NOT NULL REFERENCES findings(id),
    description  TEXT        NOT NULL,
    target_date  TIMESTAMPTZ NOT NULL,
    is_completed BOOLEAN     NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_actions_finding ON action_plans (finding_id);
