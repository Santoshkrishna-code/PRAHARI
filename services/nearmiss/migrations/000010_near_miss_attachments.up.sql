-- Migration 010: S3 attachment references

CREATE TABLE IF NOT EXISTS near_miss_attachments (
    id           VARCHAR(64)  PRIMARY KEY,
    near_miss_id VARCHAR(64)  NOT NULL REFERENCES near_misses(id),
    file_name    VARCHAR(500) NOT NULL,
    file_size    BIGINT       NOT NULL,
    content_type VARCHAR(200) NOT NULL,
    storage_path TEXT         NOT NULL,
    uploaded_by  VARCHAR(64)  NOT NULL,
    uploaded_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_attachments_nearmiss ON near_miss_attachments (near_miss_id);
