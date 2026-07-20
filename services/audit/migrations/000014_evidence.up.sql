-- Migration 014: Evidence references metadata uploads

CREATE TABLE IF NOT EXISTS evidence (
    id             VARCHAR(64) PRIMARY KEY,
    audit_id       VARCHAR(64) NOT NULL REFERENCES audits(id),
    uploaded_by_id VARCHAR(64) NOT NULL,
    storage_path   TEXT        NOT NULL,
    collected_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_evidence_audit ON evidence (audit_id);
