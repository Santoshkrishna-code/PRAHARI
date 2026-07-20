-- Migration 007: Assigned auditors roster mappings

CREATE TABLE IF NOT EXISTS auditors (
    id       VARCHAR(64) PRIMARY KEY,
    audit_id VARCHAR(64) NOT NULL REFERENCES audits(id),
    user_id  VARCHAR(64) NOT NULL,
    role     VARCHAR(64) NOT NULL
);

CREATE INDEX idx_auditors_audit ON auditors (audit_id);
