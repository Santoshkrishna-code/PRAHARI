-- Migration 007: CAPA actions

CREATE TABLE IF NOT EXISTS inspection_actions (
    id            VARCHAR(64)  PRIMARY KEY,
    inspection_id VARCHAR(64)  NOT NULL REFERENCES inspections(id),
    finding_id    VARCHAR(64)  REFERENCES inspection_findings(id),
    action_type   VARCHAR(32)  NOT NULL,
    description   TEXT         NOT NULL,
    assignee_id   VARCHAR(64)  NOT NULL,
    due_date      TIMESTAMPTZ  NOT NULL,
    completed_at  TIMESTAMPTZ,
    status        VARCHAR(32)  NOT NULL DEFAULT 'OPEN',
    verified_by   VARCHAR(64),
    verified_at   TIMESTAMPTZ,
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_actions_inspection ON inspection_actions (inspection_id);
CREATE INDEX idx_actions_status     ON inspection_actions (status);
