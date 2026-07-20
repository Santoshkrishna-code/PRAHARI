-- Migration 004: Threaded incident comments

CREATE TABLE IF NOT EXISTS incident_comments (
    id                VARCHAR(64)  PRIMARY KEY,
    incident_id       VARCHAR(64)  NOT NULL REFERENCES incidents(id),
    author_id         VARCHAR(64)  NOT NULL,
    parent_comment_id VARCHAR(64)  REFERENCES incident_comments(id),
    body              TEXT         NOT NULL,
    created_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    is_deleted        BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_comments_incident ON incident_comments (incident_id);
CREATE INDEX idx_comments_parent   ON incident_comments (parent_comment_id);
