-- Migration 007: Compliance check controls procedures

CREATE TABLE IF NOT EXISTS controls (
    id            VARCHAR(64) PRIMARY KEY,
    compliance_id VARCHAR(64) NOT NULL REFERENCES compliance_register(id),
    description   TEXT        NOT NULL
);

CREATE INDEX idx_controls_compliance ON controls (compliance_id);
