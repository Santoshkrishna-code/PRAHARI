-- Migration 005: Incident attachments

CREATE TABLE IF NOT EXISTS incident_attachments (
    id           VARCHAR(64)  PRIMARY KEY,
    incident_id  VARCHAR(64)  NOT NULL REFERENCES incidents(id),
    file_name    VARCHAR(500) NOT NULL,
    file_size    BIGINT       NOT NULL,
    content_type VARCHAR(200) NOT NULL,
    storage_path TEXT         NOT NULL,
    uploaded_by  VARCHAR(64)  NOT NULL,
    uploaded_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_attachments_incident ON incident_attachments (incident_id);
