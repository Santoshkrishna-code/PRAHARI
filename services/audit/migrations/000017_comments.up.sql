-- Migration 017: Comment threads

CREATE TABLE IF NOT EXISTS comments (
    id                VARCHAR(64)  PRIMARY KEY,
    audit_id          VARCHAR(64)  NOT NULL REFERENCES audits(id),
    author_id         VARCHAR(64)  NOT NULL,
    parent_comment_id VARCHAR(64)  REFERENCES comments(id),
    body              TEXT         NOT NULL,
    created_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    is_deleted        BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_comments_audit ON comments (audit_id);
