-- Migration 018: S3 evidence attachments

CREATE TABLE IF NOT EXISTS attachments (
    id           VARCHAR(64)  PRIMARY KEY,
    audit_id     VARCHAR(64)  NOT NULL REFERENCES audits(id),
    file_name    VARCHAR(500) NOT NULL,
    file_size    BIGINT       NOT NULL,
    content_type VARCHAR(200) NOT NULL,
    storage_path TEXT         NOT NULL,
    uploaded_by  VARCHAR(64)  NOT NULL,
    uploaded_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_attachments_audit ON attachments (audit_id);
