-- Migration 007: Workflow digital signatures approvals

CREATE TABLE IF NOT EXISTS risk_approvals (
    id            VARCHAR(64) PRIMARY KEY,
    risk_id       VARCHAR(64) NOT NULL REFERENCES risk_assessments(id),
    approver_id   VARCHAR(64) NOT NULL,
    approved_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    signature     TEXT        NOT NULL
);

CREATE INDEX idx_approvals_risk ON risk_approvals (risk_id);
