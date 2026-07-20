-- Migration 009: Compliance findings gaps logs

CREATE TABLE IF NOT EXISTS findings (
    id            VARCHAR(64) PRIMARY KEY,
    compliance_id VARCHAR(64) NOT NULL REFERENCES compliance_register(id),
    severity      VARCHAR(32) NOT NULL DEFAULT 'MINOR',
    description   TEXT        NOT NULL
);

CREATE INDEX idx_findings_compliance ON findings (compliance_id);
