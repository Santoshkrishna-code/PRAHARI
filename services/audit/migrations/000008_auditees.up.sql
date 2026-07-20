-- Migration 008: Assigned auditees roster mappings

CREATE TABLE IF NOT EXISTS auditees (
    id       VARCHAR(64) PRIMARY KEY,
    audit_id VARCHAR(64) NOT NULL REFERENCES audits(id),
    user_id  VARCHAR(64) NOT NULL
);

CREATE INDEX idx_auditees_audit ON auditees (audit_id);
