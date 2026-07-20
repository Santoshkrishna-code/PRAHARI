-- Migration 010: Attachments

CREATE TABLE IF NOT EXISTS permit_attachments (
    id           VARCHAR(64)  PRIMARY KEY,
    permit_id    VARCHAR(64)  NOT NULL REFERENCES permits(id),
    file_name    VARCHAR(500) NOT NULL,
    file_size    BIGINT       NOT NULL,
    content_type VARCHAR(200) NOT NULL,
    storage_path TEXT         NOT NULL,
    uploaded_by  VARCHAR(64)  NOT NULL,
    uploaded_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_attachments_permit ON permit_attachments (permit_id);
