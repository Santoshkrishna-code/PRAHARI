-- Migration 009: Findings gaps checklist logs

CREATE TABLE IF NOT EXISTS findings (
    id           VARCHAR(64) PRIMARY KEY,
    audit_id     VARCHAR(64) NOT NULL REFERENCES audits(id),
    finding_type VARCHAR(64) NOT NULL,
    description  TEXT        NOT NULL
);

CREATE INDEX idx_findings_audit ON findings (audit_id);
