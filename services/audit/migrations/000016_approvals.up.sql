-- Migration 016: Workflow digital signature approvals

CREATE TABLE IF NOT EXISTS approvals (
    id            VARCHAR(64) PRIMARY KEY,
    audit_id      VARCHAR(64) NOT NULL REFERENCES audits(id),
    approver_id   VARCHAR(64) NOT NULL,
    approved_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    signature     TEXT        NOT NULL
);

CREATE INDEX idx_approvals_audit ON approvals (audit_id);
