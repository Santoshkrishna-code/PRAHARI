-- Migration 011: Comments thread

CREATE TABLE IF NOT EXISTS maintenance_comments (
    id                VARCHAR(64)  PRIMARY KEY,
    maintenance_id    VARCHAR(64)  NOT NULL REFERENCES maintenance(id),
    author_id         VARCHAR(64)  NOT NULL,
    parent_comment_id VARCHAR(64)  REFERENCES maintenance_comments(id),
    body              TEXT         NOT NULL,
    created_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    is_deleted        BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_comments_maint ON maintenance_comments (maintenance_id);
