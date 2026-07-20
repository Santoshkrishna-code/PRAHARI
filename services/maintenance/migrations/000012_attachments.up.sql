-- Migration 012: Attachments S3 references

CREATE TABLE IF NOT EXISTS maintenance_attachments (
    id             VARCHAR(64)  PRIMARY KEY,
    maintenance_id VARCHAR(64)  NOT NULL REFERENCES maintenance(id),
    file_name      VARCHAR(500) NOT NULL,
    file_size      BIGINT       NOT NULL,
    content_type   VARCHAR(200) NOT NULL,
    storage_path   TEXT         NOT NULL,
    uploaded_by    VARCHAR(64)  NOT NULL,
    uploaded_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_attachments_maint ON maintenance_attachments (maintenance_id);
