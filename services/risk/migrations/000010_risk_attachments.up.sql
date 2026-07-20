-- Migration 010: S3 attachment references evidence

CREATE TABLE IF NOT EXISTS risk_attachments (
    id           VARCHAR(64)  PRIMARY KEY,
    risk_id      VARCHAR(64)  NOT NULL REFERENCES risk_assessments(id),
    file_name    VARCHAR(500) NOT NULL,
    file_size    BIGINT       NOT NULL,
    content_type VARCHAR(200) NOT NULL,
    storage_path TEXT         NOT NULL,
    uploaded_by  VARCHAR(64)  NOT NULL,
    uploaded_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_risk_attachments_id ON risk_attachments (risk_id);
